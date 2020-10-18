package main

import(
	"fmt"
	"sort"
)

func main(){
	//create empty slice of capcity 3 
	usrSlice := make([] int, 0,3)
	var temp int;
	fmt.Printf("\n Initial length of slice is %d and capcity is %d \n\n", len(usrSlice), cap(usrSlice));
	for{
		fmt.Println("Enter int value : ")
		_, err := fmt.Scan(&temp);
		if(err != nil){
			break;
		}
		usrSlice = append(usrSlice, temp);
		sort.Ints(usrSlice)
		fmt.Println(usrSlice);
	}
	//fmt.Printf("\n final length of slice is %d and capcity is %d \n\n", len(usrSlice), cap(usrSlice));
}