package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	fmt.Print(strings.Repeat("=== ", 2))
	fmt.Print("Bank CLI Go")
	fmt.Println(strings.Repeat(" === ", 2))

	сreateDemoAcc()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		printMenu()
		fmt.Println("Выберите действие:")
		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())
		switch choice {
		case "1":
			viewInfo()
		case "2":
			depositMoney(scanner)
		case "3":
			withdrawMoney(scanner)
		case "4":
			createNewAccount(scanner)
		case "0":
			fmt.Println("До свидания !!!")
			return
		default:
			fmt.Println("Неверный выбор. За информацией обратитесь в меню.")
		}
	}
}

func сreateDemoAcc() {
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
	if currentAccount == nil {
		fmt.Println("Аккаунт не выбран")
		return
	}
	fmt.Printf("Здравствуйте вот ваши данные:\n ID: %s\n Name: %s\n Balance: %.2f\n ",
		currentAccount.ID, currentAccount.Name, currentAccount.Balance)
}

func printMenu() {
	fmt.Print(strings.Repeat("--- ", 3))
	fmt.Print("Меню")
	fmt.Println(strings.Repeat(" ---", 3))
	fmt.Println("1. Посмотреть баланс")
	fmt.Println("2. Внести деньги")
	fmt.Println("3. Снять деньги")
	fmt.Println("4. Создать новый аккаунт")
	fmt.Println("0. Выйти")
	fmt.Println(strings.Repeat(" --- ", 5))
}

func depositMoney(scanner *bufio.Scanner) {
	if currentAccount == nil {
		fmt.Println("Аккаунт не выбран")
		return
	}
	fmt.Println("Внесите деньги")
	scanner.Scan()
	amountStr := scanner.Text()
	amountFlt, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		fmt.Printf("Ошибка преобразования строки в число %w\n", err)
		return
	}
	if amountFlt <= 0 {
		fmt.Println("Ошибка: сумма должна быть положительной")
		return
	}
	currentAccount.Balance += amountFlt
	accounts[currentAccount.ID] = *currentAccount

	fmt.Printf("Успешно внесено %.2f\n", amountFlt)
	fmt.Printf("Ваш баланс: %.2f\n", currentAccount.Balance)
}

func withdrawMoney(scanner *bufio.Scanner) {
	if currentAccount == nil {
		fmt.Println("Аккаунт не выбран")
		return
	}
	fmt.Println("Внесите сумму для снятия")
	scanner.Scan()
	amountStr := scanner.Text()

	amountFlt, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		fmt.Printf("Ошибка преобразования строки в число %w\n", err)
		return
	}
	if amountFlt <= 0 {
		fmt.Println("Ошибка: сумма должна быть положительной")
		return
	}

	currentAccount.Balance -= amountFlt
	accounts[currentAccount.ID] = *currentAccount
	fmt.Printf("Успешно выведено %.2f\n", amountFlt)
	fmt.Printf("Ваш баланс: %.2f\n", currentAccount.Balance)
}

func createNewAccount(scanner *bufio.Scanner) {
	fmt.Println(strings.Repeat(" --- ", 5))
	fmt.Println("Создание нового счета")
	fmt.Println("Введите имя")
	scanner.Scan()
	name := strings.TrimSpace(scanner.Text())

	if name == "" {
		fmt.Println("Имя не может быть пустым")
		return
	}

	fmt.Println("Введите баланс")
	scanner.Scan()
	balanceStr := strings.TrimSpace(scanner.Text())
	balanceFlt, err := strconv.ParseFloat(balanceStr, 64)
	if err != nil {
		fmt.Printf("Ошибка преобразования строки в число %w\n", err)
		return
	}
	if balanceFlt <= 0 {
		fmt.Println("Баланс не может быть пустым или орицательным")
		return
	}

	id := generateID()
	account := Account{
		ID:      id,
		Name:    name,
		Balance: balanceFlt,
	}
	accounts[id] = account
	if currentAccount == nil {
		currentAccount = &account
	}
	fmt.Println(strings.Repeat(" --- ", 5))
	fmt.Println("Аккаунт создан")
	fmt.Println("ID аккаунта:", account.ID)
	fmt.Println("Имя клиента:", account.Name)
	fmt.Println("Баланс клиента:", account.Balance)
	fmt.Println(strings.Repeat(" --- ", 5))

	fmt.Println("Сделать счет основным ?")
	scanner.Scan()
	answer := strings.TrimSpace(scanner.Text())
	if answer == "yes" || answer == "да" || answer == "y" {
		currentAccount = &account
	}
}
