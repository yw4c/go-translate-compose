package utest

import "translate/P10User/src/config"
import mLog "translate/P10User/src/lib/log"

func Init() {
	config.LoadEnv("$GOPATH/src/go-translate-compose/p010user")
	mLog.Setup()
}

