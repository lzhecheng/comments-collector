package utils

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// FindAllFiles finds all files in the path.
func FindAllFiles(root string) []string {
	var files []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}

	return files
}

// CheckFile checks a file.
func CheckFile(path string) string {
	if strings.HasPrefix(path, ".git/") {
		return ""
	}

	title := "=== " + path + " ===\n"

	file, err := os.Open(path)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	var result, current string

	for i, line := range txtlines {
		if strings.HasPrefix(line, "//") {
			current += line[2:] + " "
		} else {
			if len(current) != 0 {
				result += path + ":" + strconv.Itoa(i) + current + "\n"
				current = ""
			}
		}
	}
	if current != "" {
		result += current + "\n"
	}

	if len(result) == 0 {
		return ""
	}
	return title + result
}
