package service

import (
	"strconv"
	"tantan-simplify/model"
	"tantan-simplify/pkg/constant"
	"tantan-simplify/types"
)

func CreateUser(name string) (userResult types.User, err error) {
	user, err := model.AddUser(name)
	if err != nil {
		return userResult, err
	}
	userResult.Id = strconv.Itoa(user.Id)
	userResult.Name = user.Name
	userResult.Type = constant.UserStr
	return userResult, nil
}

func ListAllUsers() (userList []types.User, err error) {
	users, err := model.ListAllUsers()
	if err != nil {
		return nil, err
	}
	userList = make([]types.User, len(users))
	for k, user := range users {
		userList[k].Id = strconv.Itoa(user.Id)
		userList[k].Name = user.Name
		userList[k].Type = constant.UserStr
	}
	return userList, nil
}
