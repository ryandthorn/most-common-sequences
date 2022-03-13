package main

import (
	"fmt"
	"io"
	"os"
)

func loadFiles(filepaths []string) ([]*file, error) {
	files := []*file{}
	for _, filepath := range filepaths {
		f, err := loadFile(filepath)
		if err != nil {
			return nil, err
		}
		files = append(files, f)
	}
	return files, nil
}

func loadFile(filepath string) (*file, error) {
	fileInfo, err := os.Stat(filepath)
	if err != nil {
		return nil, err
	}

	contents, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("error reading file '%s': %w", filepath, err)
	}

	return &file{name: fileInfo.Name(), contents: contents}, nil
}

func loadStdin() (*file, error) {
	isEmpty, err := isStdinEmpty()
	if err != nil {
		return nil, err
	}

	if isEmpty {
		return nil, nil
	}

	b, err := io.ReadAll(os.Stdin)
	if err != nil {
		return nil, fmt.Errorf("error reading from stdin: %w", err)
	}
	return &file{name: "stdin", contents: b}, nil
}

// checks whether stdin is empty to prevent hanging on empty input during read
// more info: https://stackoverflow.com/a/38612652
func isStdinEmpty() (bool, error) {
	stdinInfo, err := os.Stdin.Stat()
	if err != nil {
		return true, err
	}
	return stdinInfo.Mode()&os.ModeNamedPipe == 0, nil
}
