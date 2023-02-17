package main

import (
  "github.com/odysseus/vigenere"
)

func startEncryptor(msgs chan Message) {
	password := ""

	for {
		msg := <-msgs
		if msg.command == Quit {
			return
		} else if msg.command == Password {
			password = msg.message
			msgs <- Message{
				command: Result,
				message: "Password set",
			}
		} else if msg.command == Encrypt {
			if password == "" {
				msgs <- Message{
					command: Error,
					message: "Password not set",
				}
			} else {
				msgs <- Message{
					command: Result,
					message:  vigenere.Encipher(msg.message, password),
				}
			}
		} else if msg.command == Decrypt {
			if password == "" {
				msgs <- Message{
					command: Error,
					message: "Password not set",
				}
			} else {
				msgs <- Message{
					command: Result,
					message: vigenere.Decipher(msg.message, password),
				}
			}
		} else {
			msgs <- Message{
				command: Error,
				message: "Unrecognized command sent to Encryptor",
			}
		}
	}
}
