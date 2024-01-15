package linklist

type LinkNode struct {
	Val  int
	Next *LinkNode
}

func (ln *LinkNode) AddElem(node *LinkNode) *LinkNode {
	ln.Next = node
	return ln.Next
}

func NewLinkNode(value int) *LinkNode {
	return &LinkNode{
		Next: nil,
		Val:  value,
	}
}

func NewDefaultLinkNode() *LinkNode {
	lk := NewLinkNode(1)
	lk.AddElem(NewLinkNode(2)).
		AddElem(NewLinkNode(3)).
		AddElem(NewLinkNode(4)).
		AddElem(NewLinkNode(5)).
		AddElem(NewLinkNode(6)).AddElem(NewLinkNode(7))
	return lk
}
