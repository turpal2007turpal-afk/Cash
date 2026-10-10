package request_db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

func AddProductPSQL (ctx context.Context, conn *pgx.Conn, name string, price int64, quantity int) error {
	sqlquery := `
	INSERT INTO product (name, price, quantity)
	VALUES ($1, $2, $3);
		`
	_, err := conn.Exec(ctx, sqlquery, name, price, quantity)
	if err != nil {
		return err
	}

	return nil	
}

func CreateUserAccount (ctx context.Context, conn *pgx.Conn, username string, pass int, balance int) error {
	sqlquery := `
	INSERT INTO users (username, password_hash, balance)
	VALUES ($1, $2, $3);
	`
	_, err := conn.Exec(ctx, sqlquery, username, pass, balance )
	if err != nil {
		return err
	}

	return nil
}

func CreateCashTicket (ctx context.Context, conn *pgx.Conn, name string, price int, quantity int, totalprice int, paymethod string, timeinfo time.Time) error {
	sqlquery := `
	INSERT INTO (name, price, quantity, total_price, payment_method, time_info)
	VALUES ($1, $2, $3, $4, $5, $6);
	`
	_, err := conn.Exec(ctx, sqlquery, name, price, quantity, totalprice, paymethod, timeinfo)
	if err != nil {
		return err
	}
	return nil
}