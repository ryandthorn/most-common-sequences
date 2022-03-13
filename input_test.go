package main

import (
	"testing"
)

func Test_loadFile(t *testing.T) {
	tests := []struct {
		name     string
		filepath string
		wantErr  bool
	}{
		{"loads test file", "testfiles/simple.txt", false},
		{"error when file does not exist", "foo.bar", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := loadFile(tt.filepath)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
