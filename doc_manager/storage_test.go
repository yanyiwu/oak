package doc_manager

import (
	"os"
	"path/filepath"
	"testing"
)

func TestStorage(t *testing.T) {
	dirpath := filepath.Join(os.TempDir(), "settledbstoragetest")
	if _, err := os.Stat(dirpath); os.IsExist(err) {
		t.Fatal(dirpath + " already exists, maybe you need to remove it manually")
	}
	defer os.RemoveAll(dirpath)
	st, err := NewStorage(dirpath)
	if err != nil {
		t.Fatal(err)
	}
	defer st.Close()

	key := "1"
	doc := &Document{[]string{}}
	if err := st.Put(key, doc); err != nil {
		t.Fatal(err)
	}

	doc2, err := st.Get(key)
	if err != nil {
		t.Fatal(err)
	}
	if !doc.Equal(doc2) {
		t.Fatal()
	}
}
