package main

import (
	"log"
	"net/http"
	"net/http/httputil"
)

func main11() {
	http.HandleFunc("/dingtalk/", func(writer http.ResponseWriter, req *http.Request) {
		data, _ := httputil.DumpRequest(req, true)
		log.Printf("%s\n", data)
		writer.Write([]byte("hello world"))
	})

	log.Println("static web server start!")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
