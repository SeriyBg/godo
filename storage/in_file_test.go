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
	repo := inFileRepository{fileName: "../testdata/storage/add_show_storage"}
	assert.NoError(t, repo.Create("test_add_and_show", "generated_description"))

	notes := showAllNotes(t, repo)
	assert.Len(t, notes, 1)
	assert.Equal(t, "test_add_and_show", notes[0].name)
	assert.Equal(t, "generated_description", notes[0].description)
	assert.Equal(t, New, notes[0].state)
}

func TestInFileRepository_CompleteById(t *testing.T) {
	repo := inFileRepository{fileName: "../testdata/storage/complete_storage"}
	noteId := createNewNote(t, repo)
	assert.NoError(t, repo.CompleteById(noteId))

	notes := showAllNotes(t, repo)
	assert.Len(t, notes, 1)
	assert.Equal(t, Done, notes[0].state)
}

func TestInFileRepository_DeleteById(t *testing.T) {
	repo := inFileRepository{fileName: "../testdata/storage/delete_storage"}
	noteId := createNewNote(t, repo)
	assert.NoError(t, repo.DeleteById(noteId))

	notes := showAllNotes(t, repo)
	assert.Len(t, notes, 0)
}

func createNewNote(t *testing.T, repo inFileRepository) (id string) {
	assert.NoError(t, repo.Create("generated_name", "generated_description"))
	notes := showAllNotes(t, repo)
	assert.Len(t, notes, 1)
	return notes[0].id
}

func showAllNotes(t *testing.T, repo Repository) []Note {
	notes, err := repo.GetAll()
	assert.NoError(t, err)
	return notes
}
