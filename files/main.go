package main

import (
	"files/fileutils"
	"fmt"
	"os"
)

func main() {

	rootPath, _ := os.Getwd()
	content, err := fileutils.ReadTextFile(rootPath + "/data/text.txt")
	if err == nil {
		fmt.Println(content)
	} else {
		fmt.Printf("Error panic %v", err)
	}
}
