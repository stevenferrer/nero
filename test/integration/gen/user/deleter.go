// Code generated by nero, DO NOT EDIT.
package user

type Deleter struct {
	collection string
	pfs        []PredFunc
}

func NewDeleter() *Deleter {
	return &Deleter{
		collection: collection,
	}
}

func (d *Deleter) Where(pfs ...PredFunc) *Deleter {
	d.pfs = append(d.pfs, pfs...)
	return d
}
