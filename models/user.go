package models

import "fmt"

func GetUserInfo(username, password string) *User {
	var user User
	var param QueryParams

	param.PageSize = 1
	param.Keyword = fmt.Sprintf("username = \"%s\"", username)
	err := QueryData(&user, param)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	//if user.Username == username && user.Password != user.Password {
	//	return nil
	//}
	return &user

}

func CreateAdminUser() {
	var user User
	user.Username = "admin"
	user.Password = "123456"

	keyword := fmt.Sprintf("username = \"%s\"", "admin")
	userRet := QueryUser(keyword, 0)

	if len(userRet) == 0 {
		err := InsertData(&user)
		if err != nil {
			return
		}
	}
}

func QueryUser(keyword string, size int) []User {
	var user []User
	var param QueryParams

	if param.PageSize != 0 {
		param.PageSize = size
	}
	if param.Keyword != "" {
		param.Keyword = keyword
	}

	err := QueryData(&user, param)
	if err != nil {
		return nil
	}
	return user
}
