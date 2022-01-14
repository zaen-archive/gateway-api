package proxy

import (
	"encoding/json"

	"github.com/InVisionApp/conjungo"
)

type Meta struct {
	body map[string]interface{}
}

func (meta *Meta) bodyToBytes() []byte {
	val, _ := json.Marshal(meta.body)
	return val
}

func (meta *Meta) appendBodyBytes(m []byte, deep bool) error {
	var v map[string]interface{}
	json.Unmarshal(m, &v)

	if deep {
		return conjungo.Merge(&meta.body, v, nil)
	} else {
		return meta.appendBody(v)
	}
}

func (meta *Meta) appendBody(m map[string]interface{}) error {
	if meta.body == nil {
		meta.body = m
		return nil
	}

	for k, v := range m {
		meta.body[k] = v
	}

	return nil
}
