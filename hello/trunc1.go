package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Float Truncator")
	fmt.Println("======================")

	for {
		fmt.Print("Number for trunc: ")
		txt, _ := reader.ReadString('\n')
		txt = strings.Replace(txt, "\n", "", -1)
		if value, err := strconv.ParseFloat(txt, 64); err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("Truncated value: %.0f\n", value)
		}
	}

}