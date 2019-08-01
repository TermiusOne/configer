package configer

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// New creates new configer.
func New(path string, config interface{}) *Configer {
	return &Configer{
		FilePath: path,
		Config:   config,
	}
}

// Configer can create and read JSON configs from path.
type Configer struct {
	FilePath string
	Config   interface{}
}

// Read config from path.
func (c *Configer) Read() error {
	f, err := ioutil.ReadFile(c.FilePath)
	if err != nil {
		return err
	}

	if !json.Valid(f) {
		return errors.New("Invalid JSON in" + c.FilePath)
	}

	return json.Unmarshal(f, c.Config)
}

// Create creates config file if it not exists.
func (c *Configer) Create() error {
	if _, err := os.Stat(filepath.Dir(c.FilePath)); os.IsNotExist(err) {
		err := os.Mkdir(filepath.Dir(c.FilePath), 0755)
		if err != nil {
			return err
		}
	}

	if _, err := os.Stat(c.FilePath); os.IsNotExist(err) {
		f, err := os.Create(c.FilePath)
		if err != nil {
			return err
		}

		defer f.Close()

		jm, err := json.MarshalIndent(c.Config, "", "\t")
		if err != nil {
			return err
		}

		_, err = f.Write(jm)
		if err != nil {
			return err
		}

		log.Printf("Config created. Add data to %v.", c.FilePath)

	} else {
		log.Printf("Config exists. Check data in %v.", c.FilePath)
	}
	return nil
}

// ReadConfig reads and nmarshal JSON file in you struct or other valid type.
func ReadConfig(path string, v interface{}) error {
	var c = Configer{FilePath: path, Config: v}
	return c.Read()
}

// CreateConfig creates all path and the JSON config file (if it not exists)
// from you struct or other valid type.
func CreateConfig(path string, v interface{}) error {
	var c = Configer{FilePath: path, Config: v}
	return c.Create()
}
