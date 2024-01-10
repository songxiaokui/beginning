package algorithm

import (
	"fmt"
	"testing"
)

func Test_removeElement(t *testing.T) {
	l1 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	target := 2
	fmt.Println(removeElement(l1, target))
}
