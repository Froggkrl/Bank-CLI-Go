package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Account struct {
	ID      string
	Name    string
	Balance float64
}

var accounts = make(map[string]Account)
var currentAccount *Account

func main() {
	fmt.Print(strings.Repeat("&===& ", 2))
	fmt.Print("Bank CLI Go")
	fmt.Println(strings.Repeat(" &===& ", 2))

	CreateDemoAcc()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		printMenu()
		fmt.Println("Выберите действие:")
		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())
		switch choice {
		case "1":
			viewInfo()
		}

	}
}

func CreateDemoAcc() {
	account := Account{
		ID:      generateID(),
		Name:    "Ivan",
		Balance: 1000.0,
	}
	accounts[account.ID] = account
	currentAccount = &account
	fmt.Printf("Создан демо-счет %s\n", account.ID)
}

func generateID() string {
	return fmt.Sprintf("ACC%d", len(accounts)+1)
}

func viewInfo() {
	fmt.Printf(`Здравствуйте вот ваши данные:\n ID: %s\n Name: %s\n Balance: %.2f\n `,
		currentAccount.ID, currentAccount.Name, currentAccount.Balance)
}

func printMenu() {
}
