package apis

type errResp struct {
	Error string `json:"error,omitempty"`
}

func newErrResp(err error) *errResp {
	return &errResp{Error: err.Error()}
}

type msgReq struct {
	ProjectID      uint64 `json:"projectID"        binding:"required"`
	ProjectVersion string `json:"projectVersion"   binding:"required"`
	Data           string `json:"data"             binding:"required"`
}
