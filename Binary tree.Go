//1 Binary Tree in Golang
package main
  
import (
    "fmt"
    "os"
    "io"
)
  
type BinaryNode struct {
    left  *BinaryNode
    right *BinaryNode
    data  int64
}
  
type BinaryTree struct {
    root *BinaryNode
}
  
func (t *BinaryTree) insert(data int64) *BinaryTree {
    if t.root == nil {
        t.root = &BinaryNode{data: data, left: nil, right: nil}
    } else {
        t.root.insert(data)
    }
    return t
}
  
func (n *BinaryNode) insert(data int64) {
    if n == nil {
        return
    } else if data <= n.data {
        if n.left == nil {
            n.left = &BinaryNode{data: data, left: nil, right: nil}
        } else {
            n.left.insert(data)
        }
    } else {
        if n.right == nil {
            n.right = &BinaryNode{data: data, left: nil, right: nil}
        } else {
            n.right.insert(data)
        }
    }   
}
  
func print(w io.Writer, node *BinaryNode, ns int, ch rune) {
    if node == nil {
        return
    }
  
    for i := 0; i < ns; i++ {
        fmt.Fprint(w, " ")
    }
    fmt.Fprintf(w, "%c:%v\n", ch, node.data)
    print(w, node.left, ns+2, 'L')
    print(w, node.right, ns+2, 'R')
}
  
func main() {
    tree := &BinaryTree{}
    tree.insert(100).
        insert(-20).
        insert(-50).
        insert(-15).
        insert(-60).
        insert(55).
        insert(66).
        insert(55).
        insert(80).
        insert(20).
        insert(57).
        insert(-10)
    print(os.Stdout, tree.root, 0, 'M')
}
$go run main.go
M:100
  L:-20
    L:-50
      L:-60
    R:-15
      R:55
        L:55
          L:20
            L:-10
        R:66
          L:57
          R:80

Program exited.
