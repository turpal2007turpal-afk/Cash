package postgres_connection

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateTableOrders(ctx context.Context, conn *pgx.Conn) (error) {
	sqlquery := `
	CREATE TABLE IF NOT EXISTS orders (
	id SERIAL PRIMARY KEY,
	name VARCHAR(30) NOT NULL,
	price NUMERIC(10,2) NOT NULL CHECK (price >= 0)
	); 
	`
	_, err := conn.Exec(ctx, sqlquery,)
	if err != nil {
		return err
	}
	return err

}