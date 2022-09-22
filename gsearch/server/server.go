package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jimuyang/lets-go/gsearch/google"
	"github.com/jimuyang/lets-go/gsearch/userip"
	"log"
	"net/http"
	"time"
)

// google搜索的实际栗子 /search?q=golang&timeout=1s
func handleSearch(w http.ResponseWriter, req *http.Request) {
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)
	timeout, err := time.ParseDuration(req.FormValue("timeout"))
	if err == nil {
		// 请求有timeout
		ctx, cancel = context.WithTimeout(context.Background(), timeout)
	} else {
		ctx, cancel = context.WithCancel(context.Background())
	}
	defer cancel()

	query := req.FormValue("q")
	if query == "" {
		http.Error(w, "no query", http.StatusBadRequest)
		return
	}
	fmt.Println(query)
	userIP, err := userip.FromRequest(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(userIP)
	ctx = userip.NewContext(ctx, userIP)

	// Run the Google search and print the results.
	//start := time.Now()
	results, err := google.Search(ctx, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//elapsed := time.Since(start)

	bytes, err := json.Marshal(results)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(bytes)
}

func main() {
	http.HandleFunc("/search", handleSearch)
	log.Println("static web server start!")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
