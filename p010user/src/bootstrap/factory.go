package bootstrap

func New(p Type) IBootstrap{
	return mapping[p]
}
type Type int
const(
	INTEGRATION Type = iota
	TEST
	CMD
)

var mapping = map[Type]IBootstrap {
	INTEGRATION: &IntegrationBootstrap{},
	TEST: &TestBootstrap{},
	CMD: &CmdBootstrap{},
}