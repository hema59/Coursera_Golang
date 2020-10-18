package main
import "fmt"
import "sort"
import "strconv"

func main() {
	slice := make( []int, 3)
    fmt.Printf("Enter element : ")
	var choice string
	fmt.Scan(&choice)
	var i = 0
	for choice != "X" {
		num,_ := strconv.Atoi(choice)
		if i < 1 {
			slice[i] = num
			i++
			slice = slice[:i]
		} else {
			slice = append( slice, num)
		}
		sort.Ints(slice)
		fmt.Println("Sorted slice: ", slice)
		fmt.Printf("Enter element : ")
		fmt.Scan(&choice)
	}

}