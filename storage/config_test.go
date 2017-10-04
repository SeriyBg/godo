package storage

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_ReadingConfigFile(t *testing.T) {
	os.Setenv("GODO_CONFIG", "../testdata/.godoconfig")
	config := readConfig()
	assert.Equal(t, GodoConfig{FilePath: "testFilePath"}, config)
}
