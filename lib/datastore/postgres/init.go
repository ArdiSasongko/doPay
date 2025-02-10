package postgres

import (
	"fmt"

	"github.com/ArdiSasongko/doPay/lib/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Database struct {
	cfg      *config.Config
	DBClient DBInterface
}

// initialization postgres
func InitDB(cfg *config.Config) *Database {
	return &Database{
		cfg: cfg,
	}
}

func (d *Database) InitiateConnect() (DBInterface, error) {
	cfg := d.cfg.Database
	connectDB := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.DBName,
		cfg.SSLMode,
	)

	return d.connect(connectDB)
}

func (d *Database) connect(conn string) (DBInterface, error) {
	cfg := d.cfg.Database
	db, err := sqlx.Connect("postgres", conn)
	if err != nil {
		return nil, err
	}

	if cfg.MaxOpenConns != -1 {
		db.SetMaxOpenConns(cfg.MaxOpenConns)
	}

	if cfg.MaxIdleConns != -1 {
		db.SetMaxIdleConns(cfg.MaxIdleConns)
	}

	return db, nil
}
