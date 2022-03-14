package main

import (
	"reflect"
	"testing"
)

func Test_sortWordGroups(t *testing.T) {
	tests := []struct {
		name       string
		wordGroups []wordGroup
		want       []wordGroup
	}{
		{"sorts word groups by count in descending order", []wordGroup{{"a a a", 1}, {"b b b", 2}}, []wordGroup{{"b b b", 2}, {"a a a", 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortWordGroups(tt.wordGroups); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortWordGroups() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_processWordGroups(t *testing.T) {
	tests := []struct {
		name        string
		occurrences map[string]int
		want        []wordGroup
	}{
		{"returns array of wordGroups sorted by count in descending order", map[string]int{"a a a": 1, "b b b": 2, "c c c": 3}, []wordGroup{{"c c c", 3}, {"b b b", 2}, {"a a a", 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := processWordGroups(tt.occurrences); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("processWordGroups() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countOccurrences(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want map[string]int
	}{
		{"creates a map of three-word groups to their frequency in s", "one two three one two three four", map[string]int{"one two three": 2, "two three one": 1, "three one two": 1, "two three four": 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countOccurrences(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countOccurrences() = %v, want %v", got, tt.want)
			}
		})
	}
}
