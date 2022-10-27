package main

import (
	"container/heap"
	"fmt"
	"math/rand"
)

type LNode struct {
	Data interface{}
	Next *LNode
}

func CreateNode(node *LNode) {
	cur := node
	for i := 1; i < 10; i++ {
		cur.Next = &LNode{}
		cur.Next.Data = rand.Intn(4)
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

func RemoveDup(head *LNode) {
	if head == nil || head.Next == nil {
		return
	}
	outerCur := head.Next
	var innerCur *LNode
	var innerPre *LNode
	for ; outerCur != nil; outerCur = outerCur.Next {
		for innerCur, innerPre = outerCur.Next, outerCur; innerCur != nil; {
			if outerCur.Data == innerCur.Data {
				innerPre.Next = innerCur.Next
				innerCur = innerCur.Next
			} else {
				innerPre = innerCur
				innerCur = innerCur.Next
			}
		}
	}
}

func removeDupRecursionChild(head *LNode) *LNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pointer *LNode
	cur := head
	head.Next = removeDupRecursionChild(head.Next)
	pointer = head.Next
	for pointer != nil {
		if head.Data == pointer.Data {
			cur.Next = pointer.Next
			pointer = cur.Next
		} else {
			pointer = pointer.Next
			cur = cur.Next
		}
	}
	return head
}

func RemoveDupRecursion(head *LNode) {
	if head == nil {
		return
	}
	head.Next = removeDupRecursionChild(head.Next)
}

func main() {
	head := &LNode{}
	fmt.Println("删除重复结点")
	CreateNode(head)
	PrintNode("删除前的链表", head)
	RemoveDup(head)
	PrintNode("删除后的链表", head)
}
