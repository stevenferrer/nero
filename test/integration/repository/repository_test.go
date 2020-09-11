package repository

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type errTx struct{}

func (et *errTx) Commit() error {
	return nil
}

func (et *errTx) Rollback() error {
	return errors.New("tx error")
}

type okTx struct{}

func (ot *okTx) Commit() error {
	return nil
}

func (ot *okTx) Rollback() error {
	return nil
}

func Test_rollback(t *testing.T) {
	e := &errTx{}
	err := rollback(e, errors.New("an error"))
	assert.Equal(t, "rollback error: tx error: an error", err.Error())

	o := &okTx{}
	err = rollback(o, errors.New("an error"))
	assert.Equal(t, "an error", err.Error())
}
