package util

import (
	"os"
	"strings"
)

func GetFile(filename string) []string {
	file, err := os.ReadFile(filename)
	if err != nil {
		panic("Unable to load file: " + err.Error())
	}
	return strings.Split(strings.ReplaceAll(string(file), "\r\n", "\n"), "\n") //Windows
}