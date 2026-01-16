package database

import (
	"errors"
	"fmt"
	"go-config/utils"
	"time"
)

type Options struct {
	Host            string
	Port            string
	User            string
	Password        string
	DatabaseName    string
	Protocol        string
	ConnMaxLifetime time.Duration
	MaxIdleConns    int
	MaxOpenConns    int
	PARAM           string
}

type Database interface {
	Open(options Options)
	Close()
	Get() interface{}
	Ping() error
}

func BuildDns(options Options) (string, error) {
	handleError := func(msg string, err error) (string, error) {
		return "", errors.New(msg)

	}
	if utils.IsBlank(options.Host) {
		return handleError("Host is required", nil)
	}
	if utils.IsBlank(options.Port) {
		return handleError("Port is required", nil)
	}
	if utils.IsBlank(options.User) {
		return handleError("User is required", nil)
	}
	if utils.IsBlank(options.Password) {
		return handleError("Password is required", nil)
	}
	if utils.IsBlank(options.DatabaseName) {
		return handleError("DatabaseName is required", nil)
	}
	if utils.IsBlank(options.Protocol) {
		return handleError("Protocol is required", nil)
	}

	if utils.IsBlank(options.PARAM) {
		return handleError("PARAM is required", nil)
	}
	dns := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", options.Host, options.Port, options.User, options.Password, options.DatabaseName, options.Protocol)
	return dns, nil
}
