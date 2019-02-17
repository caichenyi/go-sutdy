package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// 更新
func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		panic(err)
	}

	// 定义好格式，？为占位符
	stmt, err := db.Prepare("UPDATE user_info SET username=? where id=?")
	res, err := stmt.Exec("zhangsan", 11)
	id, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println(id)
}
