package sample

import (
	"database/sql"
	"errors"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Sample struct {
	Name        string        `yaml:"name"`
	Filename    string        `yaml:"filename"`
	Description string        `yaml:"description"`
	Content     string        `yaml:"content"`
	Regex       string        `yaml:"regex"`
	Rows        []interface{} `yaml:"rows"`
	SQL         string        `yaml:"sql"`
	Spec        Spec          `yaml:"spec"`
	Value       float64       `yaml:"value"`
	DB          *sql.DB
	TmpFile     string
	Assignment  string
}

func LoadFromFile(filepath string) (*Sample, error) {
	f, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	return LoadFromBytes(f)
}

func LoadFromBytes(bytes []byte) (*Sample, error) {
	sample := Sample{}
	err := yaml.Unmarshal(bytes, &sample)
	if err != nil {
		return nil, err
	}
	err = sample.CheckConsistency()
	return &sample, err
}

func (s Sample) CheckConsistency() error {
	spec := s.Spec
	if spec.Type == Regex && s.Regex == "" {
		return errors.New("regular expression cannot be empty")
	}
	return nil
}

func (s Sample) Tag() string {
	return s.Name + "-" + s.Assignment
}
