package database

import (
	"database/sql"
	"errors"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type OrmDB struct {
	OrmInstance *gorm.DB
	Database    Database
}

func OpenOrmWithDatabase(database Database) (*OrmDB, error) {

	ormDb := OrmDB{}
	if database == nil {
		return nil, errors.New("database is nil")
	}
	var err error

	d := database.Get().(*sql.DB)
	gormDb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: d,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	ormDb.Database = database
	ormDb.OrmInstance = gormDb

	return &ormDb, nil
}

func OpenOrm(host string, port string, user string, password string, database string) (*OrmDB, error) {
	postgresDb, err := OpenPostgresSqlDatabase(host, port, user, password, database)
	if err != nil {
		return nil, err
	}
	return OpenOrmWithDatabase(postgresDb)
}
