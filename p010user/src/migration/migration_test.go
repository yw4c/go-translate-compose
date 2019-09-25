package migration

import (
	"log"
	"runtime/debug"
	"testing"
	"translate/P10User/src/lib/db"
	"translate/P10User/src/lib/utest"
)

func init() {
	utest.Init()
}

func TestMigrate(t *testing.T) {
	defer func() {
		if r:= recover();r != nil {
			log.Println(r)
			log.Println("[panic]  trace : ", string(debug.Stack()))
		}
	}()
	m := New(&db.LocalConn{})
	m.Up()
	m.Rollback()

}
