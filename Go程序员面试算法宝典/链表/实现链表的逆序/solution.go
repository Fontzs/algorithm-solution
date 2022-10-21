// 三指针
package main

import (
	"fmt"
)

type LNode struct {
	Data interface{}
	Next *LNode
}

func CreateNode(node *LNode, max int) {
	cur := node
	for i := 1; i < max; i++ {
		cur.Next = &LNode{}
		cur.Next.Data = i
		cur = cur.Next
	}
}

func PrintNode(info string, node *LNode) {
	fmt.Println(info)
	for cur := node.Next; cur != nil; cur = cur.Next {
		fmt.Print(cur.Data, " ")
	}
	fmt.Println()
}

// O(n) O(1)
func Reverse(node *LNode) {
	if node == nil || node.Next == nil {
		return
	}
	var pre *LNode
	var next *LNode
	cur := node.Next

	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	node.Next = pre
}

// O(n)
func RecursiveReverseChild(node *LNode) *LNode {
	if node == nil || node.Next == nil {
		return node
	}
	newHead := RecursiveReverseChild(node.Next)
	node.Next.Next = node
	node.Next = nil
	return newHead
}
func RecursiveReverse(node *LNode) {
	firstNode := node.Next
	//递归调用
	newHead := RecursiveReverseChild(firstNode)
	node.Next = newHead
}

func main() {

	head := &LNode{}
	fmt.Print("三指针")
	CreateNode(head, 8)
	PrintNode("逆序前", head)
	Reverse(head)
	PrintNode("逆序后", head)

}
