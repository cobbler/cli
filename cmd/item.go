package cmd

import (
	"fmt"
	cobbler "github.com/cobbler/cobblerclient"
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

// resolveUID resolves the target item identified by the --name/--uid flag pair to its
// Cobbler UID. Cobbler 4.0.0's get_item/get_<type> and remove_item/remove_<type> XML-RPC
// calls require an object UID rather than a name (names aren't guaranteed to be globally
// unique for every item type, e.g. NetworkInterface names are only unique per-system), so
// every CLI subcommand that lets a human identify its target by name has to resolve that
// name to a UID before calling into cobblerclient.
//
// If uid is non-empty it is returned directly without a server round-trip. Otherwise name
// is resolved via Client.FindItems using an exact-match, expanded search:
//   - zero matches is an error ("no <what> found with name ...")
//   - exactly one match returns that item's uid
//   - more than one match is an error telling the caller to use --uid instead
//
// If neither name nor uid is supplied, an error asking for one of --name/--uid is returned.
func resolveUID(client *cobbler.Client, what, name, uid string) (string, error) {
	if uid != "" {
		return uid, nil
	}
	if name == "" {
		return "", fmt.Errorf("either --name or --uid must be supplied to identify the %s", what)
	}
	results, err := client.FindItems(what, map[string]interface{}{"name": name}, "", true)
	if err != nil {
		return "", err
	}
	switch len(results) {
	case 0:
		return "", fmt.Errorf("no %s found with name %q", what, name)
	case 1:
		asMap, ok := results[0].(map[string]interface{})
		if !ok {
			return "", fmt.Errorf("unexpected result type for %s %q", what, name)
		}
		uidValue, ok := asMap["uid"].(string)
		if !ok || uidValue == "" {
			return "", fmt.Errorf("no uid found in result for %s %q", what, name)
		}
		return uidValue, nil
	default:
		return "", fmt.Errorf("multiple %s items found with name %q; use --uid to disambiguate", what, name)
	}
}

// addUIDFlag registers a --uid sibling flag next to a target-identifying --name flag,
// letting the user identify the item this command acts on by its Cobbler UID instead of
// (or, together with resolveUID, alongside) its name. what is used only for the help text,
// e.g. addUIDFlag(cmd, "distro") registers "the distro uid (alternative to --name)".
func addUIDFlag(command *cobra.Command, what string) {
	command.Flags().String("uid", "", fmt.Sprintf("the %s uid (alternative to --name)", what))
}

// RemoveItemRecursive accesses the given flags and attempts to remove a given item
func RemoveItemRecursive(cmd *cobra.Command, args []string, what string) error {
	_ = args
	itemName, err := cmd.Flags().GetString("name")
	if err != nil {
		return err
	}
	itemUID, err := cmd.Flags().GetString("uid")
	if err != nil {
		return err
	}
	recursiveDelete, err := cmd.Flags().GetBool("recursive")
	if err != nil {
		return err
	}
	uid, err := resolveUID(&Client, what, itemName, itemUID)
	if err != nil {
		return err
	}
	return Client.RemoveItem(what, uid, recursiveDelete)
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
