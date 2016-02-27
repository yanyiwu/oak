package storage

import (
	"encoding/json"

	"github.com/yanyiwu/settledb/Godeps/_workspace/src/github.com/syndtr/goleveldb/leveldb"
)

type Storage struct {
	db *leveldb.DB
}

func NewStorage(dirpath string) (*Storage, error) {
	db, err := leveldb.OpenFile(dirpath, nil)
	if err != nil {
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
