package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type DbConn struct {
	DbPool *sql.DB
}

func InitPool(config *DbConfig) *DbConn {
	var db *sql.DB

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.dbHost, config.dbPort, config.dbUser, config.dbPassword, config.dbName)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Panicf("error opening database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Panicf("error pinging database: %v", err)
	}

	if config.dbMigrationPath != nil {
		if err := runMigration(db, *config.dbMigrationPath); err != nil {
			log.Panicf("error pinging database: %v", err)
		}
	}

	return &DbConn{
		DbPool: db,
	}
}

func (db *DbConn) ClosePool() error {
	if db.DbPool != nil {
		return db.DbPool.Close()
	}
	return nil
}

func runMigration(db *sql.DB, path string) error {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", path),
		"mysql",
		driver,
	)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
