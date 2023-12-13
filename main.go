package main

import (
	"os"
)

func main(){
	if len(os.Args) < 3 {
		printHelp()
	}

	handleCommands(os.Args[2])
}

func handleCommands(input string) {
	switch input {
	case "help":
		printHelp()
	case "encrypt":
		encryptFile()
	case "decrypt":
		decryptFile()
	default:
		fmt.Println("Invalid command entered. Tr `go run . fileName.img encrypt` ")
	}
}

func printHelp(){}

func encryptFile(){}

func decryptFile(){}

func getPassword(){}

func validatePassword(){}
