package kubernetes

type Secret struct {
	Name      string `json:"name" form:"name" binding:"required"`
	NameSpace string `json:"namespace" form:"namespace"`
}

type Create struct {
	SecretName string `json:"secret_name" binding:"required"` // 配置文件名称，kubectl get secret 查看服务显示时的名称
	NameSpace  string `json:"namespace"`
	Text       string `json:"text" binding:"required"`
	Type       string `json:"type"`
}
