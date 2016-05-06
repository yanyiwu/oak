package doc_manager

import "reflect"

type Document struct {
	Fields []string `json:"fields"`
}

func (this *Document) Equal(doc *Document) bool {
	return reflect.DeepEqual(this.Fields, doc.Fields)
}
