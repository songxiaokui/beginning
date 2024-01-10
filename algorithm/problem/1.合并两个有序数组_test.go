package problem

import (
	"fmt"
	"testing"
)

func Test_merge2(t *testing.T) {
	l1 := []int{1, 2, 3, 0, 0, 0}
	l2 := []int{2, 5, 6}
	//merge2(l1, 3, l2, 3)
	merge3(l1, 3, l2, 3)
	fmt.Println(l1)
}
