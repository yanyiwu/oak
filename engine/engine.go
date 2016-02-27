package engine

type Engine struct {
	kv map[string]interface{}
}

func NewEngine() *Engine {
	engine := &Engine{
		make(map[string]interface{}),
	}
	return engine
}

func (ng *Engine) GetKeys() []string {
	keys := make([]string, 0, len(ng.kv))
	for k, _ := range ng.kv {
		keys = append(keys, k)
	}
	return keys
}

func (ng *Engine) Get(key string) interface{} {
	return ng.kv[key]
}

func (ng *Engine) Set(key string, value interface{}) error {
	ng.kv[key] = value
	return nil
}

//func (ng *Engine) SetSchema(key string, value *Schema) error {
//	if _, ok := ng.kv[key]; ok {
//		return errors.New("key " + key + " already exists")
//	}
//	ng.kv[key] = value
//	return nil
//}
//
//func (ng *Engine) GetSchema(key string) *Schema {
//	return ng.kv[key].(*Schema)
//}
