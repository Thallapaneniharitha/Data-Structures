//1.Bruteforce
package main

import (
    "fmt"
)

func findElement(arr[10] int, k int) bool {
    var i int
    for i=0; i< 10; i++ {
        if arr[i]==k {
            return true
        }
    }
    return false
}
// main method
func main() {
    var arr = [10]int{1,4,7,8,3,9,2,4,1,8}
    var check bool = findElement(arr,10)
    fmt.Println(check)
    var check2 bool = findElement(arr,9)
    fmt.Println(check2)
}

$go run main.go
false
true
