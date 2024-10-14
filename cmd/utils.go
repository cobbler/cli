package cmd

import (
	"encoding/json"
	"fmt"
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

func printDumpVars(blendedData map[string]interface{}) {
	for key, value := range blendedData {
		if value == nil {
			fmt.Printf("%s:\n", key)
			continue
		}
		valueType := reflect.TypeOf(value).Kind()
		switch valueType {
		case reflect.Bool:
			fmt.Printf("%s: %t\n", key, value.(bool))
		case reflect.Int64:
			fmt.Printf("%s: %d\n", key, value.(int64))
		case reflect.Int32:
			fmt.Printf("%s: %d\n", key, value.(int32))
		case reflect.Int16:
			fmt.Printf("%s: %d\n", key, value.(int16))
		case reflect.Int8:
			fmt.Printf("%s: %d\n", key, value.(int8))
		case reflect.Int:
			fmt.Printf("%s: %d\n", key, value.(int))
		case reflect.Float32:
			fmt.Printf("%s: %f\n", key, value.(float32))
		case reflect.Float64:
			fmt.Printf("%s: %f\n", key, value.(float64))
		case reflect.Slice, reflect.Array:
			arr := reflect.ValueOf(value)
			fmt.Printf("%s: [", key)
			for i := 0; i < arr.Len(); i++ {
				if i+1 != arr.Len() {
					fmt.Printf("'%v', ", arr.Index(i).Interface())
				} else {
					fmt.Printf("'%v'", arr.Index(i).Interface())
				}
			}
			fmt.Printf("]\n")
		case reflect.Map:
			res2B, _ := json.Marshal(value)
			fmt.Printf("%s: %s\n", key, string(res2B))
		default:
			fmt.Printf("%s: %s\n", key, value)
		}
	}
}
