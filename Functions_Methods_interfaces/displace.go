package main

import "fmt"
import "math"

func main() {
    var a float64
    fmt.Println("Enter acceleration :" )
    fmt.Scan(&a)
	var v0 float64
	fmt.Println("Enter the initial velocity :")
	fmt.Scan(&v0)
	var s0 float64
	fmt.Println("Enter initial displacement: ")
	fmt.Scan(&s0)
	var t float64
	fmt.Println("Enter time: ")
	fmt.Scan(&t)
	fn := GenDisplaceFn(a,v0,s0)
	fmt.Println(fn(3))
	fmt.Println(fn(5))
}
func GenDisplaceFn(a,v0,s0 float64) func(float64) float64{
	return func(t float64) float64 {
		return 0.5*a*math.Pow(t,2)+ v0*t +s0
	}
}