package yaml

import (
	"errors"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"

	"github.com/youshy/nju/pkg/types"
	"gopkg.in/yaml.v2"
)

var (
	fileIsNotExistError = errors.New("File does not exists")
	configPath          = "/.config/nju/config.yaml"
)

func GetDefaultPath() string {
	usr, err := user.Current()
	// that should never error, hence panic as it'd point to some big issues with the system
	if err != nil {
		panic(err)
	}

	path := filepath.Join(usr.HomeDir, configPath)

	return path
}

func ReadConfig(path string) (types.Config, error) {
	t := types.Config{}

	ok := validatePath(path)
	if !ok {
		return t, fileIsNotExistError
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return t, err
	}

	err = yaml.Unmarshal(data, &t)
	if err != nil {
		return t, err
	}

	return t, nil
}

func validatePath(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
