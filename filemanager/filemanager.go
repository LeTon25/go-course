package filemanager

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	defer file.Close()

	return lines, nil
}

func (fm FileManager) WriteJson(data interface{}) error {
	file, err := os.Create(fm.OutputFilePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}

	time.Sleep(3 * time.Second) // Simulate a delay for demonstration purposes	)

	encoder := json.NewEncoder(file)

	err = encoder.Encode(data)

	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return err
	}

	defer file.Close()

	return nil
}

func NewFileManager(inputFilePath, outputFilePath string) FileManager {
	return FileManager{
		InputFilePath:  inputFilePath,
		OutputFilePath: outputFilePath,
	}
}
