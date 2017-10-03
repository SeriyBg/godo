package repository

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNote_MarshalJSON(t *testing.T) {
	timeNow := time.Date(2017, 11, 3, 1, 12, 42, 1, time.UTC)
	note := &Note{
		id:          "any_gen_id",
		name:        "note_name",
		description: "note_desc",
		state:       New,
		created:     timeNow,
		updated:     timeNow,
	}
	expectedJson := "{\"id\":\"any_gen_id\",\"name\":\"note_name\"," +
		"\"description\":\"note_desc\"," +
		"\"state\":\"New\"," +
		"\"created\":\"2017-11-03T01:12:42.000000001Z\"," +
		"\"updated\":\"2017-11-03T01:12:42.000000001Z\"}"
	bytes, err := json.Marshal(note)
	assert.NoError(t, err)
	assert.Equal(t, expectedJson, string(bytes))
}

func TestNote_UnmarshalJSON(t *testing.T) {
	jsonNote := "{\"id\":\"any_gen_id\",\"name\":\"note_name\"," +
		"\"description\":\"note_desc\"," +
		"\"state\":\"New\"," +
		"\"created\":\"2017-11-03T01:12:42.000000001Z\"," +
		"\"updated\":\"2017-11-03T01:12:42.000000001Z\"}"

	note := &Note{}
	err := json.Unmarshal([]byte(jsonNote), note)
	assert.NoError(t, err)

	timeNow := time.Date(2017, 11, 3, 1, 12, 42, 1, time.UTC)
	expectedNote := Note{
		id:          "any_gen_id",
		name:        "note_name",
		description: "note_desc",
		state:       New,
		created:     timeNow,
		updated:     timeNow,
	}
	assert.Equal(t, expectedNote, *note)
}

func TestNote_String(t *testing.T) {
	timeNow := time.Date(2017, 11, 3, 1, 12, 42, 1, time.UTC)
	note := Note{
		id:          "any_gen_id",
		name:        "note_name",
		description: "note_desc",
		state:       New,
		created:     timeNow,
		updated:     timeNow,
	}

	assert.Equal(t, "Id: any_gen_id, Name: note_name, Description: note_desc, State: New, Last change: 03 Nov 17 03:12 EET", note.String())
}
