package schema

import "testing"

func TestSchema(t *testing.T) {
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
	schema, err := NewSchema(b)
	if err != nil {
		t.Fatal(err)
	}
	if schema.GetPrimaryKey() != "id" {
		t.Fatal()
	}
	if err := schema.Dump(); err == nil {
		t.Fatal()
	}

	s2, err := NewSchemaFromFile("testdata/schema1.json")
	if err != nil {
		t.Fatal(err)
	}
	if s2.GetPrimaryKey() != "id" {
		t.Fatal()
	}
}
