package model

type HelloReq struct {
	Name string `json:"name,omitempty"`
}

type ErrRsp struct {
	Err string `json:"error,omitempty"`
}

type HelloRsp struct {
	Msg string `json:"message,omitempty"`
}
