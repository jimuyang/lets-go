package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// 用于http.HandleFunc
type httpHandler func(http.ResponseWriter, *http.Request)

// 业务逻辑
type appHandler func(http.ResponseWriter, *http.Request) error

// UserError 可向客户端输出的Error
type UserError interface {
	error
	Message() string
}

// 一个简单的string版UserError实现
type userError string

func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}

// appHandler => httpHandler
func errorWrapper(handler appHandler) httpHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic: %v", r)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		err := handler(w, r)
		if err != nil {
			log.Printf("Error occured handling request: %s", err.Error())
			code := http.StatusInternalServerError
			if userError, ok := err.(UserError); ok {
				http.Error(w, userError.Message(), http.StatusBadRequest)
				return
			}
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			}
			http.Error(w, http.StatusText(code), code)
		}
	}
}

// 静态资源文件的读取和输出
func fileHandler(writer http.ResponseWriter, request *http.Request) error {
	path := request.URL.Path[len("/file/"):]

	if path == "error" {
		return userError("you want a error?")
	}

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}
	if fileInfo.IsDir() {
		fileInfos, err := ioutil.ReadDir(path)
		if err != nil {
			return err
		}
		for _, fileInfo := range fileInfos {
			writer.Write([]byte(fileInfo.Name() + "\n"))
		}
	} else {
		all, err := ioutil.ReadAll(file)
		if err != nil {
			return err
		}
		writer.Write(all)
	}
	return nil
}

func main1() {
	http.HandleFunc("/file/", errorWrapper(fileHandler))

	log.Println("static web server start!")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func newFunc() {
}
func newFunc1() {
}
