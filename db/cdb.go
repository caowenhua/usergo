package db

import (
	"database/sql"
	"errors"
	"fmt"

	"me.user/bean"

	"github.com/Centny/gwf/dbutil"
	_ "github.com/go-sql-driver/mysql"
)

var cdb *sql.DB

func GetDb() *sql.DB {
	return cdb
}

func CSetupDb() {
	var err error
	cdb, err = sql.Open("mysql", "root:123456@/user?charset=utf8")
	if err != nil {
		panic(err.Error())
	}

}

func CCloseDb() {
	cdb.Close()
}

func CAddUser(account string, password string) (user bean.User, err error) {
	isIn := CIsUserExist(account, password)
	if isIn {
		err = errors.New("account is existed")
	} else {
		userid, err := dbutil.DbInsert(cdb, "insert into tb_account (account,password) values(?,?)", account, password)
		checkErr(err)

		_, err = dbutil.DbInsert(cdb, "insert into tb_user (userid,username) values(?,?)", userid, account)
		checkErr(err)

		user, err = cfindUserById(userid)
		checkErr(err)
	}
	return

}

func CLogin(account string, password string) (user bean.User, err error) {
	isIn, userid := CIsUserExistWithId(account, password)
	if !isIn {
		err = errors.New("invalid account or password")
	} else {
		user, err = cfindUserById(userid)
		checkErr(err)
	}
	return
}

func CIsUserExistWithId(account string, password string) (bool, int64) {
	ac := []bean.Account{}
	err := dbutil.DbQueryS(cdb, &ac, "select * from tb_account where account=? and password=?", account, password)
	if err == nil && len(ac) > 0 {
		return true, ac[0].UserId
	} else {
		return false, 0
	}
}

func CIsUserExist(account string, password string) bool {
	ac := []bean.Account{}
	err := dbutil.DbQueryS(cdb, &ac, "select * from tb_account where account=? and password=?", account, password)

	if err == nil && len(ac) > 0 {
		return true
	} else {
		return false
	}
}

func CFillInfo(userid, age, gender int64, userName string) (user bean.User, err error) {
	_, err = dbutil.DbUpdate(cdb, "update tb_user set username=?,age=?,gender=? where userid=?", userName, age, gender, userid)
	if err == nil {
		user, err = cfindUserById(userid)
	}
	return
}

func CShowUser(page, pageCount int64) (userSlice []bean.User) {
	//	m, _ := dbutil.DbQuery(cdb, "select * from tb_user limit ?,?", (page-1)*pageCount, pageCount)
	err := dbutil.DbQueryS(cdb, &userSlice, "select * from tb_user limit ?,?", (page-1)*pageCount, pageCount)
	checkErr(err)
	return
}

func CDeleteUser(userid int64) (string, error) {
	i, err := dbutil.DbUpdate(cdb, "delete from tb_user where userid = ?", userid)
	if err == nil {
		if i == 0 {
			return "no such user", err
		}
		d, e := dbutil.DbUpdate(cdb, "delete from tb_account where userid = ?", userid)
		if d == 0 {
			return "no such user", err
		} else {
			return "success", e
		}
	}
	return "config error", err
}

func cfindUserById(userid int64) (user bean.User, err error) {
	userSlice := []bean.User{}
	err = dbutil.DbQueryS(cdb, &userSlice, "select * from tb_user where userid=?", userid)
	if len(userSlice) > 0 {
		user = userSlice[0]
	} else {
		err = errors.New("no such user")
	}
	return
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
