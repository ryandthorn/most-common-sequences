package main

import (
	"fmt"
	"net/http"
	"strings"
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
	fmt.Fprint(w, stringifyMostCommonSequences(f))
}

func handleTestFull(w http.ResponseWriter, req *http.Request) {
	f, err := mostCommonSequences("testfiles/darwin-full.txt")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	fmt.Fprint(w, stringifyMostCommonSequences(f))
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

func printMostCommonSequences(f *file) {
	fmt.Print(stringifyMostCommonSequences(f))
}

func stringifyMostCommonSequences(f *file) string {
	b := strings.Builder{}
	b.WriteString(fmt.Sprintf("\nResults for %s\n", f.Name()))
	b.WriteString("---------------------------------\n")
	for _, wg := range f.GetTop100() {
		b.WriteString(fmt.Sprintf("%s: %d\n", wg.Text(), wg.Count()))
	}
	return b.String()
}
