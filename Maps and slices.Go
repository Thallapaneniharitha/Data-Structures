//1.Maps and Slices 
package main 
import "fmt"
func main() { 
    m_a_p := map[int]string{ 
  
        90: "Dog", 
        91: "Cat", 
        92: "Cow", 
        93: "Bird", 
        94: "Rabbit", 
    } 
  
    // Iterating map using for rang loop 
    for id, pet := range m_a_p { 
  
        fmt.Println(id, pet) 
    } 
} 
$go run main.go
94 Rabbit
90 Dog
91 Cat
92 Cow
93 Bird

//2.Slice example
package main 
import "fmt"
func main() { 
 
    var my_slice_1 = []string{"haritha", "helloworld", "saritha"} 
  
    fmt.Println("My Slice 1:", my_slice_1) 
  
    
    my_slice_2 := []int{12, 45, 67, 56, 43, 34, 45} 
    fmt.Println("My Slice 2:", my_slice_2) 
} 
$go run main.go
My Slice 1: [haritha helloworld saritha]
My Slice 2: [12 45 67 56 43 34 45]