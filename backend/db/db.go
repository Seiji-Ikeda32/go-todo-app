package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func OpenDB() *sql.DB {
	driver := os.Getenv("DRIVER")
	dsn := os.Getenv("DSN")

	db, err := sql.Open(driver, dsn)
	if err != nil {
		fmt.Println("faild open database", err)
	} else {
		fmt.Println("success open database")
	}
	return db
}
