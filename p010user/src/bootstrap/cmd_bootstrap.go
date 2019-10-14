package bootstrap
import (
	"go/build"
	"os"
	"translate/P10User/src/config"
)
import mLog "translate/P10User/src/lib/log"

type CmdBootstrap struct {

}

func (b *CmdBootstrap) Boot() {
	gopath := os.Getenv("GOPATH")
	if gopath=="" {
		gopath = build.Default.GOPATH
	}

	config.LoadEnv(os.Getenv("GOPATH")+"/src/go-translate-compose/p010user")
	mLog.Setup()
}