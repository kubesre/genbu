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

// 部门相关

type InterfaceDept interface {
	AddDept(dept *mod.Dept) error
	DeptList() ([]mod.Dept, error)
	DeptInfo(did int) ([]mod.Dept, error)
	DelDept(did int) error
}

type deptInfo struct{}

func NewDeptInterface() InterfaceDept {
	return &deptInfo{}
}

// 创建部门

func (d *deptInfo) AddDept(dept *mod.Dept) error {
	err := global.GORM.Model(&mod.Dept{}).Create(&dept).Error
	return err
}

// 部门列表

func (d *deptInfo) DeptList() ([]mod.Dept, error) {
	var depts []mod.Dept
	if err := global.GORM.Where("parent_id", 0).Preload("Users").Preload("Users.Role").
		Preload("Users.Dept").Find(&depts).Error; err != nil {
		return nil, err
	}
	for i := range depts {
		err := childrenDept(&depts[i])
		if err != nil {
			return nil, err
		}
	}
	return depts, nil
}

// 部门详情

func (d *deptInfo) DeptInfo(did int) ([]mod.Dept, error) {
	var depts []mod.Dept
	if err := global.GORM.Where("id = ?", did).Preload("Users").First(&depts).Error; err != nil {
		return nil, err
	}
	for i := range depts {
		err := childrenDept(&depts[i])
		if err != nil {
			return nil, err
		}
	}
	return depts, nil
}

// 删除部门 如果删除了父部门，其子部门也会被删除哦
var childrenID []int

func (d *deptInfo) DelDept(did int) error {
	childrenID = nil
	var depts mod.Dept
	if err := global.GORM.Where("id = ?", did).First(&depts).Error; err != nil {
		return err
	}
	ids, err := delchildrenDept(&depts)
	if err != nil {
		return err
	}
	if err = global.GORM.Model(&mod.Dept{}).Where("id IN ?", ids).Delete(&mod.Dept{}).Error; err != nil {
		return err
	}
	return nil
}

// 定义一个函数用户查询部门以及子部门

func childrenDept(dept *mod.Dept) error {
	var children []*mod.Dept
	if err := global.GORM.Where("parent_id = ?", dept.ID).Preload("Users").Preload("Users.Role").
		Preload("Users.Dept").Find(&children).Error; err != nil {
		return err
	}
	for i := range children {
		if err := childrenDept(children[i]); err != nil {
			return err
		}
	}
	dept.Children = children
	return nil
}

// 定义一个函数用户删除部门以及子部门

func delchildrenDept(dept *mod.Dept) ([]int, error) {
	/*
		ID 2
		dept.Children  ID  3 4
			循环  ID 3 parentID 为3 的有 id 为5
	*/
	if err := global.GORM.Where("parent_id", dept.ID).Find(&dept.Children).Error; err != nil {
		return nil, err
	}
	childrenID = append(childrenID, int(dept.ID))
	for i := range dept.Children {
		_, err := delchildrenDept(dept.Children[i])
		if err != nil {
			return nil, err
		}
	}
	return childrenID, nil
}
