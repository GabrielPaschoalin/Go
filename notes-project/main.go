package main

import (
	"bufio"
	"fmt"
	"gabriel/notes/note"
	"os"
	"strings"
)

func main() {

	title, content := getNoteData()
	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
	}

	userNote.DisplayNote()

	err = userNote.Save()

	if err != nil {
		fmt.Println(err)
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
