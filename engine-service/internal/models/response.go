package models

type Response struct {
	Action string      `json:"action"`
	Data   interface{} `json:"data"`
	Err    string      `json:"err"`
	ConnID string      `json:"conn_id"`
}
