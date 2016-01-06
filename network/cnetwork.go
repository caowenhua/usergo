package network

import (
	"fmt"
	"net/http"

	"github.com/Centny/gwf/routing"

	"me.user/db"
)

func Cregister(hs *routing.HTTPSession) routing.HResult {
	//	fmt.Println(hs.W, *hs.R, hs.S, *hs.Mux)
	//	fmt.Println(*hs.Mux)
	var account string
	var password string
	err := hs.ValidCheckVal(`
		account,R|S,L:0~50;
		password,R|S,L:0~50;
		`, &account, &password)
	if err == nil {
		s, err := db.CAddUser(account, password)
		if err != nil {
			return hs.MsgResErr2(1, "account is existed", err)
		} else {
			return hs.MsgRes(s)
		}
	} else {
		return hs.MsgResErr2(1, "config error", err)
	}
}

func Clogin(hs *routing.HTTPSession) routing.HResult {
	var account string
	var password string
	err := hs.ValidCheckVal(`
		account,R|S,L:0~50;
		password,R|S,L:0~50;
		`, &account, &password)
	if err != nil {
		return hs.MsgResErr2(1, "config error", err)
	} else {
		s, err := db.CLogin(account, password)
		if err == nil {
			return hs.MsgRes(s)
		} else {
			return hs.MsgResErr2(1, "invalid account or password", err)
		}
	}
}

func CshowUser(hs *routing.HTTPSession) routing.HResult {
	var page int64
	var pageCount int64
	err := hs.ValidCheckVal(`
		page,O|I,R:0~9999999;
		pageCount,O|I,R:0~500;
		`, &page, &pageCount)
	if err != nil {
		return hs.MsgResErr2(1, "config error", err)
	} else {
		s := db.CShowUser(page, pageCount)
		return hs.MsgRes(s)
	}
}

func CfillInfo(hs *routing.HTTPSession) routing.HResult {
	var username string
	var userid int64
	var age int64
	var gender int64
	err := hs.ValidCheckVal(`
		username,O|S,L:0~50;
		userid,R|I,R:0;
		age,O|I,R:0~3000;
		gender,O|I,R:-1~2;
		`, &username, &userid, &age, &gender)
	if err != nil {
		return hs.MsgResErr2(1, "config error", err)
	} else {
		s, err := db.CFillInfo(userid, age, gender, username)
		if err != nil {
			return hs.MsgResErr2(1, "no such user", err)
		} else {
			return hs.MsgRes(s)
		}
	}
}

func CdeleteUser(hs *routing.HTTPSession) routing.HResult {
	var userid int64
	err := hs.ValidCheckVal(`
		userid,R|I,R:0;
		`, &userid)
	if err != nil {
		return hs.MsgResErr2(1, "config error", err)
	} else {
		s := db.CDeleteUser(userid)
		return hs.MsgRes(s)
	}
}

func CListen() {
	db.CSetupDb()
	fmt.Println("listen begin")
	sb := routing.NewSrvSessionBuilder("", "/", "example", 60*60*1000, 10000)
	mux := routing.NewSessionMux("", sb)
	mux.HFunc("/user/reg", Cregister)
	mux.HFunc("/user/login", Clogin)
	mux.HFunc("/user/delete", CdeleteUser)
	mux.HFunc("/account/showuser", CshowUser)
	mux.HFunc("/user/fill", CfillInfo)
	//	mux.HFunc("^/api/list(\\?.*)?$", List)
	//	mux.HFunc("^/api/query(\\?.*)?$", Query)
	fmt.Println(http.ListenAndServe(":9090", mux))
}

func CClose() {
	db.CCloseDb()
}
