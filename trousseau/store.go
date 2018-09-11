package trousseau

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type KV interface {
	Set(key, value string) error
	Get(key string) (string, error)
	Delete(key string) error
}

type Store struct {
	Meta    Meta    `json:"meta"`
	Config  Config  `json:"config"`
	Section Section `json:"sections"`

	separator string
}

func Open(filename string) (*Store, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var store *Store
	err = json.Unmarshal(b, store)
	if err != nil {
		return nil, err
	}

	return store, nil
}

func (s *Store) WriteTo(filename string) error {
	b, err := json.Marshal(s)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, b, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) Set(key, value string) error {
	return nil
}

func (s *Store) Get(key string) (string, error) {
	return "", nil
}

func (s *Store) Delete(key string) error {
	return nil
}

func (s *Store) GetSection(key string) *Section {
	return nil
}

func (s *Store) AddSection(key string) error {
	return nil
}

func (s *Store) DeleteSection(key string) error {
	return nil
}

func NewStore() *Store {
	return &Store{
		separator: "/",
	}
}
