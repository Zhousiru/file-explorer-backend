package api

type Resp struct {
	Payload interface{} `json:"payload"`
	Err     string      `json:"err"`
}
