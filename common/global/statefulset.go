package global

/*
@auth: Meersburg
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/2/02
*/
import "time"

type StatefulSetData struct {
	Name              string    `json:"name"`
	Namespace         string    `json:"namespace"`
	CreationTimestamp time.Time `json:"creationTimestamp"`
	Replicas          int32     `json:"replicas"`
	ReadyReplicas     int32     `json:"readyReplicas"`
	CurrentReplicas   int32     `json:"currentReplicas"`
}

type StatusfulSetsResp struct {
	Items []StatefulSetData `json:"items"`
	Total int               `json:"total"`
}
