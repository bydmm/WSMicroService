package main

import (
	"encoding/json"
	"fmt"
)

// 班级登记表单表单
type ClassServerData struct {
	Name string `json:"name"`
}

// 班级登记表单表单
type ClassServerRequest struct {
	WSRequest
	Data ClassServerData `json:"data"`
}

// 班级登记表单
func ClassServer(rBody []byte) WSResponse {
	var request ClassServerRequest
	err := json.Unmarshal(rBody, &request)
	if err != nil {
		fmt.Printf("ClassServer request json fail: %s", err.Error())
	}
	// 做你的业务去吧，比如写数据库....
	return WSResponse{
		Data: "收到",
	}
}
