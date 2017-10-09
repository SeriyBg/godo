package storage

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	setUp()
	run := m.Run()
	cleanUp()
	os.Exit(run)
}

func setUp() {
	os.MkdirAll("../testdata/storage", os.ModePerm)
}

func cleanUp() {
	os.RemoveAll("../testdata/storage")
}

func TestInFileRepository_AddAndShowNote(t *testing.T) {
	repo := inFileRepository{fileName: "../testdata/add_show_storage"}
	assert.NoError(t, repo.Create("test_add_and_show", "generated_description"))

	notes := showAllNotes(t, repo)
	assert.Len(t, notes, 1)
	assert.Equal(t, "test_add_and_show", notes[0].name)
	assert.Equal(t, "generated_description", notes[0].description)
	assert.Equal(t, New, notes[0].state)
}

func TestInFileRepository_CompleteById(t *testing.T) {
	repo := inFileRepository{fileName: "../testdata/complete_storage"}
	assert.NoError(t, repo.Create("generated_name", "generated_description"))

	notes := showAllNotes(t, repo)
	assert.Len(t, notes, 1)

	noteId := notes[0].id
	assert.NoError(t, repo.CompleteById(noteId))

	notes = showAllNotes(t, repo)
	assert.Len(t, notes, 1)
	assert.Equal(t, Done, notes[0].state)
}

func showAllNotes(t *testing.T, repo Repository) []Note {
	notes, err := repo.GetAll()
	assert.NoError(t, err)
	return notes
}
