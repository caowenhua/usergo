package bean

type User struct {
	UserName string `m2s:"username" json:"userName"`
	UserId   int    `m2s:"userid" json:"userId"`
	Age      int    `m2s:"age" json:"age"`
	Gender   int    `m2s:"gender" json:"gender"`
}
