package files

import (
	"os"
	"strings"
)

// ReadFile reads the content of the file and returns it as a string
func ReadFile(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// ParseData splits the data into lines and then into characters
func ParseData(data string) [][]string {
	var lineArr [][]string
	dataPerLine := strings.Split(data, "\n")
	for _, line := range dataPerLine {
		lineArr = append(lineArr, strings.Split(line, ""))
	}
	return lineArr
}
