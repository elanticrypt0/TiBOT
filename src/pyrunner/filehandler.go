package pyrunner

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

func OpenFile(file string) []byte {
	if _, err := os.Stat(file); err != nil {
		log.Fatalln("File doesn't exists: " + file)
	}
	filedata, err := os.ReadFile(file)
	if err != nil {
		log.Fatalln("Can not open file: " + file)
	}
	return filedata
}

func ReadAndParseJson[T any](file string, stru *T) {
	fileData := strings.NewReader(string(OpenFile(file)))
	jsonParser := json.NewDecoder(fileData)
	jsonParser.Decode(&stru)
}
