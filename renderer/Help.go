package renderer

import (
	"fmt"
)

func PrintHelp() {
	fmt.Print("\n  / /____  _________ ___  / /_  ____ _____  / /__\n ")
	fmt.Print("/ __/ _ \\/ ___/ __ `__ \\/ __ \\/ __ `/ __ \\/ //_/\n/ ")
	fmt.Print("/_/  __/ /  / / / / / / /_/ / /_/ / / / / ,<   \n\\__/\\")
	fmt.Println("___/_/  /_/ /_/ /_/_.___/\\__,_/_/ /_/_/|_|  \n")
	fmt.Println("  Usage: termbank [options]\n")
	fmt.Println("  Options:\n")
	printOption("h", "Output usage information")
	fmt.Print("\n")
}

func printOption(flag string, info string) {
	fmt.Println("     -" + flag + ", --" + flag + "\t " + info)
}
