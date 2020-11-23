//1.string compare
package main
 
import (
  "fmt"
  "strings"
)
 
func main() {
  fmt.Println(strings.Compare("A", "B"))  
  fmt.Println(strings.Compare("B", "A"))  
  fmt.Println(strings.Compare("Japan", "Australia"))  
  fmt.Println(strings.Compare("Australia", "Japan"))  
  fmt.Println(strings.Compare("Germany", "Germany"))  
  fmt.Println(strings.Compare("Germany", "GERMANY"))  
  fmt.Println(strings.Compare("", " ")) 
}
$go run main.go
-1
1
1
-1
0
1
-1


//2.string index
package main 
import (
  "fmt"
  "strings"
)
 
func main() {
  fmt.Println(strings.Index("Australia", "Aus"))
  fmt.Println(strings.Index("Australia", "aus"))
  fmt.Println(strings.Index("Australia", "A"))
  fmt.Println(strings.Index("Australia", "a"))
  fmt.Println(strings.Index("Australia", "Jap"))
  fmt.Println(strings.Index("Japan-124", "-"))
  fmt.Println(strings.Index("Japan-124", ""))
}
$go run main.go
0
-1
0
5
-1
5
0

