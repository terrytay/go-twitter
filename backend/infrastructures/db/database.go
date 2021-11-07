package db

import (
	"log"
)

type Database struct {
	log    *log.Logger
	Client DBClient
}

type Data interface {
	GetId() string
}

func NewDatabase(l *log.Logger, c DBClient) *Database {
	return &Database{l, c}
}
