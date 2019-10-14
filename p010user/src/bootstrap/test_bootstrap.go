package bootstrap

import (
	"os"
	"translate/P10User/src/config"
)
import mLog "translate/P10User/src/lib/log"

type TestBootstrap struct {

}

func (b *TestBootstrap) Boot() {
	gopath := os.Getenv("GOPATH")
	config.LoadEnv(gopath+"/src/go-translate-compose/p010user")
	mLog.Setup()
}
