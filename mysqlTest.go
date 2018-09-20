package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@/mysql?charset=utf8")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer db.Close()

	stmt, err := db.Prepare("insert test set id=?,name =?,age=?")
	if err != nil {
		fmt.Println(err)
	}

	res, err := stmt.Exec("7", "wahaha", "20")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	res, err = stmt.Exec("8", "kangshifu", "20")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// res.RowsAffected()
	id, err := res.LastInsertId()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(id)

}
