package create_table

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateTableProduct (ctx context.Context, conn *pgx.Conn) (error) {
	sqlquery := `
	CREATE TABLE IF NOT EXISTS product (
	id SERIAL PRIMARY KEY,
	name VARCHAR(25) NOT NULL,
	price NUMERIC(10,2) NOT NULL CHECK (price >= 0),
	quantity INTEGER NOT NULL DEFAULT 0 CHECK (quantity >= 0)
	);
	`
	_, err := conn.Exec(ctx, sqlquery)
	if err != nil {
		return err
	}
	return nil
} 

func CreateTableCashTicket (ctx context.Context, conn *pgx.Conn) (error) {
	sqlquery := `
	CREATE TABLE IF NOT EXISTS cashticket (
	id SERIAL PRIMARY KEY,
	name VARCHAR(25) NOT NULL,
	quantity INTEGER DEFAULT 0 NOT NULL,
	price NUMERIC(10,2) NOT NULL,
	total_price NUMERIC(10,2) NOT NULL,
	payment_method VARCHAR(25) NOT NULL,
	time_info TIMESTAMP NOT NULL
	);
	`
	_, err := conn.Exec(ctx, sqlquery)
	if err != nil {
		return err
	}

	return nil
}

func CreateTableUser (ctx context.Context, conn *pgx.Conn) error {
 sqlquery := `
 CREATE TABLE IF NOT EXISTS users (
 id SERIAL PRIMARY KEY,
 username VARCHAR(100) NOT NULL UNIQUE,
 password_hash TEXT NOT NULL,
 balance NUMERIC(10,2) NOT NULL,
 avatar_url TEXT,
 created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
 );
 `
 _, err := conn.Exec(ctx, sqlquery)
 if err != nil {
	return err
 }
 return nil
}