package main

import (
	"GitandPSQL/postgres_connection"
	"context"
	"fmt"
)

func main () {
	ctx := context.Background()

	conn, err := postgres_connection.CreateConnection(ctx)
	if err != nil {
		panic(err)
	}

	if err := postgres_connection.CreateTableOrders(ctx, conn); err != nil {
		panic(err)
	}
	

	fmt.Println("success!")

}