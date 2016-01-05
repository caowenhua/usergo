package db

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	"../bean"

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

func AddUser(account string, password string) (result bean.Response) {
	isIn := IsUserExist(account, password)
	if isIn {
		result = bean.Response{"error", "account is existed", ""}
	} else {
		//插入数据
		stmt, err := db.Prepare("INSERT tb_account SET account=?,password=?")
		checkErr(err)

		res, err := stmt.Exec(account, password)
		checkErr(err)

		userid, err := res.LastInsertId()
		checkErr(err)

		fmt.Println("userid is ", userid)

		stmt, err = db.Prepare("INSERT tb_user SET userid=?, username=?")
		checkErr(err)

		_, err = stmt.Exec(userid, account)
		checkErr(err)

		user, err := findUserById(userid)
		checkErr(err)

		//		js, _ := json.Marshal(user)
		//		var buff bytes.Buffer
		//		buff.Write(js)
		//		result = bean.Response{"ok", "success", buff.String()}
		result = bean.Response{"ok", "success", user}
	}
	return

}

func Login(account string, password string) (result bean.Response) {
	isIn, userid := IsUserExistWithId(account, password)
	if !isIn {
		result = bean.Response{"error", "invalid account or password", ""}
	} else {
		user, err := findUserById(userid)
		checkErr(err)
		//		js, _ := json.Marshal(user)
		//		var buff bytes.Buffer
		//		buff.Write(js)
		//		result = bean.Response{"ok", "success", buff.String()}
		result = bean.Response{"ok", "success", user}
	}
	return
}

func IsUserExistWithId(account string, password string) (bool, int64) {
	var buffer bytes.Buffer
	buffer.WriteString("SELECT * FROM tb_account WHERE account='")
	buffer.WriteString(account)
	buffer.WriteString("' and password='")
	buffer.WriteString(password)
	buffer.WriteString("'")
	//	stmt, err := db.Prepare("SELECT * FROM tb_account WHERE account=? and password=?")
	//	checkErr(err)

	fmt.Println(buffer.String())
	rows, err := db.Query(buffer.String())
	//	res, err := stmt.Exec(account, password)
	checkErr(err)

	fmt.Println(rows)
	index := 0
	var userid int64
	for rows.Next() {
		var ac string
		var pw string
		err = rows.Scan(&userid, &ac, &pw)
		checkErr(err)
		//		fmt.Println(userid)
		//		fmt.Println(account)
		//		fmt.Println(password)
		index++
	}
	if index == 0 {
		return false, 0
	} else {
		return true, userid
	}
}

func IsUserExist(account string, password string) bool {
	var buffer bytes.Buffer
	buffer.WriteString("SELECT * FROM tb_account WHERE account='")
	buffer.WriteString(account)
	buffer.WriteString("' and password='")
	buffer.WriteString(password)
	buffer.WriteString("'")
	//	stmt, err := db.Prepare("SELECT * FROM tb_account WHERE account=? and password=?")
	//	checkErr(err)

	fmt.Println(buffer.String())
	rows, err := db.Query(buffer.String())
	//	res, err := stmt.Exec(account, password)
	checkErr(err)

	fmt.Println(rows)
	index := 0
	for rows.Next() {
		var userid int
		var account string
		var password string
		err = rows.Scan(&userid, &account, &password)
		checkErr(err)
		//		fmt.Println(userid)
		//		fmt.Println(account)
		//		fmt.Println(password)
		index++
	}
	rows.Close()
	if index == 0 {
		return false
	} else {
		return true
	}
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

func ShowUser(page, pageCount int) (result bean.Response) {
	var buffer bytes.Buffer
	buffer.WriteString("SELECT * FROM tb_user LIMIT ")
	buffer.WriteString(strconv.Itoa((page - 1) * pageCount))
	buffer.WriteString(", ")
	buffer.WriteString(strconv.Itoa(pageCount))
	rows, err := db.Query(buffer.String())
	checkErr(err)

	userSlice := make([]bean.User, 0)
	//	var buffer bytes.Buffer

	for rows.Next() {
		var username string
		var userid int
		var age int
		var gender int
		err = rows.Scan(&username, &userid, &age, &gender)
		checkErr(err)

		u := bean.User{username, userid, age, gender}
		userSlice = append(userSlice, u)

		//		buffer.WriteString("\n\n")
		//		buffer.WriteString("\nuserId: ")
		//		buffer.WriteString(strconv.Itoa(userid))
		//		buffer.WriteString("\nuserName: ")
		//		buffer.WriteString(username)
		//		buffer.WriteString("\nage: ")
		//		buffer.WriteString(strconv.Itoa(age))
		//		buffer.WriteString("\ngender: ")
		//		buffer.WriteString(strconv.Itoa(gender))
		//		buffer.WriteString("\n\n")
	}
	rows.Close()

	//	js, _ := json.Marshal(userSlice)
	//	var buff bytes.Buffer
	//	buff.Write(js)
	//	result = bean.Response{"ok", "success", buff.String()}
	result = bean.Response{"ok", "success", userSlice}
	//	result = buffer.String()
	return
}

func findUserById(userid int64) (user bean.User, err error) {
	var buffer bytes.Buffer
	buffer.WriteString("SELECT * FROM tb_user WHERE userid='")
	buffer.WriteString(strconv.FormatInt(userid, 10))
	buffer.WriteString("'")

	fmt.Println(buffer.String())
	rows, e := db.Query(buffer.String())
	checkErr(e)

	if rows.Next() {
		var username string
		var userid int
		var age int
		var gender int
		err = rows.Scan(&username, &userid, &age, &gender)
		checkErr(err)
		user = bean.User{username, userid, age, gender}
	} else {
		err = errors.New("user is not existed")
	}
	return
}
