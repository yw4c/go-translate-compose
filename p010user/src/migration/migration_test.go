package migration

import (
	"log"
	"runtime/debug"
	"testing"
	"translate/P10User/src/bootstrap"
	"translate/P10User/src/lib/db"

)

func init() {
	bootstrap.New(bootstrap.TEST).Boot()
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
