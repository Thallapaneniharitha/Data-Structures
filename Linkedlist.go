//1.Linked list
package main 
import "fmt" 
type Node struct {
    prev *Node
    next *Node
    key  interface{}
}
 
type List struct {
    head *Node
    tail *Node
}
 
func (L *List) Insert(key interface{}) {
    list := &Node{
        next: L.head,
        key:  key,
    }
    if L.head != nil {
        L.head.prev = list
    }
    L.head = list
 
    l := L.head
    for l.next != nil {
        l = l.next
    }
    L.tail = l
}
 
func (l *List) Display() {
    list := l.head
    for list != nil {
        fmt.Printf("%+v ->", list.key)
        list = list.next
    }
    fmt.Println()
}
 
func Display(list *Node) {
    for list != nil {
        fmt.Printf("%v ->", list.key)
        list = list.next
    }
    fmt.Println()
}
 
func ShowBackwards(list *Node) {
    for list != nil {
        fmt.Printf("%v <-", list.key)
        list = list.prev
    }
    fmt.Println()
}
 
func (l *List) Reverse() {
    curr := l.head
    var prev *Node
    l.tail = l.head
 
    for curr != nil {
        next := curr.next
        curr.next = prev
        prev = curr
        curr = next
    }
    l.head = prev
    Display(l.head)
}
 
func main() {
    link := List{}
    link.Insert(5)
    link.Insert(9)
    link.Insert(13)
    link.Insert(22)
    link.Insert(28)
    link.Insert(36)
     
    fmt.Println("\n==============================\n")
    fmt.Printf("Head: %v\n", link.head.key)
    fmt.Printf("Tail: %v\n", link.tail.key)
    link.Display()
    fmt.Println("\n==============================\n")
    fmt.Printf("head: %v\n", link.head.key)
    fmt.Printf("tail: %v\n", link.tail.key)
    link.Reverse()
    fmt.Println("\n==============================\n")
}

$go run main.go
==============================

Head: 36
Tail: 5
36 ->28 ->22 ->13 ->9 ->5 ->

==============================

head: 36
tail: 5
5 ->9 ->13 ->22 ->28 ->36 ->

==============================


//1.Singly Linked List
package main

import "fmt"

type ele struct {
    name string
    next *ele
}

type singleList struct {
    len  int
    head *ele
}

func initList() *singleList {
    return &singleList{}
}

func (s *singleList) AddFront(name string) {
    ele := &ele{
        name: name,
    }
    if s.head == nil {
        s.head = ele
    } else {
        ele.next = s.head
        s.head = ele
    }
    s.len++
    return
}

func (s *singleList) AddBack(name string) {
    ele := &ele{
        name: name,
    }
    if s.head == nil {
        s.head = ele
    } else {
        current := s.head
        for current.next != nil {
            current = current.next
        }
        current.next = ele
    }
    s.len++
    return
}

func (s *singleList) RemoveFront() error {
    if s.head == nil {
        return fmt.Errorf("List is empty")
    }
    s.head = s.head.next
    s.len--
    return nil
}

func (s *singleList) RemoveBack() error {
    if s.head == nil {
        return fmt.Errorf("removeBack: List is empty")
    }
    var prev *ele
    current := s.head
    for current.next != nil {
        prev = current
        current = current.next
    }
    if prev != nil {
        prev.next = nil
    } else {
        s.head = nil
    }
    s.len--
    return nil
}

func (s *singleList) Front() (string, error) {
    if s.head == nil {
        return "", fmt.Errorf("Single List is empty")
    }
    return s.head.name, nil
}

func (s *singleList) Size() int {
    return s.len
}

func (s *singleList) Traverse() error {
    if s.head == nil {
        return fmt.Errorf("TranverseError: List is empty")
    }
    current := s.head
    for current != nil {
        fmt.Println(current.name)
        current = current.next
    }
    return nil
}

func main() {
     singleList := initList()
    fmt.Printf("AddFront: A\n")
    singleList.AddFront("A")
    fmt.Printf("AddFront: B\n")
    singleList.AddFront("B")
    fmt.Printf("AddBack: C\n")
    singleList.AddBack("C")

    fmt.Printf("Size: %d\n", singleList.Size())
   
    err := singleList.Traverse()
    if err != nil {
        fmt.Println(err.Error())
    }
    
    fmt.Printf("RemoveFront\n")
    err = singleList.RemoveFront()
    if err != nil {
        fmt.Printf("RemoveFront Error: %s\n", err.Error())
    }
    
    fmt.Printf("RemoveBack\n")
    err = singleList.RemoveBack()
    if err != nil {
        fmt.Printf("RemoveBack Error: %s\n", err.Error())
    }
    
    fmt.Printf("RemoveBack\n")
    err = singleList.RemoveBack()
    if err != nil {
        fmt.Printf("RemoveBack Error: %s\n", err.Error())
    }
    
    fmt.Printf("RemoveBack\n")
    err = singleList.RemoveBack()
    if err != nil {
        fmt.Printf("RemoveBack Error: %s\n", err.Error())
    }
    
    err = singleList.Traverse()
    if err != nil {
        fmt.Println(err.Error())
    }
    
   fmt.Printf("Size: %d\n", singleList.Size())
}
$go run main.go
AddFront: A
AddFront: B
AddBack: C
Size: 3
B
A
C
RemoveFront
RemoveBack
RemoveBack
RemoveBack
RemoveBack Error: removeBack: List is empty
TranverseError: List is empty
Size: 0

//2.Double Linkedlist
package main
import (
    "fmt"
    "errors"
    "strings"
)
type Value struct {
    Name string
    MilesAway int
}
type Node struct {
    Value               
    next, prev  *Node
}
type List struct {
    head, tail *Node
}
func (l *List) First() *Node {
    return l.head
}
func (n *Node) Next() *Node {
    return n.next
}
func (n *Node) Prev() *Node {
    return n.prev
}
func (l *List) Push(v Value) *List {
    n := &Node{Value: v}
    if l.head == nil {
        l.head = n      // First node
    } else {
        l.tail.next = n 
        n.prev = l.tail 
    }
    l.tail = n          
    return l
}
func (l *List) Find(name string) *Node {
    found := false
    var ret *Node = nil
    for n := l.First(); n != nil && !found; n = n.Next() {
        if n.Value.Name == name {
            found = true
            ret = n
        }
    }
    return ret
}
func (l *List) Delete(name string) bool {
    success := false
    node2del := l.Find(name)
    if node2del != nil {
        fmt.Println("Delete - FOUND: ", name)
        prev_node := node2del.prev
        next_node := node2del.next
        prev_node.next = node2del.next
        next_node.prev = node2del.prev
        success = true
    }
    return success
}
var errEmpty = errors.New("ERROR - List is empty")

func (l *List) Pop() (v Value, err error) {
    if l.tail == nil {
        err = errEmpty
    } else {
        v = l.tail.Value
        l.tail = l.tail.prev
        if l.tail == nil {
            l.head = nil
        }
    }
    return v, err
}

func main() {
    dashes := strings.Repeat("-", 50)
    l := new(List)  // Create Doubly Linked List

    l.Push(Value{Name: "Atlanta", MilesAway: 0})
    l.Push(Value{Name: "Las Vegas", MilesAway: 1961})
    l.Push(Value{Name: "New York", MilesAway: 881})

    processed := make(map[*Node]bool)

    fmt.Println("First time through list...")
    for n := l.First(); n != nil; n = n.Next() {
        fmt.Printf("%v\n", n.Value)
        if processed[n] {
            fmt.Printf("%s as been processed\n", n.Value)
        }
        processed[n] = true
    }
    fmt.Println(dashes)
    fmt.Println("Second time through list...")
    for n := l.First(); n != nil; n = n.Next() {
        fmt.Printf("%v", n.Value)
        if processed[n] {
            fmt.Println(" has been processed")
        } else { fmt.Println() }
        processed[n] = true
    }

    fmt.Println(dashes)
    var found_node *Node
    city_to_find := "New York"
    found_node = l.Find(city_to_find)
    if found_node == nil {
        fmt.Printf("NOT FOUND: %v\n", city_to_find)
    } else {
        fmt.Printf("FOUND: %v\n", city_to_find)
    }

    city_to_find = "Chicago"
    found_node = l.Find(city_to_find)
    if found_node == nil {
        fmt.Printf("NOT FOUND: %v\n", city_to_find)
    } else {
        fmt.Printf("FOUND: %v\n", city_to_find)
    }

    fmt.Println(dashes)
    city_to_remove := "Las Vegas"
    successfully_removed_city := l.Delete(city_to_remove)
    if successfully_removed_city {
        fmt.Printf("REMOVED: %v\n", city_to_remove)
    } else {
        fmt.Printf("DID NOT REMOVE: %v\n", city_to_remove)
    }

    city_to_remove = "Chicago"
    successfully_removed_city = l.Delete(city_to_remove)
    if successfully_removed_city {
        fmt.Printf("REMOVED: %v\n", city_to_remove)
    } else {
        fmt.Printf("DID NOT REMOVE: %v\n", city_to_remove)
    }

    fmt.Println(dashes)
    fmt.Println("* Pop each value off list...")
    for v, err := l.Pop(); err == nil; v, err = l.Pop() {
        fmt.Printf("%v\n", v)
    }
    fmt.Println(l.Pop()) 
}

$go run main.go
First time through list...
{Atlanta 0}
{Las Vegas 1961}
{New York 881}
--------------------------------------------------
Second time through list...
{Atlanta 0} has been processed
{Las Vegas 1961} has been processed
{New York 881} has been processed
--------------------------------------------------
FOUND: New York
NOT FOUND: Chicago
--------------------------------------------------
Delete - FOUND:  Las Vegas
REMOVED: Las Vegas
DID NOT REMOVE: Chicago
--------------------------------------------------
* Pop each value off list...
{New York 881}
{Atlanta 0}
{ 0} ERROR - List is empty


//3.Circular list
package main
import "fmt"
type node struct {
    val  int
    next *node
    prev *node
}
type circularlinkedlist struct {
    head *node
    tail *node
}

func newNode(val int) *node {
    n := &node{}
    n.val = val
    n.next = nil
    n.prev = nil
    return n
}
func (ll *circularlinkedlist) addAtBeg(val int) {
    n := newNode(val)
    n.next = ll.head
    ll.head = n
    if ll.tail == nil {
        ll.tail = n
    }
    n.prev = ll.tail
    ll.tail.next = n
}
func (ll *circularlinkedlist) addAtEnd(val int) {
    n := newNode(val)
    if ll.head == nil {
        ll.head = n
        ll.tail = n
        return
    }
    if ll.head == ll.tail {
        ll.head.next = n
        ll.head.prev = n
        n.next = ll.head
        n.prev = ll.head
        ll.tail = n
        return
    }
    cur := ll.head
    for ; cur.next != ll.tail; cur = cur.next {
    }
    cur.next.next = n
    n.prev = cur.next
    n.next = ll.head
    ll.tail = n
}
func (ll *circularlinkedlist) delAtBeg() int {
    if ll.head == nil {
        return -1
    }
    if ll.head == ll.tail {
        value := ll.head
        ll.head = nil
        ll.tail = nil
        return value.val
    }
    cur := ll.head
    ll.head = cur.next
    if ll.head != nil {
        ll.head.prev = ll.tail
        ll.tail.next = ll.head
    }
    return cur.val
}
func (ll *circularlinkedlist) delAtEnd() int {
   
    if ll.head == nil {
        return -1
    }
    
    if ll.head == ll.tail {
        return ll.delAtBeg()
    }
    
    cur := ll.head
    for ; cur.next != ll.tail; cur = cur.next {
    }
    retval := cur.next.val
    cur.next = ll.head
    ll.tail = cur
    ll.head.prev = cur
    return retval
}
func (ll *circularlinkedlist) count() int {
    var ctr int = 0
    for cur := ll.head; cur != ll.tail; cur = cur.next {
        ctr++
    }
    return ctr
}
func (ll *circularlinkedlist) reverse() {
    var prev, next *node
    if ll.head == ll.tail {
        return
    }
    cur := ll.head.next
    ll.tail = ll.head
    if ll.head.next != nil {
        for cur != ll.head {
            next = cur.next
            cur.next = prev
            cur.prev = next
            prev = cur
            cur = next
        }
    }
    ll.head = prev
    cur.next = ll.tail.next
    ll.tail.next.next = ll.tail
    ll.tail.next = ll.head
    ll.head.prev = ll.tail
}
func (ll *circularlinkedlist) display() {
    if ll.head != nil {
        fmt.Print(ll.head.val, " ")
        if ll.head.next != nil {
            for cur := ll.head.next; cur != ll.head; cur = cur.next {
                fmt.Print(cur.val, " ")
            }
        }
    }
    fmt.Print("
")
}
func (ll *circularlinkedlist) displayReverse() {
    if ll.head == nil {
        return
    }
    var cur *node
    for cur = ll.head; cur.next != ll.tail; cur = cur.next {
    }
    for ; cur != ll.tail; cur = cur.prev {
        fmt.Print(cur.val, " ")
    }
    fmt.Print("
")
}
func main() {
    ll := circularlinkedlist{}
    ll.addAtBeg(10)
    ll.addAtEnd(20)
    ll.reverse()
    ll.display()
    fmt.Print("Deleting an element at the beginning : ", ll.delAtBeg(), "
")
    fmt.Print("Deleting an element at the end : ", ll.delAtEnd(), "
")
    ll.display()
}

