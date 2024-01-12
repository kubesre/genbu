/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2023/12/4
*/

package users

import (
	"errors"
	"fmt"
	"genbu/common/global"
	"genbu/dao"
	"genbu/models"
	"strconv"
)

type InterfaceUsers interface {
	Register(user *models.User) error
	UserInfo(id interface{}) (*models.User, error)
	UserList(username string, limit, page int) (*models.UserList, error)
	UserUpdate(id uint, userData *models.User) error
	UserAdd(user *models.User) error
}
type userInfo struct{}

func NewUserInfo() InterfaceUsers {
	return &userInfo{}
}

// 用户注册

func (u *userInfo) Register(user *models.User) error {
	err := dao.NewUserInterface().Register(user)
	if err != nil {
		global.TPLogger.Error("用户注册失败：", err)
		return errors.New("用户注册失败")
	}
	return err
}

// 用户详情

func (u *userInfo) UserInfo(id interface{}) (*models.User, error) {
	idUint := fmt.Sprintf("%d", id)
	idInt, err := strconv.Atoi(idUint)
	if err != nil {
		global.TPLogger.Error("用户详情查看失败：")
		return nil, errors.New("用户详情查看失败")
	}
	data, err := dao.NewUserInterface().UserInfo(idInt)
	if err != nil {
		global.TPLogger.Error("用户详情查看失败：", err)
		return nil, errors.New("用户详情查看失败")
	}
	return data, nil
}

// 用户列表

func (u *userInfo) UserList(username string, limit, page int) (*models.UserList, error) {
	data, err := dao.NewUserInterface().UserList(username, limit, page)
	if err != nil {
		global.TPLogger.Error("UserList失败：", err)
		return nil, errors.New("UserList失败")
	}
	return data, nil
}

// 用户更新

func (u *userInfo) UserUpdate(id uint, userData *models.User) error {
	err := dao.NewUserInterface().UserUpdate(int(id), userData)
	if err != nil {
		global.TPLogger.Error("用户更新失败：", err)
		return errors.New("用户更新失败")
	}
	return nil
}

// 用户添加

func (u *userInfo) UserAdd(user *models.User) error {
	err := dao.NewUserInterface().UserAdd(user)
	if err != nil {
		global.TPLogger.Error("用户添加失败：", err)
		return errors.New("用户添加失败")
	}
	return nil
}
