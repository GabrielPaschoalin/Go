package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func New(content string) (Todo, error) {

	if content == "" {
		return Todo{}, errors.New("Insert valid input")
	}

	return Todo{
		Text: content,
	}, nil
}

func (todo Todo) Display() {
	fmt.Println(todo.Text)
}

func (todo Todo) Save() error {

	fileName := "todo.json"
	finalName := "./todo/files/" + fileName

	json, err := json.Marshal(todo)

	if err != nil {
		return err
	}

	err = os.WriteFile(finalName, json, 0644)

	if err != nil {
		return err
	}

	return nil
}
