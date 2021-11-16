package config

import (
	"bytes"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/evenlab/go-kit/drive"
	"github.com/evenlab/go-kit/errors"
	"gopkg.in/yaml.v3"
)

type (
	// Config represents the main app's configuration.
	Config struct {
		Host              string        `yaml:"host"`
		Port              int           `yaml:"port"`
		MaxConnectionIdle time.Duration `yaml:"max_connection_idle"`
	}
)

// NewConfig constructs an empty configuration.
func NewConfig() *Config {
	return &Config{}
}

// Init creates configs form defaults or reads
// existing yamlFileName in app root path.
func (c *Config) Init() error {
	filePath := filepath.Join(OsAppRootPath(), yamlFileName)
	if !isFileExists(filePath) { // create from default
		log.Println("Create config form default into: " + filePath)
		c.flushToDefault()
		if err := c.writeToFile(filePath); err != nil {
			return errors.WrapErr("create config failed", err)
		}
	} else {
		log.Println("Read config from file: " + filePath)
		if err := c.readFromFile(filePath); err != nil {
			return errors.WrapErr("read config failed", err)
		}
	}

	return nil
}

// flushToDefault flushes Config to the default values.
func (c *Config) flushToDefault() {
	*c = Config{
		Host:              "localhost",
		Port:              4141,
		MaxConnectionIdle: 60 * time.Second,
	}
}

// readFromFile reads Config from the provided file path.
func (c *Config) readFromFile(path string) error {
	f, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return errors.WrapErr("open config failed", err)
	}
	defer func(f *os.File) { _ = f.Close() }(f)

	if err := yaml.NewDecoder(f).Decode(&c); err != nil {
		return errors.WrapErr("decode config failed", err)
	}

	return nil
}

// writeToFile creates default Config and writes it to the provided file path.
func (c *Config) writeToFile(path string) error {
	dir := filepath.Dir(path)
	if err := drive.MakeDirs(dir); err != nil {
		return errors.WrapErr("create dir failed", err)
	}

	buf := bytes.NewBuffer(nil)
	enc := yaml.NewEncoder(buf)
	if err := enc.Encode(c); err != nil {
		return errors.WrapErr("encode config failed", err)
	}
	defer func(enc *yaml.Encoder) { _ = enc.Close() }(enc)

	if err := os.WriteFile(path, buf.Bytes(), drive.DefaultFilePerm); err != nil {
		return errors.WrapErr("write config failed", err)
	}

	return nil
}
