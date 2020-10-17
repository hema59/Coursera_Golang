package main

import "fmt"
import "strings"

func main() {
	fmt.Println("Enter string : ")
	var s string
	fmt.Scan(&s)
	s = strings.ToLower(s)
	if strings.HasPrefix(s,"i") && strings.Contains(s,"a") && s[len(s)-1 : ]== "n" {
		fmt.Println("Found!")	
		return
	}
	fmt.Println("Not found!")
}