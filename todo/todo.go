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

func (Todo Todo) Display() {
	fmt.Printf("Your Todo has the following content:\n\n%v\n\n", Todo.Text)
}

func (Todo Todo) Save() error {
	fileName := "todo.json"

	json, err := json.Marshal(Todo)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func New(text string) (Todo, error) {
	if text == "" {
		return Todo{}, errors.New("Invalid input.")
	}

	return Todo{
		Text: text,
	}, nil
}
