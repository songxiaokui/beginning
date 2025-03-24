package main

import (
	"fmt"
	"reflect"
	"sync"
)

type IDGenerator struct {
	counter int64
	mutex   sync.Mutex
}

// NewIDGenerator initializes a new IDGenerator
func NewIDGenerator() *IDGenerator {
	return &IDGenerator{counter: 0}
}

// GenerateID generates a new unique ID
func (gen *IDGenerator) GenerateID() int64 {
	gen.mutex.Lock()
	defer gen.mutex.Unlock()
	gen.counter++
	return gen.counter - 1
}
func AssignUniqueNames2(obj interface{}, parentName string, gen *IDGenerator) error {
	v := reflect.ValueOf(obj)
	t := reflect.TypeOf(obj)

	// Ensure the input is a pointer and not nil
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return fmt.Errorf("the input is not a pointer type or is nil")
	}

	// Retrieve the value and type corresponding to the pointer
	v = v.Elem()
	t = t.Elem()

	// Ensure the input is a struct
	if v.Kind() != reflect.Struct {
		return fmt.Errorf("the input object is not a struct")
	}

	// Iterate over the struct fields
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldValue := v.Field(i)

		// Check if the field type is GeomGroup
		if field.Type == reflect.TypeOf(GeomGroup{}) {
			// Generate a unique name, which can combine struct name, field name, and a snowflake algorithm
			uniqueName := fmt.Sprintf("%s_%d", parentName, gen.GenerateID())

			// Set the Name field of GeomGroup
			geomGroup := fieldValue.Interface().(GeomGroup)
			geomGroup.GroupName = uniqueName

			// Set the modified value back
			fieldValue.Set(reflect.ValueOf(geomGroup))
		} else if field.Type.Kind() == reflect.Struct { // Recursively handle nested structs
			err := AssignUniqueNames2(fieldValue.Addr().Interface(), parentName, gen)
			if err != nil {
				return err
			}
		} else if field.Type.Kind() == reflect.Slice && field.Type.Elem() == reflect.TypeOf(GeomGroup{}) { // Handle []GeomGroup slices
			for j := 0; j < fieldValue.Len(); j++ {
				geomGroup := fieldValue.Index(j).Interface().(GeomGroup)
				uniqueName := fmt.Sprintf("%s_%d", parentName, gen.GenerateID())
				geomGroup.GroupName = uniqueName
				fieldValue.Index(j).Set(reflect.ValueOf(geomGroup))
			}
		}
	}

	return nil
}

// 给GeomGroup类型的字段分配唯一名称
func AssignUniqueNames(obj interface{}, parentName string) error {
	// 获取传入对象的反射类型和值
	v := reflect.ValueOf(obj)
	t := reflect.TypeOf(obj)

	// 确保传入的是指针，否则返回错误
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return fmt.Errorf("传入的不是指针类型或为空")
	}

	// 取出指针对应的值和类型
	v = v.Elem()
	t = t.Elem()

	// 确保传入的是结构体
	if v.Kind() != reflect.Struct {
		return fmt.Errorf("传入的对象不是结构体")
	}

	// 遍历结构体字段
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldValue := v.Field(i)

		// 检查字段类型是否为 define.GeomGroup
		if field.Type == reflect.TypeOf(GeomGroup{}) {
			// 生成唯一名称，可以结合结构体名、字段名和 snowflake 算法
			uniqueName := fmt.Sprintf("%s_%s_%d", parentName, field.Name, i)

			// 设置 GeomGroup 的 Name 字段
			geomGroup := fieldValue.Interface().(GeomGroup)
			geomGroup.GroupName = uniqueName

			// 将修改后的值设置回去
			fieldValue.Set(reflect.ValueOf(geomGroup))
		} else if field.Type.Kind() == reflect.Struct { // 递归处理嵌套结构体
			nestedParentName := fmt.Sprintf("%s_%s", parentName, field.Name)
			err := AssignUniqueNames(fieldValue.Addr().Interface(), nestedParentName)
			if err != nil {
				return err
			}
		} else if field.Type.Kind() == reflect.Slice && field.Type.Elem() == reflect.TypeOf(GeomGroup{}) { // 处理 []GeomGroup 切片
			for j := 0; j < fieldValue.Len(); j++ {
				geomGroup := fieldValue.Index(j).Interface().(GeomGroup)
				uniqueName := fmt.Sprintf("%s_%s_%d_%d", parentName, field.Name, i, j)
				geomGroup.GroupName = uniqueName
				fieldValue.Index(j).Set(reflect.ValueOf(geomGroup))
			}
		}
	}

	return nil
}

type GeomGroup struct {
	GroupName string   `json:"group_name" zh:"分组名称"`
	Selected  []string `json:"selected" zh:"选中的几何信息"`
}

type B1 struct {
	C GeomGroup
	D []GeomGroup
}

func main() {
	// 示例结构体
	type Example struct {
		MasterSurface GeomGroup `json:"master_surface"`
		SlaveSurface  GeomGroup `json:"slave_surface"`
		C             B1
	}

	example := &Example{
		C: B1{
			C: GeomGroup{},
			D: []GeomGroup{
				{}, {},
			},
		},
	}

	// 调用函数给 GeomGroup 类型字段赋唯一名称
	var err error
	//err := AssignUniqueNames2(example, "cc", NewIDGenerator())
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("MasterSurface Name: %s\n", example.MasterSurface.GroupName)
		fmt.Printf("SlaveSurface Name: %s\n", example.SlaveSurface.GroupName)
		fmt.Printf("C.C Name: %s\n", example.C.C.GroupName)
		fmt.Printf("C.D[0] Name: %s\n", example.C.D[0].GroupName)
		fmt.Printf("C.D[1] Name: %s\n", example.C.D[1].GroupName)
	}
}
