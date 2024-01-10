package algorithm

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates2(nums))
	fmt.Println(nums)
}
