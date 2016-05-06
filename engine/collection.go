package engine

import "github.com/yanyiwu/settledb/schema"

type Collection struct {
	schema *schema.Schema
}

func NewCollection(schema *schema.Schema) *Collection {
	return &Collection{schema}
}
