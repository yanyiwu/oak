package schema

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/golang/glog"
)

type SchemaInfo struct {
	Fields []SchemaField `"json":"fields"`
}

type SchemaField struct {
	Name string `"json": "name"`
	Type string `"json": "type"`
}

func NewSchemaInfo(jsonstr string) (*SchemaInfo, error) {
	si := &SchemaInfo{}
	err := json.Unmarshal([]byte(jsonstr), si)
	if err != nil {
		return nil, err
	}
	return si, nil
}

func NewSchemaInfoFromFile(filepath string) (*SchemaInfo, error) {
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		glog.Error(err)
		return nil, err
	}
	return NewSchemaInfo(string(b))
}

func (this *SchemaInfo) Load(filepath string) error {
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		glog.Error(err)
		return err
	}
	if err := json.Unmarshal(b, this); err != nil {
		glog.Error(err)
		return err
	}
	return nil
}

func (this *SchemaInfo) Dump(filepath string) error {
	b, err := json.Marshal(this)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(filepath, b, os.ModePerm); err != nil {
		return err
	}
	return nil
}
