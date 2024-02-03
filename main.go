package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Person struct {
	Id   int
	Name string
	Age  int
}

func main() {
	//engine := gin.Default()
	//连接数据库
	connStr := "root:123456@tcp(127.0.0.1:3306)/gin_sql"
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	//_, err = db.Exec("create table if not exists person if not exist(" +
	//	"id int auto_increment primary key," +
	//	"name varchar(20)," +
	//	"age int default 1)")
	//if err != nil {
	//	log.Fatal(err.Error())
	//	return
	//}

	_, err = db.Exec("insert into person(name,age)"+
		"values(?,?)", "szh", 20)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	//查询数据库操作
	rows, err := db.Query("select * from person")
	if err != nil {
		log.Fatal(err.Error())
		return
	}

scan:
	if rows.Next() {
		var person Person
		err = rows.Scan(&person.Id, &person.Name, &person.Age)
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		fmt.Println(person.Id, person.Name, person.Age)
		goto scan
	}

	//engine.Run()
}
