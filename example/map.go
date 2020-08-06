package example

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/pkg/errors"
)

type Map map[string]string

func (m Map) Value() (driver.Value, error) {
	b, err := json.Marshal(m)
	if err != nil {
		return nil, errors.Wrap(err, "marshal map")
	}

	return b, nil
}

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
