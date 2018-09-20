package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type student struct {
	id    string
	name  string
	age   int
	grade int
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/belajar-golang")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func sqlPrepare() {
	db, err := connect()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("select name, grade from tb_student where id =?")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var result = student{}
	stmt.QueryRow("1").Scan(&result.name, &result.grade)
	fmt.Printf("name:%s\ngrade: %d\n", result.name, result.grade)

}

func main() {
	sqlPrepare()
}
