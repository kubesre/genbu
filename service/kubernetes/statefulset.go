package kubernetes

/*
@auth: Meersburg
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/2/02
*/
import (
	"context"
	"errors"
	"genbu/common/global"
	"genbu/utils"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/kubernetes"
	"time"
)

//使用yaml创建statefulSet

func (k *k8sCluster) CreateStatefulSetYaml(cid, content string) (err error) {

	//获取clientSet所有信息
	clientSetAny, found := global.ClientSetCache.Get(cid)
	if !found {
		global.TPLogger.Error("当前集群不存在：", err)
		return errors.New("当前集群不存在")
	}

	contentsYaml, err := utils.DecodeBase64(content)
	if err != nil {
		global.TPLogger.Error("base64解码失败", err)
		return errors.New("base64解码失败")
	}
	var stateful = &appsv1.StatefulSet{}

	err = yaml.UnmarshalStrict([]byte(contentsYaml), &stateful)
	if err != nil {
		global.TPLogger.Error("解析文件内容失败", err)
		return errors.New("解析文件内容失败")
	}

	if stateful.Namespace == "" {
		stateful.Namespace = "default"
	}
	//获取clientSet
	clientSet := clientSetAny.(*kubernetes.Clientset)
	// 获取config
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)
	ctx, cancel = context.WithTimeout(context.TODO(), time.Second*2)
	defer cancel()
	_, err = clientSet.AppsV1().StatefulSets(stateful.Namespace).Create(ctx, stateful, metav1.CreateOptions{})
	if err != nil {
		global.TPLogger.Error("创建statefulSet失败", err)
		return errors.New("创建statefulSet失败")
	}
	return nil
}

//使用yaml更新statefulSet

func (k *k8sCluster) UpdateStatefulSetYaml(cid, content string) (err error) {
	//获取clientSet所有信息
	clientSetAny, found := global.ClientSetCache.Get(cid)
	if !found {
		global.TPLogger.Error("当前集群不存在：", err)
		return errors.New("当前集群不存在")
	}

	contentsYaml, err := utils.DecodeBase64(content)
	if err != nil {
		global.TPLogger.Error("base64解码失败", err)
		return errors.New("base64解码失败")
	}
	var stateful = &appsv1.StatefulSet{}

	err = yaml.UnmarshalStrict([]byte(contentsYaml), &stateful)
	if err != nil {
		global.TPLogger.Error("解析文件内容失败", err)
		return errors.New("解析文件内容失败")
	}

	if stateful.Namespace == "" {
		stateful.Namespace = "default"
	}
	//获取clientSet
	clientSet := clientSetAny.(*kubernetes.Clientset)
	// 获取config
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)
	ctx, cancel = context.WithTimeout(context.TODO(), time.Second*2)
	defer cancel()
	_, err = clientSet.AppsV1().StatefulSets(stateful.Namespace).Update(ctx, stateful, metav1.UpdateOptions{})
	if err != nil {
		global.TPLogger.Error("更新statefulSet失败", err)
		return errors.New("更新statefulSet失败")
	}
	return nil
}

// 获取statefulset列表
func (k *k8sCluster) GetStatefulSetList(cid, namespace, filterName string, limit, page int) (statefulSetsResp *global.StatusfulSetsResp, err error) {
	//获取clientSet所有信息
	clientSetAny, found := global.ClientSetCache.Get(cid)
	if !found {
		global.TPLogger.Error("当前集群不存在：", err)
		return nil, errors.New("当前集群不存在")
	}
	//获取clientSet
	clientSet := clientSetAny.(*kubernetes.Clientset)
	// 获取config
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)
	ctx, cancel = context.WithTimeout(context.TODO(), time.Second*2)
	defer cancel()
	//获取statefulSet列表
	statefulSetList, err := clientSet.AppsV1().StatefulSets(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		global.TPLogger.Error("获取statefulSet列表失败：", err)
		return nil, errors.New("获取statefulSet列表失败")
	}
	var result []global.StatefulSetData
	for _, item := range statefulSetList.Items {
		result = append(result, global.StatefulSetData{
			Name:              item.Name,
			Namespace:         item.Namespace,
			CreationTimestamp: item.CreationTimestamp.Time,
			Replicas:          item.Status.Replicas,
			ReadyReplicas:     item.Status.ReadyReplicas,
			CurrentReplicas:   item.Status.CurrentReplicas,
		})
	}

	//将result []global.StatefulSetData，放进Dataselector对象中，进行排序
	selectableData := &DataSelector{
		GenericDataList: k.StatefulSetToCells(result),
		DataSelect: &DataSelectQuery{
			Filter: &FilterQuery{Name: filterName},
			Paginatite: &PaginateQuery{
				Limit: limit,
				Page:  page,
			},
		},
	}
	filtered := selectableData.Filter()
	total := len(filtered.GenericDataList)
	data := filtered.Sort().Paginate()
	//将[]DataCell类型的转换为global.StatefulSetData列表
	statefulSets := k.StatefulSetFromCells(data.GenericDataList)

	return &global.StatusfulSetsResp{
		Items: statefulSets,
		Total: total,
	}, nil
}

// 获取statefulSet详情
func (k *k8sCluster) GetStatefulSetDetail(cid, name, namespace string) (statefulSet *appsv1.StatefulSet, err error) {
	//获取clientSet所有信息
	clientSetAny, found := global.ClientSetCache.Get(cid)
	if !found {
		global.TPLogger.Error("当前集群不存在：", err)
		return nil, errors.New("当前集群不存在")
	}
	//获取clientSet
	clientSet := clientSetAny.(*kubernetes.Clientset)
	// 获取config
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)
	ctx, cancel = context.WithTimeout(context.TODO(), time.Second*2)
	defer cancel()
	//获取statefulSet详情
	statefulSet, err = clientSet.AppsV1().StatefulSets(namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		global.TPLogger.Error("获取statefulSet详情失败：", err)
		return nil, errors.New("获取statefulSet详情失败")
	}
	//添加元数据信息，默认获取不到
	statefulSet.APIVersion = "apps/v1"
	statefulSet.Kind = "StatefulSet"
	return statefulSet, err
}

//删除statefulSet

func (k *k8sCluster) DeleteStatefulSet(cid, name, namespace string) (err error) {
	//获取clientSet所有信息
	clientSetAny, found := global.ClientSetCache.Get(cid)
	if !found {
		global.TPLogger.Error("当前集群不存在：", err)
		return errors.New("当前集群不存在")
	}
	//获取clientSet
	clientSet := clientSetAny.(*kubernetes.Clientset)
	// 获取config
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)
	ctx, cancel = context.WithTimeout(context.TODO(), time.Second*2)
	defer cancel()
	//获取statefulSet详情
	err = clientSet.AppsV1().StatefulSets(namespace).Delete(ctx, name, metav1.DeleteOptions{})
	if err != nil {
		global.TPLogger.Error("获取statefulSet详情失败：", err)
		return errors.New("获取statefulSet详情失败")
	}
	return nil
}

func (k *k8sCluster) StatefulSetToCells(std []global.StatefulSetData) []DataCell {
	cells := make([]DataCell, len(std))
	for i := range std {
		cells[i] = statefulSetCell(std[i])
	}
	return cells
}

func (k *k8sCluster) StatefulSetFromCells(cells []DataCell) []global.StatefulSetData {
	statefulSets := make([]global.StatefulSetData, len(cells))
	for i := range cells {
		statefulSets[i] = global.StatefulSetData(cells[i].(statefulSetCell))
	}

	return statefulSets
}
