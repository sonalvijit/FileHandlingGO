package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Hello(name string) {
	fmt.Println("Hello, ", name)
}

func getUserInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func SaveToFile(filename string, content string) error {
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	filename := getUserInput("Enter the filename:")
	content_ := getUserInput("Enter the content:")
	err := SaveToFile(filename, content_)
	if err != nil {
		fmt.Println("Error saving file:", err)
	} else {
		fmt.Println("File saved successfully as:", filename)
	}
}
