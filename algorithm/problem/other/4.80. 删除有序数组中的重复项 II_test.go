package other

import (
	"fmt"
	"testing"
)

func Test_removeDuplicatesII(t *testing.T) {
	l3 := []int{1, 1, 1, 2, 2, 3}
	// l3 := []int{1, 1, 1}
	fmt.Println(removeDuplicatesII2(l3))
	fmt.Println(l3)
}
