package main

import (
	"errors"
	"fmt"
	"strings"
)

func formatSlice(slice []int) string {
	var sb strings.Builder
	sb.WriteString("[")
	for i, v := range slice {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprintf("%d", v))
	}
	sb.WriteString("]")
	return sb.String()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func do() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("报错了", err)
		}
		fmt.Println("任务结束")
	}()

	fmt.Println("aaaa")
	checkErr(nil)

	checkErr(errors.New("到底了"))
}

func main() {
	do()
}
