package engine

import (
	"os"
	"path"
	"testing"
)

func TestEngine(t *testing.T) {
	dir := path.Join(os.TempDir(), "settledbengine")
	defer os.RemoveAll(dir)
	_, err := NewEngine(dir)
	if err != nil {
		t.Fatal(err)
	}
}
