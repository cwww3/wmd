package config

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	AppName    = "wmd"
	AppDir     = ".wmd"
	ConfigName = "config.yaml"
	DataName   = "data.json"
)

type LocalImagePath string

var UserHomeDir string

var SupportUploader = []string{"github"}

func (p *LocalImagePath) Init() error {
	if len(*p) == 0 {
		*p = LocalImagePath(filepath.Join(UserHomeDir, AppName))
	}
	return nil
}
func (p *LocalImagePath) Save() error {
	_, err := os.Stat(string(*p))
	if err != nil {
		if os.IsNotExist(err) {
			return os.MkdirAll(string(*p), 0755)
		}
		return err
	}
	return nil
}

type Picgo struct {
	PicBed PicBed `json:"picBed" `
}

type Config struct {
	LocalImagePath LocalImagePath `json:"localImagePath"`
	Enable         bool           `json:"enable"`
	Picgo          Picgo          `json:"picgo" yaml:"-"`
}

func (c *Config) GetImagePath() string {
	return string(c.LocalImagePath)
}

func (c *Config) Load() error {
	err := c.LocalImagePath.Init()
	if err != nil {
		return err
	}

	err = c.Picgo.Load()
	if err != nil {
		return err
	}

	// userhome/.wmd/config.yaml
	var data []byte

	config := filepath.Join(UserHomeDir, AppDir, ConfigName)
	_, err = os.Stat(config)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	data, err = ioutil.ReadFile(config)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, c)
	if err != nil {
		return err
	}

	return err
}
func (c *Config) Save() error {
	err := c.LocalImagePath.Save()
	if err != nil {
		return err
	}
	err = c.Picgo.Save()
	if err != nil {
		return err
	}

	config := filepath.Join(UserHomeDir, AppDir, ConfigName)
	data, err := yaml.Marshal(c)
	if err != nil {
		return err
	}
	return os.WriteFile(config, data, 0755)
}

func (p *Picgo) Load() error {
	// userhome/.wmd/data.json
	var data []byte
	config := filepath.Join(UserHomeDir, AppDir, DataName)
	_, err := os.Stat(config)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	data, err = ioutil.ReadFile(config)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, p)
	return err
}
func (p *Picgo) Save() error {
	if _, err := os.Stat(filepath.Join(UserHomeDir, AppDir)); err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(filepath.Join(UserHomeDir, AppDir), 0755)
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}

	config := filepath.Join(UserHomeDir, AppDir, DataName)
	data, err := json.Marshal(p)
	if err != nil {
		return err
	}
	return os.WriteFile(config, data, 0755)
}

type PicBed struct {
	Uploader string `json:"uploader"`
	Github   Github `json:"github"`
	Proxy    Proxy  `json:"proxy"`
}

type Github struct {
	Repo      string `json:"repo"`
	Branch    string `json:"branch"`
	Token     string `json:"token"`
	Path      string `json:"path"`
	CustomUrl string `json:"customUrl"`
}

type Proxy struct {
	Proxy string `json:"proxy"`
}
