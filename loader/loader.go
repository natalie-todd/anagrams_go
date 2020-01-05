package loader

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type DictionaryData struct {
	Word string `json:"word"`
}

func LoadData() []*DictionaryData {
	file, err := os.Open("dictionary.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtLines []string

	for scanner.Scan() {
		txtLines = append(txtLines, scanner.Text())
	}

	file.Close()

	scannedDictionary := []*DictionaryData{}

	for _, eachLine := range txtLines {
		word := &DictionaryData{eachLine}
		scannedDictionary = append(scannedDictionary, word)
	}
	fmt.Println(scannedDictionary)
	return scannedDictionary
}