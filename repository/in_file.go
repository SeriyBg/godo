package repository

import (
	"bufio"
	"encoding/json"
	"os"
	"time"
)

var lastId int = 1

func storageFile() (file *os.File, err error) {
	file, err = os.OpenFile("storage", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	return
}

func readLines() (notes []Note, err error) {
	file, err := storageFile()
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		note := &Note{}
		line := scanner.Text()
		json.Unmarshal([]byte(line), note)
		notes = append(notes, *note)
	}

	if err = scanner.Err(); err != nil {
		println(err.Error())
	}
	return
}

func writeToFile(b []byte) (err error) {
	file, err := storageFile()
	defer file.Close()
	file.Write(b)
	file.WriteString("\n")
	return
}

func AddNote(name string, description string) (err error) {
	now := time.Now()
	note := &Note{
		id:          lastId,
		name:        name,
		description: description,
		state:       New,
		created:     now,
		updated:     now,
	}
	lastId += 1
	marshaledJson, err := json.Marshal(note)
	writeToFile(marshaledJson)
	return
}

func ShowAll() (notes []Note, err error) {
	AddNote("Test", "Test test test")
	AddNote("Test2", "Test2 test2 test2")
	notes, err = readLines()
	return
}
