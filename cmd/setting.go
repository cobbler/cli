// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: 2021 Dominik Gedon <dgedon@suse.de>
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	cobbler "github.com/cobbler/cobblerclient"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// NewSettingCmd builds a new command that represents the setting action
func NewSettingCmd() *cobra.Command {
	settingCmd := &cobra.Command{
		Use:   "setting",
		Short: "Settings management",
		Long:  `Let you manage settings.`,
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}
	settingCmd.AddCommand(NewSettingEditCmd())
	settingCmd.AddCommand(NewSettingReportCmd())
	settingCmd.AddCommand(NewSettingExportCmd())
	return settingCmd
}

// settingsFieldByMapstructureTag walks the Settings struct and returns the
// reflect.Value of the field whose `mapstructure` tag matches name (or the
// zero Value if not found).
func settingsFieldByMapstructureTag(settings *cobbler.Settings, name string) reflect.Value {
	v := reflect.ValueOf(settings).Elem()
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Tag.Get("mapstructure") == name {
			return v.Field(i)
		}
	}
	return reflect.Value{}
}

// parseSettingValue coerces the raw string the user passed via --value into
// the Go-typed value the corresponding Settings field expects. Cobbler 4.0.0
// is strict about types on the wire, so coercion happens here, not on the
// server.
func parseSettingValue(field reflect.Value, raw string) (interface{}, error) {
	switch field.Kind() {
	case reflect.Bool:
		return strconv.ParseBool(raw)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		bitSize := field.Type().Bits()
		v, err := strconv.ParseInt(raw, 10, bitSize)
		if err != nil {
			return nil, err
		}
		switch field.Kind() {
		case reflect.Int8:
			return int8(v), nil
		case reflect.Int16:
			return int16(v), nil
		case reflect.Int32:
			return int32(v), nil
		case reflect.Int64:
			return v, nil
		default:
			return int(v), nil
		}
	case reflect.Float32, reflect.Float64:
		return strconv.ParseFloat(raw, 64)
	case reflect.String:
		return raw, nil
	case reflect.Slice:
		// Accept comma-separated list for []string and similar.
		if field.Type().Elem().Kind() == reflect.String {
			if raw == "" {
				return []string{}, nil
			}
			parts := strings.Split(raw, ",")
			out := make([]string, len(parts))
			for i, p := range parts {
				out[i] = strings.TrimSpace(p)
			}
			return out, nil
		}
	case reflect.Map:
		// Accept key=value,key=value form for map[string]interface{}.
		out := map[string]interface{}{}
		if raw == "" {
			return out, nil
		}
		for _, pair := range strings.Split(raw, ",") {
			kv := strings.SplitN(pair, "=", 2)
			if len(kv) != 2 {
				return nil, fmt.Errorf("malformed map entry %q (expected key=value)", pair)
			}
			out[strings.TrimSpace(kv[0])] = strings.TrimSpace(kv[1])
		}
		return out, nil
	}
	return nil, fmt.Errorf("unsupported setting type %s", field.Kind())
}

func NewSettingEditCmd() *cobra.Command {
	settingEditCmd := &cobra.Command{
		Use:   "edit",
		Short: "edit settings",
		Long:  `Edits a setting. The value is parsed into the same type the existing setting has, so --value=42 sends an int when the server-side type is int.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := generateCobblerClient(); err != nil {
				return err
			}
			settings, err := Client.GetSettings()
			if err != nil {
				return err
			}
			if !settings.AllowDynamicSettings {
				return errors.New("dynamic settings are turned off server-side")
			}
			name, err := cmd.Flags().GetString("name")
			if err != nil {
				return err
			}
			raw, err := cmd.Flags().GetString("value")
			if err != nil {
				return err
			}
			field := settingsFieldByMapstructureTag(settings, name)
			if !field.IsValid() {
				return fmt.Errorf("unknown setting %q; run `cobbler setting report` to list valid names", name)
			}
			typed, err := parseSettingValue(field, raw)
			if err != nil {
				return fmt.Errorf("cannot parse --value %q as %s: %w", raw, field.Kind(), err)
			}
			result, err := Client.ModifySetting(name, typed)
			if err != nil {
				return err
			}
			if result == 0 {
				_, _ = fmt.Fprintln(cmd.OutOrStdout(), "Successfully updated!")
			} else {
				_, _ = fmt.Fprintln(cmd.OutOrStdout(), "Updating settings failed!")
			}
			return nil
		},
	}
	settingEditCmd.Flags().String("name", "", "the settings name to edit (e.g. server)")
	settingEditCmd.Flags().String("value", "", "the new value (parsed according to the setting's type)")
	return settingEditCmd
}

func NewSettingReportCmd() *cobra.Command {
	settingReportCmd := &cobra.Command{
		Use:   "report",
		Short: "list settings",
		Long:  `Prints settings to stdout.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := generateCobblerClient(); err != nil {
				return err
			}
			settings, err := Client.GetSettings()
			if err != nil {
				return err
			}
			printStructured(cmd, settings)
			return nil
		},
	}
	settingReportCmd.Flags().String("name", "", "the settings name to show")
	return settingReportCmd
}

func NewSettingExportCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "export",
		Short: "export settings",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			format, err := cmd.Flags().GetString("format")
			if err != nil {
				return err
			}
			if format != "json" && format != "yaml" {
				return fmt.Errorf("format must be json or yaml")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := generateCobblerClient(); err != nil {
				return err
			}
			settings, err := Client.GetSettings()
			if err != nil {
				return err
			}
			format, err := cmd.Flags().GetString("format")
			if err != nil {
				return err
			}
			switch format {
			case "json":
				out, err := json.Marshal(settings)
				if err != nil {
					return err
				}
				_, _ = fmt.Fprintln(cmd.OutOrStdout(), string(out))
			case "yaml":
				out, err := yaml.Marshal(settings)
				if err != nil {
					return err
				}
				_, _ = fmt.Fprintln(cmd.OutOrStdout(), "---")
				_, _ = fmt.Fprintln(cmd.OutOrStdout(), string(out))
			}
			return nil
		},
	}
	cmd.Flags().String(exportStringMetadata["format"].Name, exportStringMetadata["format"].DefaultValue, exportStringMetadata["format"].Usage)
	return cmd
}
