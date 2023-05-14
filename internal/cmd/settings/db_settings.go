package settings

import (
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type dbCon struct {
	Conn *sql.DB
}

type settings struct {
	engine string
	dsn    string
}

var DbConn dbCon

func init() {
	err := DbConn.CreateConnection()
	if err != nil {
		log.Fatalln(err)
	}
}

func (d *dbCon) CreateConnection() error {
	s := &settings{}
	if err := s.loadEnvironCredentials(); err != nil {
		return err
	}

	conn, err := sql.Open(s.engine, s.dsn)
	if err != nil {
		return err
	}

	err = conn.Ping()
	if err != nil {
		return err
	}

	d.Conn = conn
	return err
}

func (s *settings) loadEnvironCredentials() error {
	engine := os.Getenv("DATABASE_TYPE")
	if engine == "" {
		return errors.New("database_type not defined on environment")
	}

	dsn := os.Getenv("DATABASE_CREDENTIALS")
	if dsn == "" {
		return errors.New("database credentials not defined on environment")
	}

	s.engine = engine
	s.dsn = dsn
	return nil

}
