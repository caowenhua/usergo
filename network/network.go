package network

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"../db"
)

func print(w http.ResponseWriter, r *http.Request) {
	//解析参数，默认是不会解析的
	r.ParseForm()
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
		s := db.AddUser(account, password)
		fmt.Fprintf(w, s.ToString())
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	print(w, r)
	account, password, err := getAccount(w, r)
	fmt.Println(account, ",", password)
	if err == nil {
		s := db.Login(account, password)
		fmt.Fprintf(w, s.ToString())
	}
}

func showUser(w http.ResponseWriter, r *http.Request) {
	print(w, r)
	page := 1
	pageCount := 10
	var err error
	for k, v := range r.Form {
		if k == "page" {
			s := strings.Join(v, "")
			page, err = strconv.Atoi(s)
			if err != nil {
				fmt.Println("page cannot to int")
				page = 1
			}
		} else if k == "pageCount" {
			s := strings.Join(v, "")
			pageCount, err = strconv.Atoi(s)
			if err != nil {
				fmt.Println("pageCount cannot to int")
				pageCount = 10
			}
		}
	}
	if page < 1 {
		page = 1
	}
	if pageCount < 1 {
		pageCount = 10
	}
	s := db.ShowUser(page, pageCount)
	fmt.Fprintf(w, s.ToString())

}

func Listen() {
	db.SetupDb()
	http.HandleFunc("/user/reg", register)
	http.HandleFunc("/user/login", login)
	http.HandleFunc("/account/showuser", showUser)
	//	http.HandleFunc("/", handle)             //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func Close() {
	db.CloseDb()
}
