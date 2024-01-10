package problem

import (
	"fmt"
	"testing"
)

func Test_rotate(t *testing.T) {
	l4 := []int{1, 2, 3, 4, 5, 6, 7}
	k2 := 3
	rotate3(l4, k2)
	// [5,6,7,1,2,3,4]
	fmt.Println(l4)
}
