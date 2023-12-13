package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/adarsh2858/file-password-lock/filecrypt"

	"golang.org/x/term"
)

const (
	commandHelp    = "help"
	commandEncrypt = "encrypt"
	commandDecrypt = "decrypt"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(0)
	}

	// follow a functional programming approach
	handleCommands(os.Args[1])
}

func printHelp() {
	fmt.Println("File Lock and Unlock by Password")
	fmt.Println()
	fmt.Println("Simple file encryptor for your day to day needs.")
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("")
	fmt.Println("To encrypt a file command is:  go run . encrypt /path/to/fileName.png")
	fmt.Println("To decrypt a file command is: go run . decrypt /path/to/fileName.png")
}

func handleCommands(function string) {
	switch function {
	case commandHelp:
		printHelp()
	case commandEncrypt:
		handleEncryption()
	case commandDecrypt:
		handleDecryption()
	default:
		fmt.Println("Invalid command entered. Tr `go run . encrypt /path/to/fileName.img` ")
		os.Exit(1)
	}
}

func handleEncryption() {
	// accept the file name as string input
	if len(os.Args) < 3 {
		fmt.Println("File name is not provided")
		fmt.Println("Run the command `go run . help` to dig deeper")
		os.Exit(0)
	}

	// validate file using os.Stat if not exists panic
	file := os.Args[2]
	ok := validateFile(file)
	if !ok {
		fmt.Println("Please try again")
		panic("File not found")
		return
	}

	password := getPassword()

	fmt.Println("Encrypting...")
	filecrypt.Encrypt(file, password)
	fmt.Println("successfully encrypted file!")
}

func validateFile(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		// panic("File does not exist")
		return false
	}
	return true
}

func getPassword() []byte {
	// using the following to not let the password be visible in terminal
	// while interacting with this file encryptor app
	fmt.Print("Enter Password")
	password1, err := term.ReadPassword(0)
	if err != nil {
		panic(err)
	}

	fmt.Print("\nConfirm Password\n")
	password2, err := term.ReadPassword(0)
	if err != nil {
		panic(err)
	}

	// no need of the following conversion:
	// password := []byte(password1)
	// confirmPassword := []byte(password2)
	if !isPasswordValid(password1, password2) {
		fmt.Println("Confirm password didn't match")
		return getPassword()
	}

	return password1
}

func isPasswordValid(password1, password2 []byte) bool {
	if bytes.Equal(password1, password2) {
		return true
	}
	return false
}

func handleDecryption() {
	if len(os.Args) < 3 {
		fmt.Println("File does not exist. For more info, try go run . help")
		os.Exit(0)
	}

	file := os.Args[2]
	if ok := validateFile(file); !ok {
		fmt.Println("Please try again")
		return
	}

	fmt.Print("Enter password")
	password, err := term.ReadPassword(0)
	if err != nil {
		panic(err)
	}

	fmt.Println("Decrypting...")
	filecrypt.Decrypt(file, password)
	fmt.Println("successfully decrypted file!")
}
