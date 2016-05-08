package schema

import (
	"errors"
	"io/ioutil"

	"github.com/golang/glog"
)

type Schema struct {
	nameIDMap    map[string]int
	primaryKey   string
	primaryKeyID int
	filepath     string

	// load info from json string
	si *SchemaInfo
}

func NewSchema(jsonstr string) (*Schema, error) {
	si, err := NewSchemaInfo(jsonstr)
	if err != nil {
		glog.Info(err)
		return nil, err
	}
	s := &Schema{
		make(map[string]int),
		"",
		0,
		"",
		si,
	}
	if err := s.parse(si); err != nil {
		glog.Error(err)
		return nil, err
	}
	return s, nil
}

func NewSchemaFromFile(filepath string) (*Schema, error) {
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		glog.Error(err)
		return nil, err
	}
	return NewSchema(string(b))
}

func (s *Schema) parse(si *SchemaInfo) error {
	for idx, field := range si.Fields {
		s.nameIDMap[field.Name] = idx
		if field.Type == "primarykey" {
			if s.primaryKey != "" {
				return errors.New("primary key need to be unique")
			}
			s.primaryKey = field.Name
			s.primaryKeyID = idx
		}
	}
	if s.primaryKey == "" {
		return errors.New("primary key should be set")
	}
	return nil
}

func (s *Schema) GetPrimaryKey() string {
	return s.primaryKey
}

func (s *Schema) Dump() error {
	if s.filepath == "" {
		return errors.New("filepath has not been set")
	}
	return s.si.Dump(s.filepath)
}
