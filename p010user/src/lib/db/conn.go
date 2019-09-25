package db

type conn struct {
	Username string
	Password string
	Addr string
	Database string
	Driver string
	ConnectionQuery string
}

func (c *conn) Init()   {

}

func (c *conn) GetConnQuery() string{
	if c.ConnectionQuery == "" {
		panic("Conn has not set up, call Init() first")
	}
	return  c.ConnectionQuery
}
func (c *conn) GetDriver() string{
	if c.Driver == "" {
		panic("Conn has not set up, call Init() first")
	}
	return c.Driver
}
func (c *conn) GetDatabase() string {
	if c.Driver == "" {
		panic("Conn has not set up, call Init() first")
	}
	return c.Database
}