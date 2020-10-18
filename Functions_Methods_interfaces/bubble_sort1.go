package main

import (
	"bufio"
	_ "fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func BubbleSort(slice []int) []int {
	length := len(slice)

	for i := 0; i < length; i++ {

		for j := 0; j < length-i-1; j++ {

			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
	return slice
}

func Swap(slice []int, i int) []int {
	slice[i], slice[i+1] = slice[i+1], slice[i]
	return slice
}

func ReadNumbers() []int {
	var text string
	log.Print("Please. Enter 10 integer numbers: ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text = scanner.Text()
		break
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	chars := strings.Split(text, " ")
	numbers := make([]int, len(chars))

	for i, v := range chars {
		numbers[i], _ = strconv.Atoi(v)
	}

	return numbers
}

func ReadIndex() int {
	var text string
	var index int

	log.Print("Please. Enter an index to Swap: ")

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text = scanner.Text()
		break
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	index, _ = strconv.Atoi(text)
	return index
}

func main() {
	numbers := BubbleSort(ReadNumbers())
	log.Println(numbers)
	index := ReadIndex()
	numbers = Swap(numbers, index)
	log.Println(numbers)
}