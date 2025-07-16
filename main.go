package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type Saver interface {
	Save() error
}

type Displayable interface {
	Display()
}

type Outputtable interface {
	Saver
	Display()
}

func saveData(saver Saver) error {
	if err := saver.Save(); err != nil {
		return fmt.Errorf("saving data failed: %w", err)
	}
	return nil
}
func outputData(data Outputtable) error {
	data.Display()
	return saveData(data)
}
func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}
	err = outputData(userNote)
	if err != nil {
		fmt.Println("Saving the note failed.")
		return
	}
	fmt.Println("Saving the note succeeded!")
	/*************/
	todoText := getTodoData()
	userTodo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = outputData(userTodo)
	if err != nil {
		fmt.Println("Saving the todo failed.")
		return
	}
	fmt.Println("Saving the todo succeeded!")

}

func getTodoData() string {
	text := getUserInput("Todo text: ")
	return text
}
func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func printSomething(value interface{}) {
	switch v := value.(type) {
	case string:
		fmt.Println("String value:", v)
	case int:
		fmt.Println("Integer value:", v)
	case float64:
		fmt.Println("Float value:", v)
	default:
		fmt.Println("Unknown type")
	}
}

func printSomethingV2(value interface{}) {
	floatValue, ok := value.(float64)
	if ok {
		fmt.Println("Float value:", floatValue)
		return
	}
	intValue, ok := value.(int)
	if ok {
		fmt.Println("Integer value:", intValue)
		return
	}
	stringValue, ok := value.(string)
	if ok {
		fmt.Println("String value:", stringValue)
		return
	}
	fmt.Println("Unknown type")
}

func add[T int | float64 | string](a, b T) T {
	return a + b
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
