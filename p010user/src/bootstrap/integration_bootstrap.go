package bootstrap

import "translate/P10User/src/config"
import mLog "translate/P10User/src/lib/log"

type IntegrationBootstrap struct {

}

func (b *IntegrationBootstrap) Boot() {
	config.LoadEnv("")
	mLog.Setup()
}
