package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		panic(err)
	}

	// 定义好格式，？为占位符
	stmt, err := db.Prepare("INSERT user_info SET username=?,departname=?,create_time=?")
	res, err := stmt.Exec("zhangsan", "技术部", "2012-12-09")
	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println(id)

}
