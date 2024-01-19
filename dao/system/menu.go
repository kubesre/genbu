/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2023/12/4
*/

package system

import (
	"genbu/common/global"
	mod "genbu/models/system"
)

// 菜单相关

type InterfaceMenus interface {
	AddMens(menu *mod.Menu) error
	MenusList() ([]mod.Menu, error)
}

type menusInfo struct{}

func NewMenusInterface() InterfaceMenus {
	return &menusInfo{}
}

// 创建菜单

func (m *menusInfo) AddMens(menu *mod.Menu) error {
	err := global.GORM.Model(&mod.Menu{}).Create(&menu).Error
	return err
}

// 菜单列表

func (m *menusInfo) MenusList() ([]mod.Menu, error) {
	var menus []mod.Menu
	if err := global.GORM.Where("parent_id = ?", 0).Find(&menus).Error; err != nil {
		return nil, err
	}
	for i := range menus {
		err := childrenMenu(&menus[i])
		if err != nil {
			return nil, err
		}
	}
	return menus, nil
}

// 定义一个函数用于查询菜单及其子菜单

func childrenMenu(menu *mod.Menu) error {
	if err := global.GORM.Where("parent_id = ?", menu.ID).Find(&menu.Children).Error; err != nil {
		return err
	}
	for i := range menu.Children {
		err := childrenMenu(&menu.Children[i])
		if err != nil {
			return err
		}
	}
	return nil
}
