package repository

import (
	"bufio"
	"encoding/json"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

const appendToFile int = os.O_APPEND | os.O_CREATE | os.O_RDWR
const fileName = "storage"

func storageFile(flag int) (file *os.File, err error) {
	file, err = os.OpenFile(fileName, flag, 0666)
	return
}

func readLines() (notes []Note, err error) {
	file, err := storageFile(os.O_CREATE | os.O_RDONLY)
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

func writeToFile(note *Note, file *os.File) (err error) {
	marshaledJson, err := json.Marshal(note)
	file.Write([]byte(marshaledJson))
	file.WriteString("\n")
	return
}

func rewriteFile(notes []Note) (err error) {
	var jsonNotes []string = make([]string, len(notes), len(notes))
	for i, note := range notes {
		jsonNote, _ := json.Marshal(&note)
		jsonNotes[i] = string(jsonNote)
	}
	output := strings.Join(jsonNotes, "\n")
	err = ioutil.WriteFile(fileName, []byte(output), 0666)
	return
}

func AddNote(name string, description string) (err error) {
	now := time.Now()
	note := &Note{
		id:          uuid.NewV4().String(),
		name:        name,
		description: description,
		state:       New,
		created:     now,
		updated:     now,
	}
	file, err := storageFile(appendToFile)
	defer file.Close()
	writeToFile(note, file)
	return
}

func ShowAll() (notes []Note, err error) {
	notes, err = readLines()
	return
}

func CompleteById(id string) (err error) {
	notes, err := readLines()

	for i, note := range notes {
		if note.id == id {
			note.state = Done
			notes[i] = note
		}
		if err != nil {
			println(err.Error())
			return
		}
	}
	return rewriteFile(notes)
}
