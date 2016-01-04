package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"../db"
)

func print(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	fmt.Println("=================================================")
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path:", r.URL.Path)
	fmt.Println("scheme:", r.URL.Scheme)
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Println("=================================================")
}

func getAccount(w http.ResponseWriter, r *http.Request) (account, password string, err error) {
	account = ""
	password = ""
	for k, v := range r.Form {
		if k == "account" {
			account = strings.Join(v, "")
		} else if k == "password" {
			password = strings.Join(v, "")
		}
	}
	if account == "" || password == "" {
		fmt.Fprintf(w, "用户名或密码不能为空")
		err = errors.New("user account or password is null.")
	}
	return
}

func register(w http.ResponseWriter, r *http.Request) {
	print(w, r)
	account, password, err := getAccount(w, r)
	fmt.Println(account, ",", password)
	if err == nil {
		//TODO 插入到数据库
		fmt.Fprintf(w, "注册")
		db.AddUser(account, password)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	print(w, r)
	account, password, err := getAccount(w, r)
	fmt.Println(account, ",", password)
	if err == nil {
		//TODO 查表
		fmt.Fprintf(w, "登录")
		db.FindUser(account, password)
	}
}

func showUser(w http.ResponseWriter, r *http.Request) {
	print(w, r)
	fmt.Fprintf(w, "showUser")
}

func main() {
	db.SetupDb()
	http.HandleFunc("/user/reg", register)
	http.HandleFunc("/user/login", login)
	http.HandleFunc("/account", showUser)
	//	http.HandleFunc("/", handle)             //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
