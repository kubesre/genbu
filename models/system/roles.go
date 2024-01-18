/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2023/12/4
*/

package system

import "gorm.io/gorm"

// 用户角色

type Role struct {
	gorm.Model
	Name   string `gorm:"column:name;comment:'角色名称';size:128" json:"name"`
	Desc   string `gorm:"column:desc;comment:'角色描述';size:128" json:"desc"`
	Status uint   `gorm:"type:tinyint(1);default:1;comment:'用户状态(正常/禁用, 默认正常)'" json:"status"`
	Menus  []Menu `gorm:"many2many:relation_role_menu" json:"menus"`
	Users  []User `gorm:"foreignkey:RoleId"`
}

func (*Role) TableName() string {
	return "role"
}
