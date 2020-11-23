// 1. Zig Zag matrix in golang
package main
  
import (
    "fmt"
)
  
func zigzag(n int) []int {
    zz := make([]int, n*n)
    i := 0
    n2 := n * 2
    for p := 1; p <= n2; p++ {
        x := p - n
        if x < 0 {
            x = 0
        }
        y := p - 1
        if y > n-1 {
            y = n - 1
        }
        j := n2 - p
        if j > p {
            j = p
        }
        for k := 0; k < j; k++ {
            if p&1 == 0 {
                zz[(x+k)*n+y-k] = i
            } else {
                zz[(y-k)*n+x+k] = i
            }
            i++
        }
    }
  
    return zz
}
  
func main() {
    num := 4
    len := 2
    for i, draw := range zigzag(num) {
        fmt.Printf("%*d ", len, draw)
        if i%num == num-1 {
            fmt.Println("")
        }
    }
}

$go run main.go
0  1  5  6 
 2  4  7 12 
 3  8 11 13 
 9 10 14 15 