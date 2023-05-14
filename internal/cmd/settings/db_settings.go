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

	createSchemesIfNotExists(DbConn.Conn)
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

func createSchemesIfNotExists(db *sql.DB) {
	rows, err := db.Query("SELECT EXISTS (SELECT FROM information_schema.tables WHERE table_name = $1)", "transactions")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var exists bool
	for rows.Next() {
		if err := rows.Scan(&exists); err != nil {
			log.Println(err)
		}
	}

	if !exists {
		_, err = db.Exec("CREATE TABLE transactions (id SERIAL PRIMARY KEY, type VARCHAR(1) NOT NULL, date DATE NOT NULL, value FLOAT NOT NULL, cpf VARCHAR(11) NOT NULL, card VARCHAR(12) NOT NULL, time TIME NOT NULL, owner VARCHAR(14) NOT NULL, market VARCHAR(19) NOT NULL)")
		if err != nil {
			log.Println(err)
		}
	}

}
