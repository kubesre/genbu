/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2023/12/4
*/

package system

import (
	"errors"
	"genbu/common/global"
	mod "genbu/models/system"
)

// 角色相关

type InterfaceRoles interface {
	AddRole(role *mod.Role) error
	RoleInfo(rid uint) (*mod.Role, error)
	UpdateRole(rid uint, roleData *mod.Role) error
	AddRelationRoleAndMenu(menuID []int, roleID int) error
	DelRole(rid []int) error
	RoleList() (roleData []*mod.Role, err error)
}

type rolesInfo struct{}

func NewRolesInterface() InterfaceRoles {
	return &rolesInfo{}
}

// 创建角色

func (r *rolesInfo) AddRole(role *mod.Role) error {
	err := global.GORM.Create(&role).Error
	return err
}

// 获取角色详情

func (r *rolesInfo) RoleInfo(rid uint) (*mod.Role, error) {
	var roleData *mod.Role
	err := global.GORM.Model(&mod.Role{}).Where("id = ?", rid).Preload("Menus").First(&roleData).Error
	return roleData, err
}

// 更新角色信息

func (r *rolesInfo) UpdateRole(rid uint, roleData *mod.Role) error {
	err := global.GORM.Model(&mod.Role{}).Where("id = ?", rid).Updates(&roleData).Error
	return err
}

// 创建角色对应的菜单

func (r *rolesInfo) AddRelationRoleAndMenu(menuID []int, roleID int) error {
	var (
		menuList []mod.Menu
		role     mod.Role
	)

	// 查询菜单列表
	global.GORM.Find(&menuList, menuID)
	if len(menuList) != len(menuID) {
		return errors.New("菜单不存在")
	}

	// 查询角色列表
	if err := global.GORM.First(&role, roleID).Error; err != nil {
		return errors.New("角色不存在或查询角色失败")
	}

	err := global.GORM.Model(&role).Association("Menus").Append(&menuList)
	if err != nil {
		return err
	}
	return nil
}

// 删除角色

func (r *rolesInfo) DelRole(rid []int) error {
	var (
		roleData []mod.Role
	)
	global.GORM.Find(&roleData, rid)
	if len(roleData) != len(rid) {
		return errors.New("角色列表中有不存在的ID")
	}
	// 清空角色与菜单的关系
	if err := global.GORM.Model(&roleData).Association("Menus").Clear(); err != nil {
		return errors.New("清空角色与菜单的关系失败:" + err.Error())
	}
	// 删除角色
	if err := global.GORM.Delete(&roleData, rid).Error; err != nil {
		return err
	}
	return nil
}

// 角色列表

func (r *rolesInfo) RoleList() (roleData []*mod.Role, err error) {
	if err = global.GORM.Model(&mod.Role{}).Find(&roleData).Error; err != nil {
		return nil, err
	}
	return roleData, nil
}
