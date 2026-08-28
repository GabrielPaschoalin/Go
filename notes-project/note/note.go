package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) (Note, error) {

	if title == "" || content == "" {
		return Note{}, errors.New("Insert valid input")
	}

	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (n Note) DisplayNote() {
	fmt.Printf("\nNote title: %s\nNote content: %s \n", n.Title, n.Content)
}

func (n Note) Save() error {

	json, err := json.Marshal(n)

	if err != nil {
		return err
	}

	fileName := strings.ReplaceAll(n.Title, " ", "_")
	fileName = strings.ToLower(fileName)
	fileName += ".json"

	finalName := "./files/" + fileName

	err = os.WriteFile(finalName, json, 0644)

	if err != nil {
		return err
	}

	return nil
}
