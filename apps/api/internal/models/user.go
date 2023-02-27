package models

import "chat/apps/api/global"

type User struct {
	UserName string
	Password string
}

func GetUserList() []*User {
	UserList := make([]*User, 0)
	global.DB.Where("id > ?", 0).Find(&UserList)
	return UserList
}
