package helper

import (
	"bufio"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		log.Println(e)
	}
}

func generateUsernames() {

	readFile, err := os.Open("users.txt")
	check(err)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		usernameList = append(usernameList, fileScanner.Text())
	}

	readFile.Close()
}

func generateMessages() {

	readFile, err := os.Open("messages.txt")
	check(err)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		messageList = append(messageList, fileScanner.Text())
	}

	readFile.Close()
}

func GenerateLists() {
	generateMessages()
	generateUsernames()
}
