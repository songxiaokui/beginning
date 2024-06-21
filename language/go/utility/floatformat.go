package main

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"os"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, errors.Wrapf(err, "The file path is incorrect err : %+v", err)
}

type m1 struct {
	N int
}

func NewM1(i int) m1 {
	return m1{
		N: i,
	}
}
func main() {
	//l := make([]m1, 0)
	//for i := 0; i < 10; i++ {
	//	l = append(l, NewM1(i))
	//}
	//
	//for i := range l {
	//	fmt.Println(i)
	//}

	mp := make(map[string]interface{})
	a, e := json.Marshal(mp)
	fmt.Println(string(a), e)
	if string(a) != "" {
		fmt.Println(string(a))
	}
}
