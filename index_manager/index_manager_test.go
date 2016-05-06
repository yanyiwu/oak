package index_manager

import (
	"os"
	"path"
	"testing"
)

func TestIndexManager(t *testing.T) {
	dir := path.Join(os.TempDir(), "settledbindexmanager")
	defer os.RemoveAll(dir)
	_, err := NewIndexManager(dir)
	if err != nil {
		t.Fatal(err)
	}
}
