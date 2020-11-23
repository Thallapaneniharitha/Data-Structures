//1.Linear search
package main  
import "fmt"

func linearsearch(datalist []int, key int) bool {
    for _, item := range datalist {
        if item == key {
            return true
        }
    }
    return false
}  
func main() {
    items := []int{45,78,46,56,78,66,45,86,96,78,245,400}
    fmt.Println(linearsearch(items,59))
}
$go run main.go
false

//2.program
package main

import "fmt"

func linearSearch(numbers []int, item int) int {

	if numbers != nil && len(numbers) > 0 {

		for i := 0; i< len(numbers); i++ {

			if numbers[i] == item {
				return numbers[i]
			}

		}

	}

	return -1
}

func main() {

	numbers := []int{5, 3, 4, 2, 1, 6, 7, 8, 10, 9}

	result := linearSearch(numbers, 10)

	if result != -1 {
		fmt.Println("Item",result,"is found!")
	}
}
$go run main.go
Item 10 is found!
