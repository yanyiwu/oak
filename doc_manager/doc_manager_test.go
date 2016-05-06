package doc_manager

import (
	"os"
	"path"
	"testing"
)

func TestDocManager(t *testing.T) {
	dir := path.Join(os.TempDir(), "settledbdocmanager")
	defer os.RemoveAll(dir)
	_, err := NewDocManager(dir)
	if err != nil {
		t.Fatal(err)
	}
}
