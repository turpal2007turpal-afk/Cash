package connect_db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func Connect(ctx context.Context) (*pgx.Conn,error) {
	return pgx.Connect(ctx, "postgres://postgres:1985@localhost:5432/postgres")
}