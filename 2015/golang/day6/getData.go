package main

import (
	"os"
)

func OpenFiles(file string) ([]byte, error) {

	data, error := os.ReadFile(file)

	return data, error
}
