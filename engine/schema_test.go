package engine

import "testing"

func TestSchema(t *testing.T) {
	b := []byte(`{
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
	}`)
	schema, err := NewSchema(b)
	if err != nil {
		t.Fatal(err)
	}
	if schema.GetPrimaryKey() != "id" {
		t.Fatal()
	}
}
