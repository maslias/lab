package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

func NewDb(
	driver string,
	addr string,
	maxOpenConns, maxIdleConns int,
	maxIdleTime,
	timeout time.Duration,
) (*sql.DB, error) {
	db, err := sql.Open(driver, addr)
	if err != nil {
		return nil, err
	}


	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxIdleTime(maxIdleTime)

	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	fmt.Printf("database is connected: %s \n", driver)

	return db, nil
}
