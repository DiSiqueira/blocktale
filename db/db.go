package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

const maxConnectAttempts = 5

// NewDB returns
func NewDB(dsn string, maxOpen, maxIdle, attempt int) (*gorp.DbMap, error) {
	errfmt := "NewDB: %s"

	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf(errfmt, err)
	}

	if err := conn.Ping(); err != nil {
		if attempt < maxConnectAttempts {
			logrus.Errorf(errfmt, err)
			time.Sleep(time.Second * attempt)
			return NewDB(dsn, maxOpen, maxIdle, attempt+1)
		}

		return nil, fmt.Errorf(errfmt, err)
	}

	conn.SetMaxOpenConns(maxOpen)
	conn.SetMaxIdleConns(maxIdle)

	mapDB := &gorp.DbMap{
		Db: conn,
		Dialect: gorp.MySQLDialect{
			Engine:   "InnoDB",
			Encoding: "UTF8",
		},
	}

	mapDB.AddTableWithName(Tale{}, "tale").SetKeys(false, "ID")

	return mapDB, nil
}
