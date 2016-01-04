package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func SetupDb() {
	var err error
	db, err = sql.Open("mysql", "root:123456@/user")
	checkErr(err)
}

func CloseDb() {
	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func AddUser(account string, password string) {
	//插入数据
	stmt, err := db.Prepare("INSERT tb_account SET account=?,password=?")
	checkErr(err)

	res, err := stmt.Exec(account, password)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
}

func FindUser(account string, password string) {

}
