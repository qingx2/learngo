package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	// pcRequest()
	mobRequest()
}

func mobRequest() {
	request, err := http.NewRequest(http.MethodGet, "http://baidu.com", nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.139 Mobile Safari/537.36")

	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("Redirect:", req)
			return nil
		},
	}

	resp, e := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bytes, e := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(e)
	}
	fmt.Printf("%s\n", bytes)
}

func pcRequest() {
	resp, err := http.Get("http://baidu.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bytes, e := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(e)
	}
	fmt.Printf("%s\n", bytes)
}
