package mem

import "testing"

func TestDeque(t *testing.T) {
	dq := NewDeque()
	if dq.Size() != 0 {
		t.Fatal()
	}
	x := 0
	dq.PushBack(x)
	if dq.Size() != 1 {
		t.Fatal()
	}
	y := dq.Get(0).(int)
	if x != y {
		t.Fatal()
	}

	for i := 0; i < BLOCK_CAP; i++ {
		dq.PushBack(i + 1)
	}

	if dq.Size() != BLOCK_CAP+1 {
		t.Fatal()
	}

	for i := 0; i < dq.Size(); i++ {
		if dq.Get(i).(int) != i {
			t.Fatal()
		}
	}
}
