package main

import (
	"fmt"
	"github.com/cockroachdb/apd"
	"reflect"
	"strconv"
	"strings"
)

const (
	defaultPrecision uint32 = 6
)

// PrecisionConversion 精度转换
type PrecisionConversion[T any] struct {
	d T
	p uint32
}

type PrecisionConversionOptions[T any] func(conversion *PrecisionConversion[T])

func WithPrecision[T any](n uint32) PrecisionConversionOptions[T] {
	return func(conversion *PrecisionConversion[T]) {
		conversion.p = n
	}
}

func NewPrecisionConversion[T any](data T, opts ...PrecisionConversionOptions[T]) *PrecisionConversion[T] {
	pc := &PrecisionConversion[T]{
		d: data,
		p: defaultPrecision,
	}
	for _, fn := range opts {
		fn(pc)
	}
	return pc
}

func (pc *PrecisionConversion[T]) Convert() T {
	ctx := apd.BaseContext
	ctx.Precision = pc.p
	d, _, err := ctx.NewFromString(fmt.Sprintf("%v", pc.d))
	if err != nil {
		return pc.d
	}
	trimmedString := strings.TrimRight(d.String(), "0")
	if strings.Contains(trimmedString, ".") {
		trimmedString = strings.TrimRight(trimmedString, ".")
	}

	typedVal := reflect.New(reflect.TypeOf(pc.d)).Elem()
	switch typedVal.Kind() {
	case reflect.String:
		typedVal.SetString(trimmedString)
	case reflect.Float64, reflect.Float32:
		if floatVal, err := strconv.ParseFloat(trimmedString, 64); err == nil {

			typedVal.SetFloat(floatVal)
		}
	default:
		return pc.d
	}

	return typedVal.Interface().(T)
}
func main() {
	// 示例使用
	p := NewPrecisionConversion(123.456).Convert()
	fmt.Printf("%v, %T\n", p, p)
	p1 := NewPrecisionConversion("2.11132").Convert()
	fmt.Printf("%v, %T\n", p1, p1)
	p2 := NewPrecisionConversion(int(1)).Convert()
	fmt.Printf("%v, %T\n", p2, p2)
}
