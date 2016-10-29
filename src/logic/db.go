package logic

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Conn struct {
	db *sql.DB
}

var conn *Conn = new(Conn)

func CreateDBConn() *Conn {
	return conn
}

func (conn *Conn) Init(host string, port string, username string, password string, database string) bool {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, database)
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		Addlog(err.Error())
		return false
	}

	conn.db = db

	return conn.Ping() == nil
}

func (conn *Conn) Ping() error {
	err := conn.db.Ping()
	if err != nil {
		Addlog(err.Error())
	}

	return err
}

func (conn *Conn) Query(sql string) (*sql.Rows, error) {
	err := conn.Ping()
	if err == nil {
		return conn.db.Query(sql)
	} else {
		return nil, err
	}
}
