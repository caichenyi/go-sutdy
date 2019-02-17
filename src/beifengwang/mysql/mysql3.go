package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// 查询
func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		panic(err)
	}

	// 定义好格式，？为占位符
	rows, err := db.Query("SELECT * FROM user_info where id=11")
	if err != nil {
		panic(err)
	}

	// 输出多行
	for rows.Next() {
		var id int
		var username string
		var department string
		var create_time string
		err = rows.Scan(&id, &username, &department, &create_time)
		fmt.Print(id, username, department, create_time)
	}
}
