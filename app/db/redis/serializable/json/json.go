package json

import (
	"bala/app/db/redis/serializable"
	"encoding/json"
)

type Json struct {
}

func (s *Json) Serialization(value interface{}) ([]byte, error) {
	return json.Marshal(value)
}

func (s *Json) Deserialization(bytes []byte, ref interface{}) error {
	return json.Unmarshal(bytes, ref)
}

func init() {
	serializable.SetInstance(new(Json))
}
