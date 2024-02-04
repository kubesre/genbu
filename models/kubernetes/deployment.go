/*
@auth: Ansu
@source: Ansu
@time: 2024/1/18
*/

package kubernetes

import "time"

type Time struct {
	time.Time `protobuf:"-"`
}
type Deployment struct {
	Name              string `json:"name" protobuf:"bytes,1,opt,name=name"`
	Namespace         string `json:"namespace" protobuf:"bytes,3,opt,name=namespace"`
	MinReadySeconds   int32  `json:"minReadySeconds" protobuf:"varint,5,opt,name=minReadySeconds"`
	Replicas          int32  `json:"replicas" protobuf:"varint,2,opt,name=replicas"`
	ReadyReplicas     int32  `json:"readyReplicas" protobuf:"varint,7,opt,name=readyReplicas"`
	CreationTimestamp Time   `json:"creationTimestamp" protobuf:"bytes,8,opt,name=creationTimestamp"`
}

type DeleteDeployment struct {
	Name      string `json:"name" form:"name" binding:"required"`
	NameSpace string `json:"namespace" form:"namespace"`
}

type UpdateDeployment2Yaml struct {
	Text string `json:"text" binding:"required"`
}

type UpdateDeployment2Arg struct {
	Name           string `json:"name" form:"name" binding:"required"`
	NameSpace      string `json:"namespace" form:"namespace"`
	ContainersName string `json:"containersName" form:"containersName"`
	Image          string `json:"image" form:"image"`
	Replicas       int32  `json:"replicas" form:"replicas"`
	LableKey       string `json:"lableKey" form:"lableKey"`
	LableValue     string `json:"lableValue" form:"lableValue"`
}
