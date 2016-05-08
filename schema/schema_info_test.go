package schema

import (
	"fmt"
	"os"
	"path"
	"reflect"
	"testing"
)

func ExampleSchemaInfo() {
	b := `{
	    "fields": [
			{
				"name": "id", 
				"type": "primarykey"
			},
			{
				"name": "title"
			},
			{
				"name": "content"
			}
		]
	}`
	si, _ := NewSchemaInfo(b)
	fmt.Println(si)

	// Output:
	// &{[{id primarykey} {title } {content }]}
}

func ExampleSchemaInfoFromFile() {
	si, _ := NewSchemaInfoFromFile("testdata/schema1.json")
	fmt.Println(si)

	// Output:
	// &{[{id primarykey} {title } {content }]}
}

func TestSchemaInfo(t *testing.T) {
	si, err := NewSchemaInfoFromFile("testdata/schema1.json")
	if err != nil {
		t.Fatal(err)
	}
	tempfile := path.Join(os.TempDir(), "testschemainfo")
	defer os.Remove(tempfile)
	si.Dump(tempfile)
	si2, err := NewSchemaInfo("{}")
	if err != nil {
		t.Fatal(err)
	}
	if reflect.DeepEqual(si, si2) {
		t.Fatal()
	}
	si2.Load(tempfile)
	if !reflect.DeepEqual(si, si2) {
		t.Fatal()
	}
}
