package main
import "fmt"
import "encoding/json"

func main() {
	var name string
	var address string
	fmt.Println("Enter first name: ")
	fmt.Scan(&name)
	fmt.Println("Enter address: ")
	fmt.Scan(&address)
	Person := map[string]interface{} {"name" : name, "address" : address }
	//creating the json object
	jsonPerson, _ := json.Marshal(Person)
	fmt.Println("The json object is : ", jsonPerson)

	//creating the json string into string
	jsonString := string(jsonPerson)
	fmt.Println("The json string is : ",jsonString)
}