package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

type Connector struct {
	connection *sqlx.DB
}

func NewConnector(name, password, host string) (*Connector, error) {
	source := fmt.Sprintf("%s:%s@%s", name, password, host)
	db, err := sqlx.Open("mysql", source)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(5)
	db.SetConnMaxLifetime(10 * time.Second)
	db.SetMaxIdleConns(10)
	db.SetConnMaxIdleTime(10 * time.Second)
	return &Connector{connection: db}, nil
}

func (c Connector) Close() {
	c.connection.Close()
}
