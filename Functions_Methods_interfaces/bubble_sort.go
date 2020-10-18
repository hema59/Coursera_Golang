package main

import (
    "fmt"
)


func Swap( nums []int, i, j int) {
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
}
func BubbleSort(nums []int) {
	for n := len(nums); n > 1; n-- {
		for j := 1; j < n; j++ {
			if nums[j-1] > nums[j] {
				Swap(nums, j-1, j)
			}
		}
	}
}
func main() {
	var numbers [10]int
	var limit int
	fmt.Println("Enter the limit on the integer array: ")
	fmt.Scan(&limit)

	fmt.Println("Enter ",limit," integers to sort: ")
	for i := 0; i < limit; i++ {
		var temp int
		fmt.Scan(&temp)
		numbers[i]= temp
	}  
	slice := numbers[:limit]
	BubbleSort(slice)
	fmt.Println("Sorted slice : ",slice)
}