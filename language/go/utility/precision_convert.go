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
	formattedStr := d.String()
	if decimalPos := strings.Index(formattedStr, "."); decimalPos != -1 {
		// 只去除小数点后的尾部零
		formattedStr = strings.TrimRight(formattedStr, "0")
		formattedStr = strings.TrimRight(formattedStr, ".")
	}

	typedVal := reflect.New(reflect.TypeOf(pc.d)).Elem()
	switch typedVal.Kind() {
	case reflect.String:
		typedVal.SetString(formattedStr)
	case reflect.Float64, reflect.Float32:
		if floatVal, err := strconv.ParseFloat(formattedStr, 64); err == nil {

			typedVal.SetFloat(floatVal)
		}
	default:
		return pc.d
	}

	return typedVal.Interface().(T)
}
func main() {
	// 示例使用
	p := NewPrecisionConversion(189.52408281652598, WithPrecision[float64](3)).Convert()
	a := fmt.Sprintf("%v", NewPrecisionConversion("189.52408281652598", WithPrecision[string](3)).Convert())
	fmt.Println(a)
	fmt.Printf("%v, %T\n", p, p)
	p1 := NewPrecisionConversion("2.11132").Convert()
	fmt.Printf("%v, %T\n", p1, p1)
	p2 := NewPrecisionConversion(int(1)).Convert()
	fmt.Printf("%v, %T\n", p2, p2)

	fmt.Println(NewPrecisionConversion("189.52408281652598", WithPrecision[string](3)).Convert())
	fmt.Println(NewPrecisionConversion(189.52408281652598, WithPrecision[float64](3)).Convert())
	fmt.Println(NewPrecisionConversion(0.52408281652598, WithPrecision[float64](3)).Convert())
	fmt.Println(NewPrecisionConversion(0.00012, WithPrecision[float64](3)).Convert())
	//fmt.Println(NewPrecisionConversion(100.00012, WithPrecision[float64](3)).Convert())
	//fmt.Println(NewPrecisionConversion(0.15034, WithPrecision[float64](3)).Convert())
	//fmt.Println(NewPrecisionConversion(10000.1234, WithPrecision[float64](3)).Convert())
	//fmt.Println(NewPrecisionConversion(10.1234, WithPrecision[float64](3)).Convert())
	//fmt.Println(NewPrecisionConversion(10.00014, WithPrecision[float64](3)).Convert())
	//fmt.Println(NewPrecisionConversion(101234.05014, WithPrecision[float64](3)).Convert())

	fmt.Println("--")
	fmt.Println(strings.Index("1000", "."))
	fmt.Println(strings.TrimRight("1000", "0"))
	fmt.Println(strings.TrimRight("1000.0", "."))
	fmt.Println(strings.TrimRight("1000.0101", "."))
}
