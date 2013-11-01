package renderer

import (
	"fmt"
	"github.com/st3redstripe/termbank/domain"
	"github.com/st3redstripe/golumn"
	"strconv"
)

func PrintAccountList(user *domain.User) {
	for i, acc := range user.Accounts {
		fmt.Println(strconv.Itoa(i+1) + ") " + acc.Name + " - " + acc.Balance)
	}
}

func PromptAndPrintAccount(user *domain.User) {
	var account *domain.Account

	for account == nil {
		fmt.Print("\nEnter valid account number: ")

		var input string
		fmt.Scanln(&input)
		accNumberInt, err := strconv.Atoi(input)

		if err == nil && accNumberInt > 0 && accNumberInt <= len(user.Accounts) {
			account = user.Accounts[accNumberInt-1]
		}
	}

	fmt.Print("\n")
	PrintAccount(account)
}

func PrintAccount(account *domain.Account) {
	fmt.Println("Account: " + account.Name)
	fmt.Println("Balance: " + account.Balance)
	fmt.Println("Fetching statement...\n")

	statement := account.StatementPretty()
	formattedStatement := golumn.ParseF(string(statement), ",", golumn.Options{
		MaxColumnWidth: 12,
		Truncate:       true,
	})
	fmt.Println(formattedStatement)
}

func PromptCredentials() map[string]string {
	credentials := make(map[string]string)
	var buffer string

	fmt.Print("User ID: ")
	fmt.Scanln(&buffer)
	credentials["id"] = buffer
	fmt.Print("Password: ")
	fmt.Scanln(&buffer)
	credentials["password"] = buffer
	fmt.Print("Memorable Info: ")
	fmt.Scanln(&buffer)
	credentials["memorable"] = buffer

	return credentials
}
