package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
)

func main() {
	title, content, err := getNoteData()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	userNote.Display()

	err = userNote.Save()

	if err != nil {
		fmt.Println("Error When Saving:", err)
		return
	}

	fmt.Println("Note saved successfully!")
}

func getNoteData() (string, string, error) {
	title, err := getUserInput("Enter the title of the note: ")
	if err != nil {
		return "", "", err
	}

	content, err := getUserInput("Enter the content of the note: ")
	if err != nil {
		return "", "", err
	}

	return title, content, nil
}

func getUserInput(promptText string) (string, error) {
	fmt.Print(promptText)
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return "", fmt.Errorf("failed to read input: %w", err)
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text, nil
}
