/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/1/9
*/

package kubernetes

import (
	corev1 "k8s.io/api/core/v1"
	"sort"
	"strings"
	"time"
)

// 过滤排序

// dataSelector 用于封装排序，过滤，分页的数据类型

type DataSelector struct {
	GenericDataList []DataCell
	DataSelect      *DataSelectQuery
}

// DataCell 接口，用于各种资源List的类型转换，转换后可以使用dataSelector的排序，过滤，分页方法
type DataCell interface {
	GetCreation() time.Time // 根据是时间排序
	GetName() string        // 获取名称
}

//定义过滤和分页的结构体 过滤：Name 分页：limit和page limit是单页的数据条数  page是第几页

type DataSelectQuery struct {
	Filter     *FilterQuery
	Paginatite *PaginateQuery
}

type FilterQuery struct {
	Name string
}

type PaginateQuery struct {
	Limit int
	Page  int
}

// 实现自定义结构的排序，需要重写len swap less方法
// len方法 用于获取数组的长度

func (d *DataSelector) Len() int {

	return len(d.GenericDataList)
}

// swap方法用于数据比较大小之后的位置变更

func (d *DataSelector) Swap(i, j int) {
	d.GenericDataList[i], d.GenericDataList[j] = d.GenericDataList[j], d.GenericDataList[i]
}

// less方法用于比较大小  可以看到我们使用的是GetCreation()方法，其获取的是资源创建时间

func (d *DataSelector) Less(i, j int) bool {
	a := d.GenericDataList[i].GetCreation()
	b := d.GenericDataList[j].GetCreation()
	return b.Before(a)
}

// 重写以上三个方法，使用sort.sort进行排序

func (d *DataSelector) Sort() *DataSelector {
	sort.Sort(d)
	return d
}

// Filter 方法用于过滤数据，比较数据的Name属性，若包含，则返回
func (d *DataSelector) Filter() *DataSelector {
	// 判断入参是否为空，若为空，则返回所有数据
	if d.DataSelect.Filter.Name == "" {
		return d
	}
	// 若不为空，则按照入参Name进行过滤

	var filtered []DataCell
	for _, value := range d.GenericDataList {
		// 定义是否匹配的标签，默认是匹配的
		matches := true
		objName := value.GetName()
		if !strings.Contains(objName, d.DataSelect.Filter.Name) {
			matches = false
			continue
		}
		if matches {
			filtered = append(filtered, value)
		}

	}
	d.GenericDataList = filtered
	return d
}

// Pagination方法用于对于数组的分页，根据limit和page的传参，取一定范围内的数据 返回

func (d *DataSelector) Paginate() *DataSelector {
	// 根据Limit和Page的传参，定义快捷变量
	limit := d.DataSelect.Paginatite.Limit
	page := d.DataSelect.Paginatite.Page
	// 检验参数的合法性
	if limit <= 0 || page <= 0 {
		return d
	}
	// 定义取出方位需要的start index 和end index
	startIndex := limit * (page - 1)
	endIndex := limit * page

	// 处理endIndex
	if endIndex > len(d.GenericDataList) {
		endIndex = len(d.GenericDataList)
	}
	d.GenericDataList = d.GenericDataList[startIndex:endIndex]
	return d

}

// pod 资源格式化

type podCell corev1.Pod

// 重写DataCell接口的两个方法   pod

func (p podCell) GetCreation() time.Time {
	return p.CreationTimestamp.Time
}

func (p podCell) GetName() string {
	return p.Name
}
func ToCells(pods []corev1.Pod) []DataCell {

	cells := make([]DataCell, len(pods))
	for i := range pods {
		cells[i] = podCell(pods[i])
	}
	// 这里其实是将pod资源数据复制到DataCell切片中了
	return cells

}
func FromCells(cells []DataCell) []corev1.Pod {
	pods := make([]corev1.Pod, len(cells))
	for i := range cells {
		//  cells[i].(podCell) 将DataCell类型转成podCell 接口断言
		pods[i] = corev1.Pod(cells[i].(podCell))
	}
	return pods
}

// node
type nodeCell corev1.Node

func (p nodeCell) GetCreation() time.Time {
	return p.CreationTimestamp.Time
}

func (p nodeCell) GetName() string {
	return p.Name
}
