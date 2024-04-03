package fileutils

import "os"

func ReadTextFile(fileName string) (string, error) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		// we couldnt read the file
		return "", err
	} else {
		// read  operation successful
		return string(content), nil
	}
}
