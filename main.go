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

	for _, f := range files {
		wgs, err := processFile(f)
		if err != nil {
			return fmt.Errorf("error processing file contents: %w", err)
		}
		f.SetTop100(wgs)
	}

	printMostCommonSequences(files)

	return nil
}

func printMostCommonSequences(files []*file) {
	for _, f := range files {
		fmt.Printf("\nResults for %s\n", f.Name())
		fmt.Println("---------------------------------")
		for _, wg := range f.GetTop100() {
			fmt.Printf("%s: %d\n", wg.Text(), wg.Count())
		}
	}
}
