//1.Tree inorder traversal
package main
 
import "fmt"
 
type Node struct {
	val   int
	left  *Node
	right *Node
}
 
func InorderRecursive(root *Node) {
	if root == nil {
		return
	}
 
	InorderRecursive(root.left)
	fmt.Printf("%d ", root.val)
	InorderRecursive(root.right)
}
 
func main() {
	
 
	root := &Node{110, nil, nil}
	root.left = &Node{210, nil, nil}
	root.right = &Node{310, nil, nil}
	root.left.left = &Node{410, nil, nil}
	root.left.right = &Node{510, nil, nil}
	root.right.right = &Node{610, nil, nil}
	root.left.left.left = &Node{710, nil, nil}
 
	fmt.Println("Inorder Traversal - recursive solution : ")
	InorderRecursive(root)
}
Inorder Traversal - recursive solution : 
710 410 210 510 110 310 610 
Program exited.

2.package main
 
import (
	"fmt"
 
	"github.com/emirpasic/gods/stacks/arraystack"
)
 
type Node struct {
	val   int
	left  *Node
	right *Node
}
 
func InorderIterative(root *Node) {
	if root == nil {
		return
	}
 
	temp := root
	stack := arraystack.New()
	for stack.Size() > 0 || temp != nil {
		if temp != nil {
			stack.Push(temp)
			temp = temp.left
		} else {
			obj, _ := stack.Pop()
			temp = obj.(*Node)
			fmt.Printf("%d ", temp.val)
			temp = temp.right
		}
	}
}
 
func main() {
	
	root := &Node{20, nil, nil}
	root.left = &Node{30, nil, nil}
	root.right = &Node{90, nil, nil}
	root.left.left = &Node{40, nil, nil}
	root.left.right = &Node{50, nil, nil}
	root.right.right = &Node{60, nil, nil}
	root.left.left.left = &Node{70, nil, nil}
 
	fmt.Println("Inorder Traversal - iterative solution : ")
	InorderIterative(root)
}
Inorder Traversal - iterative solution : 
70 40 30 50 20 90 60 
Program exited.

3.//pre order
package main

import (
	"fmt"
)

type Node struct {
	val   int
	left  *Node
	right *Node
}

func PreorderRecursive(root *Node) {
	if root != nil {
		fmt.Printf("%d ", root.val)
		PreorderRecursive(root.left)
		PreorderRecursive(root.right)
	}
}

func main() {
	
	root := &Node{10, nil, nil}
	root.left = &Node{20, nil, nil}
	root.right = &Node{30, nil, nil}
	root.left.left = &Node{40, nil, nil}
	root.left.right = &Node{50, nil, nil}
	root.right.right = &Node{60, nil, nil}
	root.left.left.left = &Node{70, nil, nil}

	fmt.Println("Preorder Traversal - Recursive Solution : ")
	PreorderRecursive(root)
}
Preorder Traversal - Recursive Solution : 
10 20 40 70 50 30 60 
Program exited.

4.package main

import (
	"fmt"

	"github.com/emirpasic/gods/stacks/arraystack"
)

type Node struct {
	val   int
	left  *Node
	right *Node
}

func PreorderIterative(root *Node) {
	if root == nil {
		return
	}

	current := root
	stack := arraystack.New()
	stack.Push(current)
	for stack.Size() > 0 {
		temp, _ := stack.Pop()
		current = temp.(*Node)
		fmt.Printf("%d ", current.val)
		if current.right != nil {
			stack.Push(current.right)
		}
		if current.left != nil {
			stack.Push(current.left)
		}
	}
}

func main() {
	
	root := &Node{10, nil, nil}
	root.left = &Node{20, nil, nil}
	root.right = &Node{30, nil, nil}
	root.left.left = &Node{40, nil, nil}
	root.left.right = &Node{50, nil, nil}
	root.right.right = &Node{60, nil, nil}
	root.left.left.left = &Node{70, nil, nil}

	fmt.Println("Preorder Traversal - Iterative Solution : ")
	PreorderIterative(root)
}
Preorder Traversal - Iterative Solution : 
10 20 40 70 50 30 60 
Program exited.
//tarvarsals
package main

import "fmt"

type Node struct {
    data        string
    left, right *Node
}

func main() {
    nodeG := Node{data: "g", left: nil, right: nil}
    nodeF := Node{data: "f", left: &nodeG, right: nil}
    nodeE := Node{data: "e", left: nil, right: nil}
    nodeD := Node{data: "d", left: &nodeE, right: nil}
    nodeC := Node{data: "c", left: nil, right: nil}
    nodeB := Node{data: "b", left: &nodeD, right: &nodeF}
    nodeA := Node{data: "a", left: &nodeB, right: &nodeC}
    fmt.Println("Preorder")
    nodeA.PrintPre()
    fmt.Println("Inorder")
    nodeA.PrintIn()
    fmt.Println("Postorder")
    nodeA.PrintPost()
}

// Preorder (Root, Left, Right)
func (root *Node) PrintPre() {
    fmt.Println(root.data)
    if root.left != nil {
        root.left.PrintPre()
    }
    if root.right != nil {
        root.right.PrintPre()
    }
}

// Inorder (Left, Root, Right)
func (root *Node) PrintIn() {
    if root.left != nil {
        root.left.PrintIn()
    }
    fmt.Println(root.data)
    if root.right != nil {
        root.right.PrintIn()
    }
}

// Postorder (Left, Right, Root)
func (root *Node) PrintPost() {
    if root.left != nil {
        root.left.PrintPost()
    }
    if root.right != nil {
        root.right.PrintPost()
    }
    fmt.Println(root.data)
}

$go run main.go
Preorder
a
b
d
e
f
g
c
Inorder
e
d
b
g
f
a
c
Postorder
e
d
g
f
b
c
a