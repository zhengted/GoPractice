package Tree

import (
	"bytes"
	"fmt"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// 		e.g. root := Tree.Node{Val: 1}
func (ts *Node) String() string {
	var (
		buf bytes.Buffer
		arr []int
	)
	arr = PreOrder(ts, arr)
	//fmt.Println(arr)
	buf.WriteByte('{')
	for _, v := range arr {
		if buf.Len() > len("{") {
			buf.WriteByte(' ')
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte('}')
	return buf.String()
}

func PreOrder(root *Node, Org []int) []int {
	if root != nil {
		Org = append(Org, (*root).Val)
		Org = PreOrder(root.Left, Org)
		Org = PreOrder(root.Right, Org)
	}
	return Org
}
