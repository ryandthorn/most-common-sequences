package main

import (
	"fmt"
	"os"
)

func main() {
	if err := mostCommonSequences(); err != nil {
		fmt.Println(err)
	}
}

func mostCommonSequences() error {
	files, err := loadFiles(os.Args[1:])
	if err != nil {
		return fmt.Errorf("error loading files from args: %w", err)
	}

	fromStdin, err := loadStdin()
	if err != nil {
		return fmt.Errorf("error loading text from stdin: %w", err)
	}

	if fromStdin != nil {
		files = append(files, fromStdin)
	}

	fmt.Printf("files loaded: %d\n", len(files))
	return nil
}
