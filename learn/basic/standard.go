package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func usehttp() {
	req, err := http.NewRequest(http.MethodGet, "http://www.imooc.com", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("redirect:", req)
			return nil
		},
	}
	// resp, err := http.DefaultClient.Do(req)
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", data)
}

// func main() {
// 	usehttp()
// }
