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
	"strings"
)

// 授权相关

type InterfaceCasbin interface {
	AddPolicy(policy []*mod.CasbinPolicy) error
	DelPolicy(policy []*mod.CasbinPolicy) error
	ListPolicy(keyWord string, limit, page int) *mod.CasbinPolicyList
}

type casbinInfo struct{}

func NewCasbinInterface() InterfaceCasbin {
	return &casbinInfo{}
}

// 添加授权

func (c *casbinInfo) AddPolicy(policy []*mod.CasbinPolicy) error {
	if len(policy) > 0 {
		for _, item := range policy {
			if ok, _ := global.CasbinEnforcer.AddPolicy(item.RoleID, item.Path, item.Method, item.Desc); !ok {
				global.TPLogger.Error("权限已经存在")
				return errors.New("权限已经存在")
			}
			global.TPLogger.Info("权限添加成功")
		}
		return nil
	}
	global.TPLogger.Error("权限不能为空")
	return errors.New("权限不能为空")
}

// 删除授权

func (c *casbinInfo) DelPolicy(policy []*mod.CasbinPolicy) error {
	if len(policy) > 0 {
		for _, item := range policy {
			if _, err := global.CasbinEnforcer.RemovePolicy(item.RoleID, item.Path, item.Method, item.Desc); err != nil {
				global.TPLogger.Error("接口删除失败")
				return errors.New("接口删除失败")
			}
			return nil
		}
	}
	global.TPLogger.Error("权限不能为空")
	return errors.New("权限不能为空")
}

// 查看授权

func (c *casbinInfo) ListPolicy(keyWord string, limit, page int) *mod.CasbinPolicyList {
	var (
		policy mod.CasbinPolicy
		total  int
	)
	policys := make([]mod.CasbinPolicy, 0)
	casbinData := global.CasbinEnforcer.GetPolicy()
	// 组装一下数据
	for _, item := range casbinData {
		if policyFilter(item, keyWord) {
			policy.RoleID = item[0]
			policy.Path = item[1]
			policy.Method = item[2]
			policy.Desc = item[3]
			policys = append(policys, policy)
		}

	}
	total = len(policys)
	// 自定义处理分页
	if limit <= 0 || page <= 0 {
		return &mod.CasbinPolicyList{
			Items: policys,
			Total: total,
		}
	}
	/*
		举例1：
		limit 2  page 1  也就是 一页2条数据
		startIndex 0
		endIndex 1
		policys[0:1]
		举例2：
		limit 1  page  1 也即是 一页一条数据
		startIndex 0
		endIndex 1
		policys[0:1]
	*/
	startIndex := limit * (page - 1)
	endIndex := limit * page
	if endIndex > total {
		endIndex = total
	}
	policys = policys[startIndex:endIndex]
	return &mod.CasbinPolicyList{
		Items: policys,
		Total: total,
	}
}

// 筛选

func policyFilter(source []string, keyWord string) bool {
	for _, item := range source {
		if strings.Contains(item, keyWord) {
			return true
		}
	}
	return false
}
