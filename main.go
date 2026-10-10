package main

import (
	"GitandPSQL/db_sql/connect_db"
	"GitandPSQL/db_sql/create_table"
	"context"
	"fmt"
)

func main () {
	ctx := context.Background

	// Cоздаем таблицы

	conn, err := connect_db.Connect(ctx())
	if err != nil {
		panic(err)
	}
	// Таблица кассового чека
	if err := create_table.CreateTableCashTicket(ctx(), conn); err != nil {
		panic(err)
	}
	// Таблица продуктов
	if err := create_table.CreateTableProduct(ctx(), conn); err != nil {
		panic(err)
	}
	// Таблица аккаунта пользователя
	if err := create_table.CreateTableUser(ctx(), conn); err != nil {
		panic(err)
	}

	fmt.Println("succed!")




}