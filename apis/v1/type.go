package v1

type List struct {
	Count int64       `json:"count"`
	Items interface{} `json:"items"`
}

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name" binding:"omin=3,max=12" msg:"用户名称不合法"`
	Password string `json:"password" binding:"min=3,max=13" msg:"用户密码不合法"`
	Email    string `json:"email"`
}
