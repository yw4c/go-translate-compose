package db

import (
	"testing"
	"translate/P10User/src/lib/utest"
)
func init() {
	utest.Init()
}

func TestConn(t *testing.T) {

	db := NewConn(&LocalConn{})
	if err := db.DB().Ping(); err != nil {
		t.Log(err.Error())
		t.Fail()
	}
}
