package funcs

import (
	"BankCLI/internal/generator"
	"BankCLI/pkg/models"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func DepositMoney(scanner *bufio.Scanner, currentAccount *models.Account) {
	if currentAccount == nil {
		fmt.Println("Аккаунт не выбран")
		return
	}
	fmt.Println("Внесите деньги")
	scanner.Scan()
	amountStr := scanner.Text()
	amountFlt, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		fmt.Printf("Ошибка преобразования строки в число %v\n", err)
		return
	}
	if amountFlt <= 0 {
		fmt.Println("Ошибка: сумма должна быть положительной")
		return
	}
	currentAccount.Balance += amountFlt
	models.Accounts[currentAccount.ID] = *currentAccount

	fmt.Printf("Успешно внесено %.2f\n", amountFlt)
	fmt.Printf("Ваш баланс: %.2f\n", currentAccount.Balance)
}
func WithdrawMoney(scanner *bufio.Scanner, currentAccount *models.Account) {
	if currentAccount == nil {
		fmt.Println("Аккаунт не выбран")
		return
	}
	fmt.Println("Внесите сумму для снятия")
	scanner.Scan()
	amountStr := scanner.Text()

	amountFlt, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		fmt.Printf("Ошибка преобразования строки в число %v\n", err)
		return
	}
	if amountFlt <= 0 {
		fmt.Println("Ошибка: сумма должна быть положительной")
		return
	}

	if amountFlt > currentAccount.Balance {
		fmt.Println("Недостатточно средств")
		return
	}

	currentAccount.Balance -= amountFlt
	models.Accounts[currentAccount.ID] = *currentAccount
	fmt.Printf("Успешно выведено %.2f\n", amountFlt)
	fmt.Printf("Ваш баланс: %.2f\n", currentAccount.Balance)
}

func ViewInfo(currentAccount *models.Account) {
	if currentAccount == nil {
		fmt.Println("Аккаунт не выбран")
		return
	}
	fmt.Printf("Здравствуйте вот ваши данные:\n ID: %s\n Name: %s\n Balance: %.2f\n ",
		currentAccount.ID, currentAccount.Name, currentAccount.Balance)
}

func PrintMenu() {
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

func CreateDemoAcc() *models.Account {
	account := models.Account{
		ID:      generator.GenerateID(),
		Name:    "Ivan",
		Balance: 1000.0,
	}
	models.Accounts[account.ID] = account
	currentAccount := &account
	fmt.Printf("Создан демо-счет %s\n", account.ID)
	return currentAccount
}
func CreateNewAccount(scanner *bufio.Scanner) *models.Account {
	fmt.Println(strings.Repeat(" --- ", 5))
	fmt.Println("Создание нового счета")
	fmt.Println("Введите имя")
	scanner.Scan()
	name := strings.TrimSpace(scanner.Text())

	if name == "" {
		fmt.Println("Имя не может быть пустым")
		return nil
	}

	fmt.Println("Введите баланс")
	scanner.Scan()
	balanceStr := strings.TrimSpace(scanner.Text())
	balanceFlt, err := strconv.ParseFloat(balanceStr, 64)
	if err != nil {
		fmt.Printf("Ошибка преобразования строки в число %v\n", err)
		return nil
	}
	if balanceFlt <= 0 {
		fmt.Println("Баланс не может быть пустым или орицательным")
		return nil
	}

	id := generator.GenerateID()
	account := models.Account{
		ID:      id,
		Name:    name,
		Balance: balanceFlt,
	}
	models.Accounts[id] = account
	fmt.Println(strings.Repeat(" --- ", 5))
	fmt.Println("Аккаунт создан")
	fmt.Println("ID аккаунта:", account.ID)
	fmt.Println("Имя клиента:", account.Name)
	fmt.Println("Баланс клиента:", account.Balance)
	fmt.Println(strings.Repeat(" --- ", 5))

	fmt.Println("Сделать счет основным ?")
	scanner.Scan()
	answer := strings.ToLower(strings.TrimSpace(scanner.Text()))
	if answer == "yes" || answer == "да" || answer == "y" {
		currentAccount := &account
		return currentAccount
	}
	return nil
}
