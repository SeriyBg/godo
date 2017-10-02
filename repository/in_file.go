package repository

import (
	"bufio"
	"encoding/json"
	"github.com/satori/go.uuid"
	"os"
	"sync"
	"time"
)

const readFromFile int = os.O_CREATE | os.O_RDONLY
const appendToFile int = os.O_APPEND | os.O_CREATE | os.O_RDWR
const rewriteToFile int = os.O_TRUNC | os.O_APPEND | os.O_RDWR
const fileName = "storage"

func storageFile(flag int) (file *os.File, err error) {
	file, err = os.OpenFile(fileName, flag, 0666)
	return
}

func readLines() (notes []Note, err error) {
	file, err := storageFile(readFromFile)
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

func rewriteFile(nc <-chan Note, file *os.File) (err error) {
	for {
		select {
		case note := <-nc:
			writeToFile(&note, file)
		}
	}
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
	file, err := storageFile(rewriteToFile)
	var wg sync.WaitGroup

	for i, note := range notes {
		if note.id == id {
			note.state = Done
			notes[i] = note
		}
		wg.Add(1)
		go func(n Note) {
			writeToFile(&n, file)
			wg.Done()
		}(note)
		if err != nil {
			println(err.Error())
			return
		}
	}
	wg.Wait()
	return
}
