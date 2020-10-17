package main

import (
	b "bufio"
	f "fmt"
	o "os"
	s "strings"
)

func main() {
	reader := b.NewReader(o.Stdin)
	f.Println("Findian")
	f.Println("======================")
	f.Print("Enter text: ")
	t, _ := reader.ReadString('\n')
	t = s.ToUpper(t)
	f.Println(t)

	if s.HasPrefix(t, "I") && s.Contains(t, "A") && s.HasSuffix(t, "N\n") {
		f.Println("Found!")
	} else {
		f.Println("Not Found!")
	}
}