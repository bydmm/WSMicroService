package main

// WS请求
type WSRequest struct {
	RequestID string `json:"rid"`
	Action    string `json:"action"`
}

// WS响应
type WSResponse struct {
	RequestID string      `json:"rid"`
	Data      interface{} `json:"data"`
}
