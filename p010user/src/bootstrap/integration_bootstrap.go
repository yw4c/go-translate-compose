package bootstrap

import (
	"os"
	"translate/P10User/src/config"
)
import mLog "translate/P10User/src/lib/log"

type IntegrationBootstrap struct {

}

func (b *IntegrationBootstrap) Boot() {
	config.LoadEnv(os.Getenv("GOPATH")+"/src/p010user")
	mLog.Setup()
}
