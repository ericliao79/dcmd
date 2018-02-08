package dcmd

import (
	"os"
	"io"
	"fmt"
	"path/filepath"
	"encoding/json"
	"io/ioutil"
)

const (
	Name        = "Dcmd"
	Usage       = "fast do docker-compose something"
	Version     = "0.0.1"
	ConfigName  = "/config.json"
	CheckSymbol = "\u2714 "
	CrossSymbol = "\u2716 "
	EditSymbol  = "\u2710 "
)

var (
	// StorePath is the default dcmd config
	StorePath = filepath.Join(os.Getenv("HOME"), ".dcmd")
)

type Config struct {
	PATH string `json:"dockerPath"`
}

func IsEmpty(path string) (bool, error) {
	f, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdirnames(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err
}

func SetConfig(p string) (int, error) {
	f, err := os.Create(StorePath + ConfigName)
	defer f.Close()
	if err != nil {
		fmt.Println("1")
		return 0, err
	}

	config := Config{
		PATH: p,
	}

	b, err := json.Marshal(config)
	if err != nil {
		return 0, err
	}
	s, err := f.Write(b)
	if err != nil {
		return 0, err
	}

	return s, nil
}

func LoadComposes() map[string]string {
	keys := map[string]string{}
	var c Config

	raw, err := ioutil.ReadFile(StorePath + ConfigName)
	if err != nil {
		fmt.Println(err.Error())
	}

	json.Unmarshal(raw, &c)

	list, error := ioutil.ReadDir(c.PATH)
	if error != nil {
		fmt.Println(error.Error())
	}

	for _, f := range list {
		if f.IsDir() {
			keys[f.Name()] = f.Name()
		}
	}

	return keys
}
