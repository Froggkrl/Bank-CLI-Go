package main

import (
	"fmt"
	"strings"
)

type Account struct {
	ID      string
	Name    string
	Balance float64
}

func main() {
	fmt.Print(strings.Repeat("&===& ", 2))
	fmt.Print("Bank CLI Go")
	fmt.Println(strings.Repeat(" &===& ", 2))

	testAccount := Account{
		ID:      "A001",
		Name:    "Ivan",
		Balance: 1000.0,
	}

	fmt.Printf("Здравствуйте вот ваши данные:\n ID: %s\n Name: %s\n Balance: %.2f\n ", testAccount.ID, testAccount.Name, testAccount.Balance)

}
