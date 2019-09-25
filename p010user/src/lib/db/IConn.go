package db

type IConn interface {
	Init()
	GetConnQuery() string
	GetDriver()string
	GetDatabase() string
}
