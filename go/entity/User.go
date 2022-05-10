package entity

func (User) TableName() string {
	return "sys_user"
}

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
	NickName string `json:"nickName"`
	Avatar string `json:"avatar"`
	Password string `json:"password"`
	Email string `json:"email"`
	Mobile string `json:"mobile"`
	DelStatus int `json:"delStatus"`
	CreateTime int64 `json:"createTime"`
}

