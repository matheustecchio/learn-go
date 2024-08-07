package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"example.com/note"
	"example.com/todo"
)

func getUserInput(promptText string) string {
	fmt.Print(promptText) // Prompt user for input
	var input string

	// Read user input and remove trailing newline character using TrimSuffix function from strings package.
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return ""
	}

	input = strings.TrimSuffix(input, "\n")

	return input
}

func getNoteData() (string, string) {
	title := getUserInput("Enter the title of your note: ")

	description := getUserInput("Enter the description of your note: ")

	return title, description
}

func getTodoData() string {
	description := getUserInput("Enter the text of your note: ")

	return description
}

type Saver interface {
	Save() error
}

type Displayer interface {
	Display()
}

type Outputtable interface {
	Saver
	Displayer
}

func saveData(data Saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving failed.")
		return err
	}

	fmt.Println("Saving succeeded!")
	return nil
}

func outputData(data Outputtable) {
	data.Display()
	saveData(data)
}

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)
	if err != nil {
		log.Fatal(err)
	}

	outputData(userNote)

	text := getTodoData()

	userTodo, err := todo.New(text)
	if err != nil {
		log.Fatal(err)
	}

	outputData(userTodo)

}
