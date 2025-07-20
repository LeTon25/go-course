package cmdmanager

import (
	"fmt"
)

type CMDManager struct {
}

func (cm *CMDManager) ReadLines() ([]string, error) {
	// Implementation for reading lines from a command input
	var lines []string
	for {
		var price string
		fmt.Print("Enter price (or 'exit' to finish): ")
		_, err := fmt.Scanln(&price)
		if err != nil || price == "exit" {
			break
		}
		lines = append(lines, price)
	}
	return lines, nil
}

func (fm *CMDManager) WriteJson(data interface{}) error {
	fmt.Println("Writing JSON data to command output")
	return nil
}

func NewCMDManager() CMDManager {
	return CMDManager{}
}
