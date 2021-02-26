package main

import "fmt"

func main() {
	Customers := make(map[string]string)
	Customers["876fbe4a-728f-4345-8368-f65a91b035dc"] = "Jane"
	Customers["70059c66-cc28-496b-a16c-109ba24c4a28"] = "John"
	Customers["a4d9ab9a-3f58-4819-8807-6e29018e1a35"] = "Paul"
	fmt.Println(len(Customers))
	fmt.Println(Customers["a4d9ab9a-3f58-4819-8807-6e29018e1a35"])
}
