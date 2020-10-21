package main

import (
	"fmt"
)

func main() {
	numbers := make([]float64, 3)
	var time float64
	fmt.Println("Enter acceleration, initial velocity, and initial displacement:")
	for i := 0; i < 3; i++ {
		fmt.Scanln(&numbers[i])
	}
	a := numbers[0]
	v0 := numbers[1]
	s0 := numbers[2]
	

	fmt.Println("Enter time in seconds:")
	fmt.Scanln(&time)
	
	fn := GenDisplaceFn(a, v0, s0)
	fmt.Printf("Displacement after %d seconds is:\n",int(time))
	fmt.Println(fn(time))
}

func GenDisplaceFn(a float64, v0 float64, s0 float64) func(float64) float64 {
	fn := func(t float64) float64 {			
		return (0.5*a*t*t + v0*t + s0)
	}
	return fn
}
