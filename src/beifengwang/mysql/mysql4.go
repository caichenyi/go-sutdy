package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// 删除
func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		panic(err)
	}

	// 定义好格式，？为占位符
	stmt, err := db.Prepare("DELETE FROM user_info where id=?")
	res, err := stmt.Exec(11)
	id, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Print(id)

}
