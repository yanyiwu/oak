package doc_manager

import (
	"encoding/json"

	"github.com/golang/glog"
	"github.com/syndtr/goleveldb/leveldb"
)

type Storage struct {
	db *leveldb.DB
}

func NewStorage(dirpath string) (*Storage, error) {
	db, err := leveldb.OpenFile(dirpath, nil)
	if err != nil {
		glog.Error(err)
		return nil, err
	}
	return &Storage{
		db,
	}, nil
}

func (this *Storage) Close() {
	this.db.Close()
}

func (this *Storage) Put(key string, doc *Document) error {
	b, err := json.Marshal(doc)
	if err != nil {
		return err
	}
	return this.db.Put([]byte(key), b, nil)
}

func (this *Storage) Get(key string) (*Document, error) {
	b, err := this.db.Get([]byte(key), nil)
	if err != nil {
		return nil, err
	}
	doc := Document{make([]string, 0)}
	if err := json.Unmarshal(b, &doc); err != nil {
		return nil, err
	}
	return &doc, nil
}
