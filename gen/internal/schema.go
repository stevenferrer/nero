package internal

type Schema struct {
	Typ   *Typ
	Ident *Col
	Cols  []*Col
	Pkg   string
}
