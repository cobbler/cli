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
// --autoinstall-meta) into the map[string]interface{} expected by the XML-RPC API. As a fallback for values that
// need to carry a list or dict rather than a flat string (e.g. Agama's agama_modules or cloud-init's per-module
// settings), any value that looks like a JSON array or object (starts with '[' or '{') is decoded as JSON; if it
// fails to parse, the original string is used unchanged so existing flat-string usage is unaffected.
func convertMapStringToMapInterface(m map[string]string) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range m {
		trimmed := strings.TrimSpace(v)
		if len(trimmed) > 0 && (trimmed[0] == '[' || trimmed[0] == '{') {
			var decoded interface{}
			if err := json.Unmarshal([]byte(trimmed), &decoded); err == nil {
				result[k] = decoded
				continue
			}
		}
		result[k] = v
	}
	return result
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
			_, _ = fmt.Fprintf(cmd.OutOrStdout(), "%s: %s\n", key, value)
		}
	}
}
