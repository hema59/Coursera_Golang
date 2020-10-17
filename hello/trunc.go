package main

import "fmt"

func main() {
    var n float64
    fmt.Printf("Enter a floating point number: ")
	fmt.Scan(&n)
	fmt.Println(int(n))
}