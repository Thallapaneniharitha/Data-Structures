//1.Hashing in golang
package main
import (
        "fmt"
        "hash/fnv"
)

func hash(s string) uint32 {
        h := fnv.New32a()
        h.Write([]byte(s))
        return h.Sum32()
}

func main() {
        fmt.Println(hash("Helloworld"))
        fmt.Println(hash("Helloworld"))
}
$go run main.go
1698852737
378916989