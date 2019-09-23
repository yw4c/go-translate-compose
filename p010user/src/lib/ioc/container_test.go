package ioc

import (
	"testing"
)


func TestGetContainerInstance(t *testing.T)  {
	c := GetContainerInstance()
	c.Invoke(func(test ITest) {
		if test.echo() != "success" {
			t.Fail()
		}
	})
}