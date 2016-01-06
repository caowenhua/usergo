package network

import (
	"fmt"
	"github.com/Centny/gwf/util"
	"strconv"
	"testing"
	"time"
)

func TestCListen(t *testing.T) {
	go CListen()
}

func TestCregister(t *testing.T) {
	ti := time.Now()
	for i := 0; i < 10; i++ {
		str := "http://localhost:9090/user/reg?account=%v&password=%v"
		name := "c" + strconv.Itoa(ti.Nanosecond()) + strconv.Itoa(i)
		fmt.Println(name)
		vmap, err := util.HGet2(str, name, "123456")
		fmt.Println(vmap, err)
	}
	// str := "http://localhost:9090/user/reg?account=%v&password=%v"
	// vmap, err := util.HGet2(str, "c1", "123456")

}

func TestClogin(t *testing.T) {
	for i := 0; i < 10; i++ {
		str := "http://localhost:9090/user/login?account=%v&password=%v"
		name := "888888"
		if i%2 == 0 {
			name = "c1"
		}
		fmt.Println(name)
		vmap, err := util.HGet2(str, name, "123456")
		fmt.Println(vmap, err)
	}
}

func TestCshowuser(t *testing.T) {
	for i := 0; i < 10; i++ {
		str := "http://localhost:9090/account/showuser?page=%v&pageCount=%v"
		vmap, err := util.HGet2(str, 1, i+1)
		fmt.Println(vmap, err)
	}
}

func TestCdeleteUser(t *testing.T) {
	for i := 0; i < 10; i++ {
		str := "http://localhost:9090/user/delete?userid=%v"
		vmap, err := util.HGet2(str, 1)
		fmt.Println(vmap, err)
	}
}
