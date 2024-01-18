package domain

import corev1 "k8s.io/api/core/v1"

// NodesResp 领域对象，为了能在接口中透传
type NodesResp struct {
	Items []corev1.Node `json:"items"`
	Total int           `json:"total"`
}
