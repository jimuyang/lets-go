package fetcher

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

const UserAgent string = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36"

// Fetch 获取结果
func Fetch(url string) ([]byte, error) {
	log.Printf("fetching: %s", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("user-agent", UserAgent)
	resp, err := http.DefaultClient.Do(req)
	//resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bodyReader := bufio.NewReader(resp.Body)
	ec := determineEncoding(bodyReader)
	// fmt.Println("encoding:", encoding)
	if resp.StatusCode == http.StatusOK {
		// 转化为utf8
		utf8Reader := transform.NewReader(bodyReader, ec.NewDecoder())
		return ioutil.ReadAll(utf8Reader)
	}
	return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	ec, _, _ := charset.DetermineEncoding(bytes, "")
	return ec
}
