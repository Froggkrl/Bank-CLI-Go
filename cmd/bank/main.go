package main

import (
	"BankCLI/internal/config"
	"BankCLI/internal/funcs"
	"BankCLI/pkg/models"
	"bufio"
	"fmt"
	"os"
	"strings"
)

var currentAccount *models.Account

func main() {

	if err := config.Init(); err != nil {
		fmt.Printf("Ошибка: %w", err)
	}

	cfg := config.GetConfig()

	fmt.Print(strings.Repeat("===", 5))
	fmt.Printf("%s %s", cfg.AppName, cfg.Version)
	fmt.Print(strings.Repeat("=== ", 5))

	currentAccount = funcs.CreateDemoAcc()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		funcs.PrintMenu()
		fmt.Println("Выберите действие:")
		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())
		switch choice {
		case "1":
			funcs.ViewInfo(currentAccount)
		case "2":
			funcs.DepositMoney(scanner, currentAccount)
		case "3":
			funcs.WithdrawMoney(scanner, currentAccount)
		case "4":
			newAccount := funcs.CreateNewAccount(scanner)
			if newAccount != nil {
				currentAccount = newAccount
			}
		case "0":
			fmt.Println("До свидания !!!")
			return
		default:
			fmt.Println("Неверный выбор. За информацией обратитесь в меню.")
		}
	}
}
