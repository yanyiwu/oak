package engine

import "testing"

func TestEngine(t *testing.T) {
	ng := NewEngine()
	if ng.Get("hello") != nil {
		t.Fatal()
	}
	v := "world"
	err := ng.Set("hello", v)
	if err != nil {
		t.Fatal()
	}
	if "world" != ng.Get("hello").(string) {
		t.Fatal()
	}
}
