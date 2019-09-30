package migration

import (
	"database/sql"
	"github.com/spf13/viper"
	"log"
	"os"
	mDB "translate/P10User/src/lib/db"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)
type IMigration interface {
	Up()
	Rollback()
}
type MigrationImpl struct {
	migrate *migrate.Migrate
}

func New(conn mDB.IConn) *MigrationImpl{

	conn.Init()

	// "user:password@tcp(host:port)/dbname?multiStatements=true"
	db, err := sql.Open(conn.GetDriver(), conn.GetConnQuery())
	if err != nil {
		log.Panic(err.Error())
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Panic(err.Error())
	}

	path := os.Getenv("GOPATH")+viper.GetString("APP_PATH")+"src/migration/migration"
	m, err := migrate.NewWithDatabaseInstance(
		"file://"+path,
		conn.GetDatabase(),
		driver,
	)
	if err != nil {
		log.Panic(err.Error())
	}

	return &MigrationImpl{
		migrate:m,
	}
}

func (this *MigrationImpl) Up() {
	if err := this.migrate.Up(); err != nil {
		if err != migrate.ErrNoChange {
			log.Panicf(err.Error())
		}
	}
}

func (this *MigrationImpl) Rollback() {
	if err := this.migrate.Steps(-1); err != nil {
		log.Panicf(err.Error())
	}
}

