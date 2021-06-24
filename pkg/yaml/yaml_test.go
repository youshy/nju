package yaml

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	mockConfigFileOk  = "dir: \"dir\""
	mockConfigFileBad = "0123456789"
	mockConfigPath    = "config_test.yaml"

	modeReadWrite = 0644
	modeWriteOnly = 0200
)

func createFile(path, file string, mode int) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode(mode))
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write([]byte(file))
	if err != nil {
		return err
	}

	return nil
}

func deleteFile(path string) error {
	return os.Remove(path)
}

func TestReadConfig(t *testing.T) {
	t.Run("ReadConfig doesn't return any error on valid path and config.yaml", func(t *testing.T) {
		dir, err := os.Getwd()
		if err != nil {
			t.Errorf("Error on getting work directory %v", err)
		}

		path := filepath.Join(dir, mockConfigPath)
		err = createFile(path, mockConfigFileOk, modeReadWrite)
		if err != nil {
			t.Errorf("Error on test file creation %v", err)
		}
		defer deleteFile(path)

		c, err := ReadConfig(path)
		assert.NoError(t, err)
		assert.NotEmpty(t, c)
		assert.Equal(t, c.Dir, "dir")
	})

	t.Run("ReadConfig returns an error if file does not exists", func(t *testing.T) {
		dir, err := os.Getwd()
		if err != nil {
			t.Errorf("Error on getting work directory %v", err)
		}

		path := filepath.Join(dir, mockConfigPath)

		_, err = ReadConfig(path)
		assert.ErrorIs(t, err, fileIsNotExistError)
	})

	t.Run("ReadConfig returns an error if unable to read the file", func(t *testing.T) {
		dir, err := os.Getwd()
		if err != nil {
			t.Errorf("Error on getting work directory %v", err)
		}

		path := filepath.Join(dir, mockConfigPath)
		err = createFile(path, mockConfigFileOk, modeWriteOnly)
		if err != nil {
			t.Errorf("Error on test file creation %v", err)
		}
		defer deleteFile(path)

		_, err = ReadConfig(path)

		assert.Error(t, err)
	})

	t.Run("ReadConfig returns an error if unable to unmarshal", func(t *testing.T) {
		dir, err := os.Getwd()
		if err != nil {
			t.Errorf("Error on getting work directory %v", err)
		}

		path := filepath.Join(dir, mockConfigPath)
		err = createFile(path, mockConfigFileBad, modeReadWrite)
		if err != nil {
			t.Errorf("Error on test file creation %v", err)
		}
		defer deleteFile(path)

		_, err = ReadConfig(path)

		assert.Error(t, err)
	})
}
