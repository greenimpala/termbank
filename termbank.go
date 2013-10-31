package main

import (
	"flag"
	"fmt"
	"termbank/domain"
	"termbank/renderer"
)

var (
	helpFlag = flag.Bool("help", false, "Account")
)

func main() {
	flag.Parse()

	if *helpFlag == true {
		renderer.PrintHelp()
	} else {
		inititialise()
	}
}

func inititialise() {
	credentials := renderer.PromptCredentials()
	user := domain.NewUser(credentials)
	user.Login()

	fmt.Println("\nLogged in!\n")
	renderer.PrintAccountList(user)
	renderer.PromptAndPrintAccount(user)
}
