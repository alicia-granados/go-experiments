package main

import (
	"files/data"
	"files/fileutils"
	"fmt"
	"os"
)

func main() {

	var price = 34.04
	var stringPrice = fmt.Sprintf("%.2f", price)
	fmt.Println(stringPrice)
	rootPath, _ := os.Getwd()
	filePath := rootPath + "/data/text.txt"
	content, err := fileutils.ReadTextFile(filePath)
	if err == nil {
		fmt.Println(content)
		newContent := fmt.Sprintf("Original %v\nDouble Original%v %v", content, content, content)
		fileutils.WriteToFile(filePath+".output.txt", newContent)
	} else {
		fmt.Printf("Error panic %v", err)
	}

	data.Types()
	/*distance := data.ToKm(200)
	fmt.Println(distance)*/
}
