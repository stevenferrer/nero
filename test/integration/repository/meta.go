package repository

// Collection is the name of the collection
const Collection = "users"

// Column is a User column
type Column int

// String implements Stringer
func (c Column) String() string {
	switch c {
	case ColumnID:
		return "id"
	case ColumnUID:
		return "uid"
	case ColumnEmail:
		return "email"
	case ColumnName:
		return "name"
	case ColumnAge:
		return "age"
	case ColumnGroup:
		return "group"
	case ColumnKv:
		return "kv"
	case ColumnUpdatedAt:
		return "updated_at"
	case ColumnCreatedAt:
		return "created_at"
	}

	return "invalid"
}

const (
	ColumnID Column = iota
	ColumnUID
	ColumnEmail
	ColumnName
	ColumnAge
	ColumnGroup
	ColumnKv
	ColumnUpdatedAt
	ColumnCreatedAt
)
