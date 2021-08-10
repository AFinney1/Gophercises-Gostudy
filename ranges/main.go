package main

import "fmt"

func main() {
	ids := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	for _, id := range ids {
		fmt.Printf("ID:%d\n", id)
	}

	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	emails := map[string]string{"Bob": "bob@gmail.com", "John": "john@gmail.com", "Mary": "mary@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range emails {
		fmt.Printf("%s: %s\n", k, emails[k])
	}
}
