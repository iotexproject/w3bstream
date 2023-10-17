package main

type errResp struct {
	Error string `json:"error,omitempty"`
}

func newErrResp(err error) *errResp {
	return &errResp{Error: err.Error()}
}

type msgReq struct {
	Data string `json:"data"        binding:"required"`
}
