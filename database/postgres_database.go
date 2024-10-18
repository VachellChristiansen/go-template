package database

import (
	"context"
	"go-template/helper"
	"os"

	"github.com/jackc/pgx/v5"
)

type PostgresDatabase interface {
	Open()
	GetConn() *pgx.Conn
	Close()
}

type PostgresDatabaseImpl struct {
	log  helper.LoggerHelper
	conn *pgx.Conn
}

func NewPostgresDatabase() PostgresDatabase {
	log := helper.NewLoggerHelper()
	return &PostgresDatabaseImpl{
		log: log,
	}
}

func (db *PostgresDatabaseImpl) Open() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("POSTGRES_URL"))
	if err != nil {
		db.log.LogErrAndExit(2, err, "Failed opening postgres database connection")
	}

	db.conn = conn
}

func (db *PostgresDatabaseImpl) GetConn() *pgx.Conn {
	return db.conn
}

func (db *PostgresDatabaseImpl) Close() {
	if db.conn != nil {
		db.conn.Close(context.Background())
	}
}
