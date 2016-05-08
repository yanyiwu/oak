// DocManager manage document data by collection name
package doc_manager

import (
	"errors"
	"path"
	"strings"

	"github.com/golang/glog"
	"github.com/yanyiwu/settledb/schema"
	"github.com/yanyiwu/settledb/util"
)

// DocManager.manager:
// key: collection nanme
// value: the pointer of DocStorage
type DocManager struct {
	manager map[string]*DocStorage
	rootdir string
}

func NewDocManager(dirpath string) (*DocManager, error) {
	util.CreateDirIfNotExists(dirpath)
	dirs := util.GetDirsFromDir(dirpath)
	glog.Info("get dirs ", dirs)
	docmgr := &DocManager{make(map[string]*DocStorage), dirpath}
	for _, dir := range dirs {
		schemafile := strings.TrimRight(dir, "/") + ".json"
		collection := path.Base(dir)
		s, err := schema.NewSchema(schemafile)
		if err != nil {
			glog.Error(err)
			return nil, err
		}
		if err := docmgr.Create(collection, s); err != nil {
			glog.Error(err)
			return nil, err
		}
	}
	return docmgr, nil
}

func (this *DocManager) Close() {
	//TODO
}

func (this *DocManager) Create(collection string, schema *schema.Schema) error {
	_, ok := this.manager[collection]
	if ok {
		err := errors.New("collection " + collection + " already exists")
		return err
	}
	sto, err := NewDocStorage(
		collection,
		path.Join(this.rootdir, collection),
		schema,
	)
	if err != nil {
		glog.Error(err)
		return err
	}
	this.manager[collection] = sto
	return nil
}
