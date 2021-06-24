package yaml

import (
	"io/ioutil"
	"os"
	"os/user"

	"github.com/youshy/pkg/types"
	"gopkg.in/yaml.v2"
)

func GetDefaultPath() string {
	usr, err := user.Current()
	// that should never error, hence panic as it'd point to some big issues with the system
	if err != nil {
		panic(err)
	}

	path := usr.HomeDir + "/.config/nju/config.yaml"

	return path
}

func ValidatePath(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func ReadConfig(path string) (types.Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	t := types.Config{}

	err = yaml.Unmarshal(data, &t)
	if err != nil {
		return nil, err
	}

	return t, nil
}
