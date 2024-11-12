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

// FindItemNames accesses the given flags and attempts to perform a search for the given item type
func FindItemNames(cmd *cobra.Command, args []string, what string) error {
	_ = args
	criteria := make(map[string]interface{})
	cmd.Flags().Visit(func(flag *pflag.Flag) {
		if flag.Name == "config" {
			return
		}
		key := strings.Replace(flag.Name, "-", "_", -1)
		criteria[key] = flag.Value.String()
	})

	// Now perform the actual search
	itemNames, err := Client.FindItemNames(what, criteria, "name")
	if err != nil {
		return err
	}
	for _, distroName := range itemNames {
		fmt.Fprintln(cmd.OutOrStdout(), distroName)
	}
	return nil
}
