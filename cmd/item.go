package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"strings"
)

var inheritedUsageFormat = "Mark %s as inherited and remove its concrete value"

func addStringFlags(command *cobra.Command, metadata map[string]FlagMetadata[string]) {
	for _, value := range metadata {
		command.Flags().String(value.Name, value.DefaultValue, value.Usage)
	}
}

func addBoolFlags(command *cobra.Command, metadata map[string]FlagMetadata[bool]) {
	for _, value := range metadata {
		if value.IsInheritable {
			var inheritFlagName = value.Name + "-inherit"
			command.Flags().Bool(value.Name, value.DefaultValue, value.Usage)
			command.Flags().Bool(
				inheritFlagName,
				false,
				fmt.Sprintf(inheritedUsageFormat, value.Name),
			)
		} else {
			command.Flags().Bool(value.Name, value.DefaultValue, value.Usage)
		}
	}
}

func addIntFlags(command *cobra.Command, metadata map[string]FlagMetadata[int]) {
	for _, value := range metadata {
		if value.IsInheritable {
			var inheritFlagName = value.Name + "-inherit"
			command.Flags().Int(value.Name, value.DefaultValue, value.Usage)
			command.Flags().Bool(
				inheritFlagName,
				false,
				fmt.Sprintf(inheritedUsageFormat, value.Name),
			)
		} else {
			command.Flags().Int(value.Name, value.DefaultValue, value.Usage)
		}
	}
}

func addFloatFlags(command *cobra.Command, metadata map[string]FlagMetadata[float64]) {
	for _, value := range metadata {
		if value.IsInheritable {
			var inheritFlagName = value.Name + "-inherit"
			command.Flags().Float64(value.Name, value.DefaultValue, value.Usage)
			command.Flags().Bool(
				inheritFlagName,
				false,
				fmt.Sprintf(inheritedUsageFormat, value.Name),
			)
		} else {
			command.Flags().Float64(value.Name, value.DefaultValue, value.Usage)
		}
	}
}

func addStringSliceFlags(command *cobra.Command, metadata map[string]FlagMetadata[[]string]) {
	for _, value := range metadata {
		if value.IsInheritable {
			var inheritedFlagName = value.Name + "-inherit"
			command.Flags().StringSlice(value.Name, value.DefaultValue, value.Usage)
			command.Flags().Bool(
				inheritedFlagName,
				false,
				fmt.Sprintf(inheritedUsageFormat, value.Name),
			)
			command.MarkFlagsMutuallyExclusive(value.Name, inheritedFlagName)
		} else {
			command.Flags().StringSlice(value.Name, value.DefaultValue, value.Usage)
		}
	}
}

func addMapFlags(command *cobra.Command, metadata map[string]FlagMetadata[map[string]string]) {
	for _, value := range metadata {
		if value.IsInheritable {
			var inheritedFlagName = value.Name + "-inherit"
			command.Flags().StringToString(value.Name, value.DefaultValue, value.Usage)
			command.Flags().Bool(
				inheritedFlagName,
				false,
				fmt.Sprintf(inheritedUsageFormat, value.Name),
			)
			command.MarkFlagsMutuallyExclusive(value.Name, inheritedFlagName)
		} else {
			command.Flags().StringToString(value.Name, value.DefaultValue, value.Usage)
		}
	}
}

func addCommonArgs(command *cobra.Command) {
	addStringFlags(command, commonStringFlagMetadata)
	addStringSliceFlags(command, commonStringSliceFlagMetadata)
}

// RemoveItemRecursive accesses the given flags and attempts to remove a given item
func RemoveItemRecursive(cmd *cobra.Command, args []string, what string) error {
	_ = args
	itemName, err := cmd.Flags().GetString("name")
	if err != nil {
		return err
	}
	recursiveDelete, err := cmd.Flags().GetBool("recursive")
	if err != nil {
		return err
	}
	return Client.RemoveItem(what, itemName, recursiveDelete)
}

// addPaginationFlags registers --page and --items-per-page on a find subcommand.
// When either is supplied the find handler routes through Client.FindItemsPaged
// and emits a trailing `# page N of M (T total)` summary line.
func addPaginationFlags(cmd *cobra.Command) {
	cmd.Flags().Int("page", 0, "page number for paginated find (1-based; omit for unpaged)")
	cmd.Flags().Int("items-per-page", 0, "results per page (only honoured when --page is set)")
}

// FindItemNames accesses the given flags and performs a search for items of the
// given type. When --page is specified it uses the paginated backend endpoint
// and prints a trailing summary line; otherwise it falls back to the unpaged
// find.
func FindItemNames(cmd *cobra.Command, args []string, what string) error {
	_ = args
	page, _ := cmd.Flags().GetInt("page")
	itemsPerPage, _ := cmd.Flags().GetInt("items-per-page")
	criteria := make(map[string]interface{})
	cmd.Flags().Visit(func(flag *pflag.Flag) {
		switch flag.Name {
		case "config", "page", "items-per-page":
			return
		}
		key := strings.ReplaceAll(flag.Name, "-", "_")
		criteria[key] = flag.Value.String()
	})

	if page > 0 {
		if itemsPerPage <= 0 {
			itemsPerPage = 20
		}
		result, err := Client.FindItemsPaged(what, criteria, "name", int32(page), int32(itemsPerPage))
		if err != nil {
			return err
		}
		for _, raw := range result.FoundItems {
			if asMap, ok := raw.(map[string]interface{}); ok {
				if name, ok := asMap["name"].(string); ok {
					_, _ = fmt.Fprintln(cmd.OutOrStdout(), name)
					continue
				}
			}
			_, _ = fmt.Fprintln(cmd.OutOrStdout(), raw)
		}
		_, _ = fmt.Fprintf(cmd.OutOrStdout(), "# page %d of %d (%d total)\n",
			result.PageInfo.Page, result.PageInfo.NumPages, result.PageInfo.NumItems)
		return nil
	}

	itemNames, err := Client.FindItemNames(what, criteria, "name")
	if err != nil {
		return err
	}
	for _, name := range itemNames {
		_, _ = fmt.Fprintln(cmd.OutOrStdout(), name)
	}
	return nil
}
