package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	//
	http.HandleFunc("/user/login", userLogin)
	// start web server
	fmt.Println("starting...")
	http.ListenAndServe(":8080", nil)
}

func userLogin(writer http.ResponseWriter, request *http.Request) {
	mobile := request.PostForm.Get("mobile")
	password := request.PostForm.Get("password")
	// fmt.Println("mobile:" + mobile)
	// fmt.Println("password:" + password)
	fmt.Println(request.Form)

	loginSuccess := false
	if mobile == password {
		loginSuccess = true
	}
	if loginSuccess {
		// data : {"id": 1, "token": "xxx"}
		data := make(map[string]interface{})
		data["id"] = 1
		data["token"] = "xxx"
		commonResponse(writer, data, "", nil)
	} else {
		commonResponse(writer, nil, "密码不正确", nil)
	}
}

// Response 定义统一返回值类型
type Response struct {
	Code    int
	Message string
	Data    interface{}
}

func commonResponse(writer http.ResponseWriter, data interface{}, message string, err error) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	var code int
	if err != nil || data == nil {
		code = -1
	} else {
		code = 0
	}
	response := Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
	ret, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err.Error())
	}
	writer.Write(ret)
}
