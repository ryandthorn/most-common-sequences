package main

import (
	"fmt"
	"os"
)

type file struct {
	name     string
	contents []byte
}

func (f *file) Name() string {
	return f.name
}

func (f *file) Contents() []byte {
	return f.contents
}

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
