package linklist

import (
	"fmt"
	"testing"
)

func TestGetMedianElement(t *testing.T) {
	lk := NewDefaultLinkNode()
	n1 := GetMedianElement(lk)
	fmt.Println("---", n1.Val)

	n2 := GetMedianElementII(lk, nil)
	fmt.Println("---", n2.Val)
}
