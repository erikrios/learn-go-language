package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"os"
	"strconv"
	"time"
)

// PostgreSQLConnection func for connection to PostgreSQL database.
func PostgreSQLConnection() (*sqlx.DB, error) {
	// Define database connection settings.
	maxConn, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))
	maxIdleConn, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
	maxLifeTimeConn, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME_CONNECTIONS"))

	// Define database connection for PostgreSQL.
	db, err := sqlx.Connect("pgx", os.Getenv("DB_SERVER_URL"))
	if err != nil {
		return nil, fmt.Errorf("error, not connected to database, %w", err)
	}

	// SSet database connection settings.
	db.SetMaxOpenConns(maxConn)                           // the default is 0 (unlimited)
	db.SetMaxIdleConns(maxIdleConn)                       // defaultMaxIdleConns = 2
	db.SetConnMaxLifetime(time.Duration(maxLifeTimeConn)) // 0, connections are reused forever

	// Try to ping database.
	if err := db.Ping(); err != nil {
		defer func(db *sqlx.DB) {
			err := db.Close()
			if err != nil {
				fmt.Printf("error, failed to close the connection %v", err)
			}
		}(db) // close database connection
		return nil, fmt.Errorf("error, not sent ping to database %w", err)
	}
	return db, nil
}
