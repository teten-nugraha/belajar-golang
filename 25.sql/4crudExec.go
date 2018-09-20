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

func sqlExec() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec("insert into tb_student values (?,?,?,?)", "2", "Yasna", 1, 100)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("insert success")

	_, err = db.Exec("update tb_student set age = ? where id =?", 2, "1")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("update success")

	_, err = db.Exec("delete from tb_student where id = ?", "2")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("delete success")
}

func main() {
	sqlExec()
}
