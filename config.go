package trousseau

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
)

type Loader interface {
	Load(data []byte) error
	LoadFrom(r io.Reader) error
}

type Dumper interface {
	Dump() ([]byte, error)
	DumpTo(w io.Writer) error
}

type LoadDumper interface {
	Loader
	Dumper
}

const DEFAULT_TROUSSEAU_CONFIGURATION_NAME = ".trousseauconfig"

type Configuration struct {
	Authentication Authentication `json:"authentication"`
	Server         string         `json:"server"`
}

type Authentication struct {
	Username string `json:"username"`
	Token    Token  `json:"token"`
}

func (c *Configuration) Load(data []byte) error {
	err := json.Unmarshal(data, c)
	if err != nil {
		return fmt.Errorf("unable to deserialize configuration: %v\n", err)
	}

	return nil
}

func (c *Configuration) LoadFrom(r io.Reader) error {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return fmt.Errorf("unable to read file buffer content: %v\n", err)
	}

	err = c.Load(data)
	if err != nil {
		return err
	}

	return nil
}

func (c *Configuration) Dump() ([]byte, error) {
	data, err := json.Marshal(c)
	if err != nil {
		return nil, fmt.Errorf("unable to serialize configuration: %v\n", err)
	}

	return data, nil
}

func (c *Configuration) DumpTo(w io.Writer) error {
	data, err := c.Dump()
	if err != nil {
		return err
	}

	_, err = w.Write(data)
	if err != nil {
		return fmt.Errorf("unable to write configuration data: %v\n", err)
	}

	return nil
}
