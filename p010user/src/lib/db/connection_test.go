package db

import (
	"testing"
	"translate/P10User/src/bootstrap"

)
func init() {
	bootstrap.New(bootstrap.TEST).Boot()
}

func TestConn(t *testing.T) {

	NewConn(&LocalConn{})

}
