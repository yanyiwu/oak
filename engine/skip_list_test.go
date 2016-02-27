package engine

import "testing"

func TestSkipList(t *testing.T) {
	list := NewSkipList()
	if list.Len() != 0 {
		t.Fatal()
	}
}
