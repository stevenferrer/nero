package internal

type Schema struct {
	Coln  string
	Typ   *Typ
	Ident *Col
	Cols  []*Col
	Pkg   string
}
