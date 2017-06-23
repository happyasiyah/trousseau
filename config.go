package trousseau

import "io"

type JSONSerializableFileReadWriter interface {
	Read(r io.Reader) ([]byte, error)
	Write(data []byte, w io.Writer) error
	Serialize(v interface{}) ([]byte, error)
	Deserialize(data []byte, v interface{}) error
}

type Authentication struct {
	Username string `json:"username"`
	Token    Token  `json:"token"`
}

type Configuration struct {
	Authentication Authentication `json"authentication"`
}

func (c *Configuration) Serialize(v interface{}) ([]byte, error) {
	return nil, nil
}

func (c *Configuration) Deserialize(data []byte, v interface{}) error {
	return nil
}

func (c *Configuration) Read(r io.Reader) ([]byte, error) {
	return nil, nil
}

func (c *Configuration) Write(r io.Writer) error {
	return nil
}

const DEFAULT_TROUSSEAU_CONFIGURATION_NAME = ".trousseauconfig"
