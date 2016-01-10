package engine

import (
	"encoding/json"
	"errors"
)

type Schema struct {
	Fields       []SchemaField  `"json":"fields"`
	NameIDMap    map[string]int `"json":"nameidmap"`
	PrimaryKey   string         `"json":"primarykey"`
	PrimaryKeyID int            `"json":"primarykeyid"`
}

type SchemaField struct {
	Name string `"json": "name"`
	Type string `"json": "type"`
}

func NewSchema(b []byte) (*Schema, error) {
	s := &Schema{
		Fields:    make([]SchemaField, 0),
		NameIDMap: make(map[string]int),
	}
	err := s.parse(b)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Schema) parse(b []byte) error {
	err := json.Unmarshal(b, s)
	if err != nil {
		return err
	}
	for idx, field := range s.Fields {
		s.NameIDMap[field.Name] = idx
		if field.Type == "primarykey" {
			if s.PrimaryKey != "" {
				return errors.New("primary key need to be unique")
			}
			s.PrimaryKey = field.Name
			s.PrimaryKeyID = idx
		}
	}
	if s.PrimaryKey == "" {
		return errors.New("primary key should be set")
	}
	return nil
}

func (s *Schema) GetPrimaryKey() string {
	return s.PrimaryKey
}
