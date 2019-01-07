package main

import (
	"testing"
)

var checkFileArray = []string{
	"Dockerfile",
	"Makefile",
	"Rakefile",
}

func TestCheckFiles(t *testing.T) {
	checkedFiles := checkFiles(checkFileArray)

	expectedCounts := 2

	if expectedCounts != len(checkedFiles) {
		t.Errorf("Expected length %d != %d", expectedCounts, len(checkedFiles))
	}

	expectedFiles := map[string]bool{
		"Dockerfile": true,
		"Makefile":   true,
		"Rakefile":   false,
	}

	for _, filename := range checkedFiles {
		if val, exists := expectedFiles[filename]; exists {
			if !val {
				t.Errorf("Expected value %v != %v", !expectedFiles[filename], expectedFiles[filename])
			}
		} else {
			t.Errorf("Expected no item %v", filename)
		}
	}
}
