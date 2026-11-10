package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// convertMapStringToMapInterface converts the flat map[string]string produced by a stringToString flag (e.g.
// --autoinstall-meta) into the map[string]interface{} expected by the XML-RPC API. Values are copied through
// literally - a flag value is never implicitly reinterpreted as JSON, since that risks silently corrupting an
// innocuous-looking flat value (e.g. an IPv6-literal kernel option like "ip=[::1]"). Fields that genuinely need
// a list or dict value (e.g. Agama's agama_modules or cloud-init's per-module settings) have a dedicated "-json"
// flag (e.g. --autoinstall-meta-json) for that instead.
func convertMapStringToMapInterface(m map[string]string) map[string]interface{} {
	result := make(map[string]interface{}, len(m))
	for k, v := range m {
		result[k] = v
	}
	return result
}

// mergeJSONObjectFlag decodes jsonFlagName's value (if the flag was set) as a JSON object and
// merges it on top of base in place, overwriting any key base already has. Used by the "-json"
// companion flags (--autoinstall-meta-json, --kernel-options-json, --kernel-options-post-json) that
// let a value carry a list/nested-object the flat key=value stringToString flag syntax can't
// represent.
func mergeJSONObjectFlag(cmd *cobra.Command, jsonFlagName string, base map[string]interface{}) error {
	if !cmd.Flags().Lookup(jsonFlagName).Changed {
		return nil
	}
	raw, err := cmd.Flags().GetString(jsonFlagName)
	if err != nil {
		return err
	}
	decoded := map[string]interface{}{}
	if err := json.Unmarshal([]byte(raw), &decoded); err != nil {
		return err
	}
	for k, v := range decoded {
		base[k] = v
	}
	return nil
}

// mergeJSONStringMapFlag is mergeJSONObjectFlag's flat map[string]string counterpart, used by
// --template-files-json.
func mergeJSONStringMapFlag(cmd *cobra.Command, jsonFlagName string, base map[string]string) error {
	if !cmd.Flags().Lookup(jsonFlagName).Changed {
		return nil
	}
	raw, err := cmd.Flags().GetString(jsonFlagName)
	if err != nil {
		return err
	}
	decoded := map[string]string{}
	if err := json.Unmarshal([]byte(raw), &decoded); err != nil {
		return err
	}
	for k, v := range decoded {
		base[k] = v
	}
	return nil
}

func covertFloatToUtcTime(t float64) (time.Time, error) {
	parts := strings.Split(fmt.Sprintf("%f", t), ".")
	seconds, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return time.Time{}, err
	}
	nanoSeconds, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		return time.Time{}, err
	}
	timezone, err := time.LoadLocation("UTC")
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(seconds, nanoSeconds).In(timezone), nil
}

func printDumpVars(cmd *cobra.Command, blendedData map[string]interface{}) {
	for key, value := range blendedData {
		if value == nil {
			_, _ = fmt.Fprintf(cmd.OutOrStdout(), "%s:\n", key)
			continue
		}
		valueType := reflect.TypeOf(value).Kind()
		switch valueType {
		case reflect.Bool:
			_, _ = fmt.Fprintf(cmd.OutOrStdout(), "%s: %t\n", key, value.(bool))
		case reflect.Int64:
			_, _ = fmt.Fprintf(cmd.OutOrStdout(), "%s: %d\n", key, value.(int64))
		case reflect.Int32:
			_, _ = fmt.Fprintf(cmd.OutOrStdout(), "%s: %d\n", key, value.(int32))
		case reflect.Int16:
			_, _ = fmt.Fprintf(cmd.OutOrStdout(), "%s: %d\n", key, value.(int16))
		case reflect.Int8:
			_, _ = fmt.Fprintf(cmd.OutOrStdout(), "%s: %d\n", key, value.(int8))
		case reflect.Int:
			_, _ = fmt.Fprintf(cmd.OutOrStdout(), "%s: %d\n", key, value.(int))
		case reflect.Float32:
			_, _ = fmt.Fprintf(cmd.OutOrStdout(), "%s: %f\n", key, value.(float32))
		case reflect.Float64:
			_, _ = fmt.Fprintf(cmd.OutOrStdout(), "%s: %f\n", key, value.(float64))
		case reflect.Slice, reflect.Array:
			arr := reflect.ValueOf(value)
			_, _ = fmt.Fprintf(cmd.OutOrStdout(), "%s: [", key)
			for i := 0; i < arr.Len(); i++ {
				if i+1 != arr.Len() {
					_, _ = fmt.Fprintf(cmd.OutOrStdout(), "'%v', ", arr.Index(i).Interface())
				} else {
					_, _ = fmt.Fprintf(cmd.OutOrStdout(), "'%v'", arr.Index(i).Interface())
				}
			}
			_, _ = fmt.Fprintf(cmd.OutOrStdout(), "]\n")
		case reflect.Map:
			res2B, _ := json.Marshal(value)
			_, _ = fmt.Fprintf(cmd.OutOrStdout(), "%s: %s\n", key, string(res2B))
		default:
			_, _ = fmt.Fprintf(cmd.OutOrStdout(), "%s: %v\n", key, value)
		}
	}
}
