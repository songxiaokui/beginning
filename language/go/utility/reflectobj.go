package main

import (
	"fmt"
	"reflect"
)

type GeomElement struct {
	ID int
}

type GeomElementGroup []GeomElement

type LocationValue struct {
	Location GeomElementGroup
	Value    string
}

type BoundaryCondition struct {
	JobName      string
	Pressure     LocationValue
	FixedSupport GeomElementGroup
}

// Recursively collects GeomElementGroup from any nested structure
func collectGeomElementsRecursive(v reflect.Value) []GeomElement {
	var elements []GeomElement

	// 遍历结构体的所有字段
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		// 检查字段的类型
		switch field.Kind() {
		case reflect.Struct:
			// 递归处理嵌套结构体
			elements = append(elements, collectGeomElementsRecursive(field)...)
		case reflect.Slice:
			// 检查切片元素的类型是否为GeomElement
			if field.Type().Elem() == reflect.TypeOf(GeomElement{}) {
				for j := 0; j < field.Len(); j++ {
					elements = append(elements, field.Index(j).Interface().(GeomElement))
				}
			}
		}
	}

	return elements
}

func collectGeomElements(data interface{}) []GeomElement {
	// 获取变量的反射Value对象
	v := reflect.ValueOf(data)
	// 检查是否为结构体
	if v.Kind() == reflect.Struct {
		return collectGeomElementsRecursive(v)
	}
	// 如果传入的不是结构体，直接返回空切片
	return []GeomElement{}
}

func main() {
	bc := BoundaryCondition{
		JobName: "ExampleJob",
		Pressure: LocationValue{
			Location: GeomElementGroup{{ID: 1}, {ID: 2}},
			Value:    "100kPa",
		},
		FixedSupport: GeomElementGroup{{ID: 3}},
	}

	elements := collectGeomElements(bc)
	fmt.Println("Collected GeomElements:", elements)
}
