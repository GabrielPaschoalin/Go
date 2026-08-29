package main

import (
	"bufio"
	"fmt"
	"gabriel/notes/note"
	"gabriel/notes/todo"
	"os"
	"strings"
)

type saver interface {
	Save() error
}

// type displayer interface {
// 	Display()
// }

type outputtable interface {
	saver
	Display()
	// DoSomething(int) string
}

// type outputtable interface {
// 	Save() error
// 	Display()
// }

func main() {

	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println()

	// Print and save todo
	err = outputData(todo)

	if err != nil {
		return
	}

	// Print and save note
	err = outputData(userNote)

	if err != nil {
		return
	}
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")
	return title, content

}

func getUserInput(text string) string {
	var input string
	fmt.Printf("%s ", text)

	// Pegar uma string completa a partir do teclado
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	// Remover o lineBreak
	text = strings.TrimSuffix(input, "\n")
	// Remover o \r
	text = strings.TrimSuffix(text, "\r")

	return text
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving note failed")
		return err
	}

	fmt.Print("Saving the note succeeded!")
	return nil
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}
