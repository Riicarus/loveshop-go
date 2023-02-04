package util

import (
	"database/sql/driver"
	"encoding/json"
)

type Void struct{}

type JSONStringSlice []string

func (s JSONStringSlice) Value() (driver.Value, error) {
	b, err := json.Marshal(s)
	return string(b), err
}

func (s *JSONStringSlice) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), s)
}

type JSONMap map[string]interface{}

func (m JSONMap) Value() (driver.Value, error) {
	b, err := json.Marshal(m)
	return string(b), err
}

func (m *JSONMap) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), m)
}
