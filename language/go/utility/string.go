package main

import (
	"errors"
	"fmt"
	"github.com/avast/retry-go/v4"
	"time"
)

type Person struct {
	Name string
	Age  int
}

// Person 类型实现了 Stringer 接口
func (p Person) String() string {
	return fmt.Sprintf("%s is %d years old", p.Name, p.Age)
}

func RetryFunc(f func() error, attempts uint, delay time.Duration) error {
	return retry.Do(
		f,
		retry.Attempts(attempts),
		retry.Delay(delay),
	)
}

func main() {
	p := Person{Name: "John Doe", Age: 30}
	fmt.Printf("This is %s.\n", p)

	err := RetryFunc(func() error {
		return errors.New("err")
	}, 3, time.Second*3)
	fmt.Println(err)
}
