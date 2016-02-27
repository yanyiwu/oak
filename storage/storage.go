package storage

import "github.com/syndtr/goleveldb/leveldb"

type Storage struct {
	db *leveldb.DB
}

func NewStorage() (*Storage, error) {
	db, err := leveldb.OpenFile(dirpath, nil)
	if err != nil {
		return nil, err
	}
	return &Storage{
		db,
	}, nil
}

func (this *Storage) Close() {
	db.Close()
}

func (this *Storage) Put(doc *Document) {
	err := db.Put()
}
