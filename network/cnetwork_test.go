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
	str := "http://localhost:9090/user/reg?account=%v&password=%v"
	for i := 0; i < 20; i++ {
		if i > 10 {
			vmap, err := util.HGet2(str, "c1", "123456")
			fmt.Println(vmap, err)
		} else {

			name := "c" + strconv.Itoa(ti.Nanosecond()) + strconv.Itoa(i)
			vmap, err := util.HGet2(str, name, "123456")
			fmt.Println(vmap, err)
		}
	}
	vmap, err := util.HGet2(str, "aaaaaaaaaaaaaaaaaaa22222222222221111111111113333333333333355555555555555", "123456")
	fmt.Println(vmap, err)
	// str := "http://localhost:9090/user/reg?account=%v&password=%v"
	// vmap, err := util.HGet2(str, "c1", "123456")

}

func TestClogin(t *testing.T) {
	str := "http://localhost:9090/user/login?account=%v&password=%v"
	for i := 0; i < 10; i++ {
		name := "888888"
		if i%2 == 0 {
			name = "c1"
		}
		vmap, err := util.HGet2(str, name, "123456")
		fmt.Println(vmap, err)
	}
	vmap, err := util.HGet2(str, "aaaaaaaaaaaaaaaaaaa22222222222221111111111113333333333333355555555555555", "123456")
	fmt.Println(vmap, err)
}

func TestCshowuser(t *testing.T) {
	for i := 0; i < 10; i++ {
		str := "http://localhost:9090/account/showuser?page=%v&pageCount=%v"
		vmap, err := util.HGet2(str, 1, i+1)
		fmt.Println(vmap, err)
	}
	str := "http://localhost:9090/account/showuser?page=%v&pageCount=%v"
	vmap, err := util.HGet2(str, "a", 10)
	fmt.Println(vmap, err)
}

func TestCdeleteUser(t *testing.T) {
	str := "http://localhost:9090/user/delete?userid=%v"
	for i := 0; i < 10; i++ {
		vmap, err := util.HGet2(str, 1)
		fmt.Println(vmap, err)
	}
	vmap, err := util.HGet2(str, "a")
	fmt.Println(vmap, err)
}

func TestCfillinfo(t *testing.T) {
	str := "http://localhost:9090//user/fill?userid=%v&username=%v&age=%v&gender=%v"
	for i := 0; i < 10; i++ {
		vmap, err := util.HGet2(str, 1, "aaaa", 10, 1)
		fmt.Println(vmap, err)
	}
	vmap, err := util.HGet2(str, 13, "aaaa", 10, 1)
	fmt.Println(vmap, err)
	vmap, err = util.HGet2(str, 13, "aaaa", "c", 1)
	fmt.Println(vmap, err)
}

func TestCclose(t *testing.T) {
	CClose()
}
