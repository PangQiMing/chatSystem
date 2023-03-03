package helper

import "strings"

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
	Data    interface{} `json:"data"`
}

type EmptyObj struct{}

// BuildResponse 建立一个响应成功的方法
func BuildResponse(status bool, message string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Error:   nil,
		Data:    data,
	}
	return res
}

// BuildErrResponse 建立一个响应失败的方法
func BuildErrResponse(message string, err string, data interface{}) Response {
	//  Error\nError
	splitErr := strings.Split(err, "\n")
	res := Response{
		Status:  false,
		Message: message,
		Error:   splitErr,
		Data:    data,
	}
	return res
}
