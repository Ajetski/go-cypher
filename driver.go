package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Message struct {
	command string
	message string
}

const (
	Quit     string = "QUIT"
	Password        = "PASSKEY"
	Encrypt         = "ENCRYPT"
	Decrypt         = "DECRYPT"
	Result          = "RESULT"
	Error           = "ERROR"
)

func main() {
	history := make([]string, 0)
	encryptor := make(chan Message)
	logger := make(chan Message)
	go startLogger("logs.txt", logger)
	go startEncryptor(encryptor)

	for {
		// get user input
		msg := menu(history)

		// log input and send msg to encryptor
		logger <- msg
		encryptor <- msg

		if msg.command == Encrypt || msg.command == Decrypt {
			if !contains(history, msg.message) {
				history = append(history, msg.message)
			}
		}

		// quit if necessary
		if msg.command == Quit {
			time.Sleep(1000000)
			return
		}

		// get and log result
		result := <-encryptor
		logger <- result
		if result.command == Error {
			println(result.command + ": " + result.message)
		} else {
			println(result.message)
		}
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func getLine() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = text[:len(text)-1]
	return text
}

func printHistory(history []string) {
	fmt.Println("History:")
	for i, h := range history {
		fmt.Println(fmt.Sprint(i) + ") ", h)
	}
}

func selectFromHistory(history []string) string {
	printHistory(history)
	fmt.Print("Enter a number to choose from the list: ")
	text := getLine()
	idx, err := strconv.Atoi(text)
	if err != nil || idx < 0 || idx >= len(history) {
		return selectFromHistory(history)
	}
	return history[idx]
}

func useHistory() bool {
	fmt.Print("Use history (say no to enter a new string)? (y/N): ")
	text := getLine()
	return text == "y" || text == "Y"
}

func menu(history []string) Message {
	fmt.Print("Enter a command (help for options): ")
	text := getLine()
	if text == "help" {
		fmt.Println("Availbe commands are decrypt, encrypt, help, password, & quit")
		return menu(history)
	} else if text == "history" {
		printHistory(history)
		return menu(history)
	} else if text == "quit" {
		return Message{
			command: Quit,
			message: "",
		}
	} else if text == "password" {
		if len(history) > 0 && useHistory() {
			return Message{
				command: Password,
				message: selectFromHistory(history),
			}
		} else {
			fmt.Print("Enter a new enryption password: ")
			text := getLine()
			return Message{
				command: Password,
				message: text,
			}
		}
	} else if text == "encrypt" {
		if len(history) > 0 && useHistory() {
			return Message{
				command: Encrypt,
				message: selectFromHistory(history),
			}
		} else {
			fmt.Print("Enter some text to encrypt: ")
			text := getLine()
			return Message{
				command: Encrypt,
				message: text,
			}
		}
	} else if text == "decrypt" {
		if len(history) > 0 && useHistory() {
			return Message{
				command: Encrypt,
				message: selectFromHistory(history),
			}
		} else {
			fmt.Print("Enter some text to decrypt: ")
			text := getLine()
			return Message{
				command: Decrypt,
				message: text,
			}
		}
	} else {
		fmt.Println("I didn't understand that one. Try again?")
		return menu(history)
	}
}
