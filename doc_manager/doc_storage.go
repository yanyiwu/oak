package doc_manager

import (
	"github.com/golang/glog"
	"github.com/yanyiwu/settledb/schema"
)

type DocStorage struct {
	name    string // collection name
	rootdir string // directory path
	sto     *Storage
	schema  *schema.Schema
}

func NewDocStorage(name string, rootdir string, schema *schema.Schema) (*DocStorage, error) {
	sto, err := NewStorage(rootdir)
	if err != nil {
		glog.Error(err)
		return nil, err
	}
	return &DocStorage{name, rootdir, sto, schema}, nil
}

func (this *DocStorage) Close() {
	this.sto.Close()
}
