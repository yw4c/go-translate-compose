package bootstrap

import "translate/P10User/src/config"
import mLog "translate/P10User/src/lib/log"

type TestBootstrap struct {

}

func (b *TestBootstrap) Boot() {
	config.LoadEnv("$GOPATH/src/go-translate-compose/p010user")
	mLog.Setup()
}
