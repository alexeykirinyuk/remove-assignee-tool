package main

import (
	"fmt"
)

func main() {
	err := RemoveAssigneesFromDoneTickets()
	if err != nil {
		fmt.Printf("Oh, no, something terrible has happened... (%s)\r\n", err.Error())
		return
	}

	fmt.Println("Done!")
}
