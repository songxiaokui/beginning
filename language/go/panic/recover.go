package main

import (
	"errors"
	"fmt"
)

func checkErr1(err error) {
	if err != nil {
		panic(err)
	}
}

func doThing() (dataPath *string, err error) {
	if dataPath == nil {
		dataPath = new(string)
	}
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)

			return
		}
		fmt.Println("44444")
		return
	}()

	*dataPath = "sxk"
	checkErr1(errors.New("报错了"))
	return
}

func main() {
	a, b := doThing()
	fmt.Println(*a, b)
}
