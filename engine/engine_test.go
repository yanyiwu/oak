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

//func TestEngineSetSchema(t *testing.T) {
//	b := []byte(`{
//		"book": {
//			"fields": [
//				{
//					"name": "id",
//					"type": "primarykey"
//				},
//				{
//					"name": "title"
//				},
//				{
//					"name": "content"
//				}
//			]
//		}
//	}
//	`)
//	var m map[string]interface{}
//	if err := json.Unmarshal(b, &m); err != nil {
//		t.Fatal(err)
//	}
//	if len(m) != 1 {
//		t.Fatal()
//	}
//	v, ok := m["book"]
//	if !ok {
//		t.Fatal()
//	}
//	b, _ = json.Marshal(v)
//	schema, err := NewSchema(b)
//	if err != nil {
//		t.Fatal(err)
//	}
//	if schema.GetPrimaryKey() != "id" {
//		t.Fatal()
//	}
//	if len(schema.Fields) != 3 {
//		t.Fatal()
//	}
//
//	ng := NewEngine()
//	if err := ng.SetSchema("book", schema); err != nil {
//		t.Fatal(err)
//	}
//	if err := ng.SetSchema("book", schema); err == nil {
//		t.Fatal(err)
//	}
//}
