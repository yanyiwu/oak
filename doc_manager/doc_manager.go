// DocManager manage document data by collection name
package doc_manager

import (
	"errors"
	"path"

	"github.com/golang/glog"
	"github.com/yanyiwu/settledb/util"
)

type DocManager struct {
	manager map[string]*Storage
}

func NewDocManager(dirpath string) (*DocManager, error) {
	util.CreateDirIfNotExists(dirpath)
	dirs := util.GetDirsFromDir(dirpath)
	glog.Info("get dirs ", dirs)
	mp := make(map[string]*Storage)
	for _, dir := range dirs {
		collection := path.Base(dir)
		sto, err := NewStorage(dir)
		if err != nil {
			glog.Error(err)
			return nil, err
		}
		_, ok := mp[collection]
		if ok {
			err := errors.New("collection " + collection + " already exists")
			return nil, err
		}
		mp[collection] = sto
	}
	return &DocManager{mp}, nil
}

func (this *DocManager) Create(collection string) {
}
