package database

import (
	"database/sql"
	"fmt"
	"go-config/config"
	"time"
)

type PostgresDatabase struct {
	db *sql.DB
}

func OpenPostgresSqlDatabase(host string, port string, user string, password string, database string) (*PostgresDatabase, error) {
	postgresDb := &PostgresDatabase{
		db: nil,
	}
	connMaxLifetime := config.Default().GetInt("db.postgres.connMaxLifetime")
	maxIdleConns := config.Default().GetInt("db.postgres.maxIdleConns")
	maxOpenConns := config.Default().GetInt("db.postgres.maxOpenConns")
	protocol := config.Default().GetString("db.postgres.protocol")
	params := config.Default().GetString("db.postgres.params")
	postgresDb.Open(Options{
		Host:            host,
		Port:            port,
		User:            user,
		Password:        password,
		DatabaseName:    database,
		Protocol:        protocol,
		PARAM:           params,
		ConnMaxLifetime: time.Duration(connMaxLifetime),
		MaxIdleConns:    maxIdleConns,
		MaxOpenConns:    maxOpenConns,
	})
	return postgresDb, nil
}

func (m *PostgresDatabase) Open(options Options) {
	dbs, err := BuildDns(options)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Opening connection on host: %s, port: %s, database: %s", options.Host, options.Port, options.DatabaseName)
	db, err := sql.Open("postgres", dbs)
	if err != nil {
		panic(err)
	}
	fmt.Sprintf("Connection opened on host: %s, port: %s, database: %s", options.Host, options.Port, options.DatabaseName)
	db.SetConnMaxLifetime(options.ConnMaxLifetime)
	db.SetMaxIdleConns(options.MaxIdleConns)
	db.SetMaxOpenConns(options.MaxOpenConns)
	m.db = db
}

func (m *PostgresDatabase) Close() {
	if m.db != nil {
		m.db.Close()
		fmt.Println("Connection closed")
	} else {
		fmt.Println("Connection already closed")
	}
}

func (m *PostgresDatabase) Get() interface{} {
	if m.db == nil {
		panic("Database connection is not open")
	}
	return m.db
}

func (m *PostgresDatabase) Ping() error {
	if m.db == nil {
		panic("Database connection is not open")
	}
	err := m.db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection pinged")
	return nil
}
