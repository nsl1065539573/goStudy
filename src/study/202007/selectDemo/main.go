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

func getConn() *sqlx.DB {
	db, err := sqlx.Open("mysql", "root:nan1065539573@tcp(127.0.0.1:3306)/demo")
	if err != nil {
		fmt.Println("open mysql connection failed", err)
		return nil
	}
	return db
}

func main(){
	DB := getConn()
	if DB == nil {
		fmt.Println("open mysql connection failed")
		return
	}
	defer DB.Close()
	var person []person
	err := DB.Select(&person, "select user_id, username, sex, email from person where user_id=?", 2) 
	if err != nil {
		fmt.Println("select person failed")
	}
	fmt.Println(person[0])
}