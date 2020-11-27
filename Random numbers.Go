//1.Random number
package main

	import "time"
import "fmt"
import "math/rand"

	func main() {
// For example, rand.Intn returns a random int n, 0 <= n < 100.
	    fmt.Print(rand.Intn(100), ",")
    fmt.Print(rand.Intn(100))
    fmt.Println()
// rand.Float64 returns a float64 f, 0.0 <= f < 1.0.
	    fmt.Println(rand.Float64())
// This can be used to generate random floats in other ranges, for example 5.0 <= f' < 10.0.
	    fmt.Print((rand.Float64()*5)+5, ",")
    fmt.Print((rand.Float64() * 5) + 5)
    fmt.Println()

	    
//The default number generator is deterministic, so it’ll produce the same sequence of numbers each time by default. To produce varying sequences, give it a seed that changes. Note that this is not safe to use for random numbers you intend to be secret, use crypto/rand for those.

s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)

/// Call the resulting rand.Rand just like the functions on the rand package.
	    fmt.Print(r1.Intn(100), ",")
    fmt.Print(r1.Intn(100))
    fmt.Println()

//// If you seed a source with the same number, it produces the same sequence of random numbers.
	    s2 := rand.NewSource(42)
    r2 := rand.New(s2)
    fmt.Print(r2.Intn(100), ",")
    fmt.Print(r2.Intn(100))
    fmt.Println()
    s3 := rand.NewSource(42)
    r3 := rand.New(s3)
    fmt.Print(r3.Intn(100), ",")
    fmt.Print(r3.Intn(100))
}
81,87
0.6645600532184904
7.1885709359349015,7.123187485356329
0,28
5,87
5,87
// Random Numbers in Golang 
package main 
  
import ( 
    "fmt"
    "math/rand"
    "time"
) 
  
func main() { 
  
     
    fmt.Println(rand.Intn(200)) 
    fmt.Println(rand.Intn(200)) 
    fmt.Println(rand.Intn(200)) 
    fmt.Println() 
  
    fmt.Println(rand.Float64()) 
  
    
    fmt.Println((rand.Float64() * 8) + 8) 
    fmt.Println((rand.Float64() * 8) + 8) 
    fmt.Println() 
  
    
    x1 := rand.NewSource(time.Now().UnixNano()) 
    y1 := rand.New(x1) 
      
    fmt.Println(y1.Intn(200)) 
    fmt.Println(y1.Intn(200)) 
    fmt.Println() 
  
    x2 := rand.NewSource(55) 
    y2 := rand.New(x2) 
    fmt.Println(y2.Intn(200)) 
    fmt.Println(y2.Intn(200)) 
    fmt.Println() 
      
    x3 := rand.NewSource(5) 
    y3 := rand.New(x3) 
    fmt.Println(y3.Intn(200)) 
    fmt.Println(y3.Intn(200)) 
}
$go run main.go 
81
87
47

0.4377141871869802
11.397099976570125
13.494584582936875

0
128

112
164

26
36