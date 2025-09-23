package belajar_golang_database

import (
	"context"
	"fmt"
	"testing"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	fmt.Println("Success connect to database")

	ctx := context.Background()

	script := "INSERT INTO customer(id, name) VALUES('regi', 'REGI')"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

