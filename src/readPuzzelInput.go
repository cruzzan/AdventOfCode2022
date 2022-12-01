package src

import (
	"bufio"
	"fmt"
	"os"
)

func ReadPuzzelInput(filePath string) []string {
	lines := make([]string, 0)
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Something went wrong", err)
		os.Exit(1)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Something went wrong", err)
	}

	return lines
}
