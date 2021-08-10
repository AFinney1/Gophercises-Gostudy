package main

import "fmt"

func main() {
	//emails := make(map[string]string)

	//emails["Robert"] = "robert@gmail.com"
	//emails["John"] = "john@gmail.com"
	//emails["Dude"] = "mike@gmail.com"
	emails := map[string]string{"Robert": "robert@gmail.com", "John": "john@gmail.com", "Dude": "dude@gmail.com"}

	fmt.Println(emails)
	fmt.Println(emails["Robert"])

	delete(emails, "Dude")
	fmt.Println(emails)
}
