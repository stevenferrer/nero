package example

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/pkg/errors"
)

// Map is a serializeable map
type Map map[string]string

// Value implements driver.Valuer
func (m Map) Value() (driver.Value, error) {
	b, err := json.Marshal(m)
	if err != nil {
		return nil, errors.Wrap(err, "marshal map")
	}

	return b, nil
}

// Scan implements sql.Scanner
func (m *Map) Scan(v interface{}) error {
	if v == nil {
		*m = Map{}
		return nil
	}

	b, ok := v.([]byte)
	if !ok {
		return errors.New("map is not of type []byte")
	}

	err := json.Unmarshal(b, m)
	if err != nil {
		return errors.Wrap(err, "unmarshal map")
	}

	return nil
}
