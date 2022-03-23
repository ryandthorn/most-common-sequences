package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/test/simple", handleTestSimple)
	http.HandleFunc("/test/full", handleTestFull)
	http.ListenAndServe(":8080", nil)
}

func handleTestSimple(w http.ResponseWriter, req *http.Request) {
	f, err := mostCommonSequences("testfiles/simple.txt")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	fmt.Fprint(w, f.Top100ToString())
}

func handleTestFull(w http.ResponseWriter, req *http.Request) {
	f, err := mostCommonSequences("testfiles/darwin-full.txt")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	fmt.Fprint(w, f.Top100ToString())
}

func mostCommonSequences(filepath string) (*file, error) {
	f, err := loadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("error loading file: %w", err)
	}

	wgs, err := processFile(f)
	if err != nil {
		return nil, fmt.Errorf("error processing file contents: %w", err)
	}
	f.SetTop100(wgs)

	return f, nil
}
