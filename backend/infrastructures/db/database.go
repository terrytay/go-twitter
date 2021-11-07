package db

type Database struct {
	Client DBClient
}

type Data interface {
	GetId() string
}

func NewDatabase(c DBClient) *Database {
	return &Database{c}
}
