package storage

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

type inFileRepository struct {
	fileName string
}

func (r inFileRepository) storageFile(flag int) (file *os.File, err error) {
	file, err = os.OpenFile(r.fileName, flag, 0666)
	return
}

func (r inFileRepository) readLines() (notes []Note, err error) {
	file, err := r.storageFile(readFromFile)
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

func (r inFileRepository) writeToFile(note *Note, file *os.File) (err error) {
	marshaledJson, err := json.Marshal(note)
	file.Write([]byte(marshaledJson))
	file.WriteString("\n")
	return
}

func (r inFileRepository) Create(name string, description string) (err error) {
	now := time.Now()
	note := &Note{
		id:          uuid.NewV4().String(),
		name:        name,
		description: description,
		state:       New,
		created:     now,
		updated:     now,
	}
	file, err := r.storageFile(appendToFile)
	defer file.Close()
	r.writeToFile(note, file)
	return
}

func (r inFileRepository) GetAll() (notes []Note, err error) {
	notes, err = r.readLines()
	return
}

func (r inFileRepository) CompleteById(id string) (err error) {
	notes, err := r.readLines()
	file, err := r.storageFile(rewriteToFile)
	var wg sync.WaitGroup
	lock := make(chan bool, 1)

	for i, note := range notes {
		if note.id == id {
			note.state = Done
			notes[i] = note
		}
		wg.Add(1)
		go func(n Note, lock chan bool) {
			lock <- true
			r.writeToFile(&n, file)
			wg.Done()
			<-lock
		}(note, lock)
		if err != nil {
			println(err.Error())
			return
		}
	}
	wg.Wait()
	return
}

func (r inFileRepository) DeleteById(id string) (err error) {
	notes, err := r.readLines()
	file, err := r.storageFile(rewriteToFile)
	var wg sync.WaitGroup
	lock := make(chan bool, 1)

	for _, note := range notes {
		if note.id != id {
			wg.Add(1)
			go func(n Note, lock chan bool) {
				lock <- true
				r.writeToFile(&n, file)
				wg.Done()
				<-lock
			}(note, lock)
			if err != nil {
				println(err.Error())
				return
			}
		}
	}
	wg.Wait()
	return
}

func (r inFileRepository) FindAllBy(filter NoteFilter) (notes []Note, err error) {
	return
}
