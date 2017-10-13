package main

import (
	"encoding/json"
	"fmt"
	"github.com/SeriyBg/godo/storage"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestAddNewNote(t *testing.T) {
	oldArgs := os.Args
	os.Args = []string{oldArgs[0], "add", "-n=some_name", "-d=some_description"}
	err := os.Setenv("GODO_CONFIG", "testdata/.godoconfig")
	assert.NoError(t, err)
	main()
	os.Args = oldArgs
}

func TestShowAll(t *testing.T) {
	//t.Skip()
	oldArgs := os.Args
	os.Args = []string{oldArgs[0], "show"}
	main()
	os.Args = oldArgs
}

func TestCompleteById(t *testing.T) {
	//t.Skip()
	oldArgs := os.Args
	var id string
	os.Args = []string{oldArgs[0], "complete", "-i=" + id}
	main()
	os.Args = oldArgs
}

func assertNotesMatches(t *testing.T, actual storage.Note, expectedMap map[string]string) {
	actualMap := make(map[string]string)
	actualBytes, _ := json.Marshal(actual)
	json.Unmarshal(actualBytes, &actualMap)

	assert.Equal(t, actualMap["name"], expectedMap["name"])
	assert.Equal(t, actualMap["description"], expectedMap["description"])
	assert.Equal(t, actualMap["state"], expectedMap["state"])
}

func TestJsonToMap(t *testing.T) {
	unmarshaed := make(map[string]string)
	jsonString := `{"name": "foo", "description": "bar"}`
	json.Unmarshal([]byte(jsonString), &unmarshaed)
	fmt.Println(unmarshaed)
}
