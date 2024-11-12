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

func convertMapStringToMapInterface(m map[string]string) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range m {
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
			fmt.Fprintf(cmd.OutOrStdout(), "%s:\n", key)
			continue
		}
		valueType := reflect.TypeOf(value).Kind()
		switch valueType {
		case reflect.Bool:
			fmt.Fprintf(cmd.OutOrStdout(), "%s: %t\n", key, value.(bool))
		case reflect.Int64:
			fmt.Fprintf(cmd.OutOrStdout(), "%s: %d\n", key, value.(int64))
		case reflect.Int32:
			fmt.Fprintf(cmd.OutOrStdout(), "%s: %d\n", key, value.(int32))
		case reflect.Int16:
			fmt.Fprintf(cmd.OutOrStdout(), "%s: %d\n", key, value.(int16))
		case reflect.Int8:
			fmt.Fprintf(cmd.OutOrStdout(), "%s: %d\n", key, value.(int8))
		case reflect.Int:
			fmt.Fprintf(cmd.OutOrStdout(), "%s: %d\n", key, value.(int))
		case reflect.Float32:
			fmt.Fprintf(cmd.OutOrStdout(), "%s: %f\n", key, value.(float32))
		case reflect.Float64:
			fmt.Fprintf(cmd.OutOrStdout(), "%s: %f\n", key, value.(float64))
		case reflect.Slice, reflect.Array:
			arr := reflect.ValueOf(value)
			fmt.Fprintf(cmd.OutOrStdout(), "%s: [", key)
			for i := 0; i < arr.Len(); i++ {
				if i+1 != arr.Len() {
					fmt.Fprintf(cmd.OutOrStdout(), "'%v', ", arr.Index(i).Interface())
				} else {
					fmt.Fprintf(cmd.OutOrStdout(), "'%v'", arr.Index(i).Interface())
				}
			}
			fmt.Fprintf(cmd.OutOrStdout(), "]\n")
		case reflect.Map:
			res2B, _ := json.Marshal(value)
			fmt.Fprintf(cmd.OutOrStdout(), "%s: %s\n", key, string(res2B))
		default:
			fmt.Fprintf(cmd.OutOrStdout(), "%s: %s\n", key, value)
		}
	}
}
