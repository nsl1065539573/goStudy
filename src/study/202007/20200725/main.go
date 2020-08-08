package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"fmt"
)

type person struct{
	UserId   int    `db:"user_id"`
    Username string `db:"username"`
    Sex      string `db:"sex"`
    Email    string `db:"email"`
}

type Place struct {
    Country string `db:"country"`
    City    string `db:"city"`
    TelCode int    `db:"telcode"`
}

var Db *sqlx.DB

func getConn() *sqlx.DB {
	database, err := sqlx.Open("mysql", "root:nan1065539573@tcp(127.0.0.1:3306)/demo")
	if err != nil {
		fmt.Println("open mysql failed", err)
		return nil
	}
	return database
}


func main() {
	DB := getConn()
	if DB == nil {
		fmt.Println("open mysql failed")
		return
	}
	defer DB.Close()
	res, err := DB.Exec("insert into person(username, sex, email) values(?, ?, ?)", "Tom", "male", "1065539573")
	if err != nil {
		fmt.Println("insert to `person` failed")
		return
	}
	fmt.Println(res)
}