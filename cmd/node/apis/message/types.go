package message

type HandleErrRsp struct {
	Error string `json:"error,omitempty"`
}

type HandleReq struct {
	ProjectID      string `json:"projectID"        binding:"required"`
	ProjectVersion string `json:"projectVersion"   binding:"required"`
	Data           string `json:"data"             binding:"required"`
}
