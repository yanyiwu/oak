package engine

import (
	"github.com/golang/glog"
	"github.com/yanyiwu/settledb/index_manager"
)

type Engine struct {
	idxmgr *index_manager.IndexManager
}

func NewEngine(indexdir string) (*Engine, error) {
	idxmgr, err := index_manager.NewIndexManager(indexdir)
	if err != nil {
		glog.Error(err)
		return nil, err
	}
	ng := &Engine{idxmgr}
	return ng, nil
}
