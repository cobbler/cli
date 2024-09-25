package cmd

import (
	"fmt"
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
