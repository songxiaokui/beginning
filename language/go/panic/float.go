package main

import (
	"fmt"
	"github.com/cockroachdb/apd"
	"log"
	"strconv"
	"strings"
)

func RetainFloatingPointPrecision(precision uint32, v float64) float64 {
	ctx := apd.BaseContext
	ctx.Precision = precision
	d, _, err := ctx.NewFromString(fmt.Sprintf("%.*f", precision, v))
	if err != nil {
		log.Printf("input: %v, convert digital string failed, err: %+v", v, err)
		return v
	}
	trimmedString := strings.TrimRight(d.String(), "0")
	if strings.Contains(trimmedString, ".") {
		trimmedString = strings.TrimRight(trimmedString, ".")
	}
	data, err := strconv.ParseFloat(trimmedString, 64)
	if err != nil {
		log.Printf("input: %v, convert digital to float64 failed, err: %+v", v, err)
		return v
	}
	return data
}

func RetainStringPrecision(precision uint32, v string) string {
	ctx := apd.BaseContext
	ctx.Precision = precision
	d, _, err := ctx.NewFromString(v)
	if err != nil {
		log.Printf("input: %v, convert digital string failed, err: %+v", v, err)
		return v
	}
	// Trim unnecessary trailing zeros and decimal point if present
	trimmedString := strings.TrimRight(d.String(), "0")
	if strings.Contains(trimmedString, ".") {
		trimmedString = strings.TrimRight(trimmedString, ".")
	}
	return trimmedString
}

const (
	defaultPrecision uint32 = 6
)

func GetDefaultPrecisionValue(v any) any {
	switch value := v.(type) {
	case string:
		return RetainStringPrecision(defaultPrecision, value)
	case float64:
		return RetainFloatingPointPrecision(defaultPrecision, value)
	default:
		return v // For unsupported types, return the input value as is
	}
}
func main() {

	pl := []string{"0.000510011111", "128.9701", "0.0000000100000000000011", "0.001001111", "1.0001"}
	for _, v := range pl {
		fmt.Println(GetDefaultPrecisionValue(v))
	}
	fmt.Println(GetDefaultPrecisionValue(fmt.Sprintf("%g", 0.000510011111)))

}
