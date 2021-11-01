package lib

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var Conn DB

type DB struct {
	Connection *sql.DB
}

// DBInit initialize db connection
func DBInit() {
	dbType := os.Getenv("DB_TYPE")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	sslmode := os.Getenv("DB_SSL")
	schema := os.Getenv("DB_SCHEMA")
	opt := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s search_path=%s",
		host,
		port,
		dbName,
		user,
		password,
		sslmode,
		schema)
	var err error
	// opt = deploy()
	if Conn.Connection, err = sql.Open(dbType, opt); err != nil {
		log.Fatalf("Connection error\n %s", err)
	}
}

func deploy() string {
	return os.Getenv("DATABASE_URL")
}

// Close connection close
func (o *DB) Close() {
	o.Connection.Close()
}

// GetRow get DB Data by query
func (o *DB) GetRow(query string, args ...interface{}) ([][]interface{}, error) {
	var result [][]interface{}
	tx, err := o.startTransaction()
	if err != nil {
		return nil, err
	}
	rows, err := tx.Query(query, args...)
	defer tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	for rows.Next() {
		cols, err := rows.Columns()
		if err != nil {
			return nil, err
		}
		row := make([]interface{}, len(cols))
		ptr := make([]interface{}, len(cols))
		for i := 0; i < len(cols); i++ {
			ptr[i] = &row[i]
		}
		rows.Scan(ptr...)
		result = append(result, row)
	}
	return result, nil
}

// Exec execute data insert update by query
func (o *DB) Exec(query string, args ...interface{}) error {
	tx, err := o.startTransaction()
	defer tx.Commit()
	if err != nil {
		return err
	}
	_, err = tx.Exec(query, args...)
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (o *DB) startTransaction() (*sql.Tx, error) {
	tx, err := o.Connection.Begin()
	if err != nil {
		err = fmt.Errorf("transaction begin Error: \n%s", err)
	}

	return tx, err
}