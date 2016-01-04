package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func SetupDb() {
	var err error
	db, err = sql.Open("mysql", "root:123456@/user?charset=utf8")
	checkErr(err)
}

func CloseDb() {
	db.Close()
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func AddUser(account string, password string) {
	//插入数据
	stmt, err := db.Prepare("INSERT tb_account SET account=?,password=?")
	checkErr(err)

	res, err := stmt.Exec(account, password)
	checkErr(err)

	userid, err := res.LastInsertId()
	checkErr(err)

	fmt.Println("userid is ", userid)

	stmt, err = db.Prepare("INSERT tb_user SET userid=?")
	checkErr(err)

	_, err = stmt.Exec(userid)
	checkErr(err)
}

func FindUser(account string, password string) {
	stmt, err := db.Prepare("SELECT * FROM tb_account WHERE account=?,password=?")
	checkErr(err)

	//	rows, err := db.Query("SELECT * FROM tb_account")
	res, err := stmt.Exec(account, password)
	checkErr(err)

	fmt.Println(res)
	//	for rows.Next() {
	//		var userid int
	//		var account string
	//		var password string
	//		var created string
	//		err = rows.Scan(&userid, &account, &password)
	//		checkErr(err)
	//		fmt.Println(userid)
	//		fmt.Println(account)
	//		fmt.Println(password)
	//	}
}

func FillInfo(userid, gender, age int, userName string) {
	stmt, err := db.Prepare("update tb_user set username=?,age=?,gender=? where uid=?")
	checkErr(err)

	res, err := stmt.Exec(userName, age, gender, userid)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
}

func ShowUser() {
	rows, err := db.Query("SELECT * FROM tb_user")
	checkErr(err)

	for rows.Next() {
		var userid int
		var username string
		var age int
		var gender int
		err = rows.Scan(&userid, &username, &age, &gender)
		checkErr(err)
		fmt.Println(userid)
		fmt.Println(username)
		fmt.Println(age)
		fmt.Println(gender)
	}
}
