package main

import (
	"fmt"
	"time"

	"github.com/guopuke/learngo/retriever/mock"
	"github.com/guopuke/learngo/retriever/real"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

type RetrieverPoster interface {
	Retriever
	Poster
	// Other...
}

func download(r Retriever) string {
	return r.Get("http://baidu.com")
}

func post(poster Poster) {
	poster.Post("http://baidu.com",
		map[string]string{
			"name":   "Hello",
			"course": "golang",
		})
}

const url = "http://baidu.com"

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked baidu.com",
	})

	return s.Get(url)
}

func main() {
	var r Retriever

	retriever := mock.Retriever{
		"Hello retriever, this is a fake baidu.com",
	}
	r = &retriever
	// mock.Retriever {Hello retriever}
	// fmt.Printf("%T %v\n", r, r)
	inspect(r)

	r = &real.Retriever{
		UserAgent: "Chrome/9.9",
		TimeOut:   time.Minute,
	}
	// *real.Retriever &{Chrome/9.9 1m0s}
	// fmt.Printf("%T %v\n", r, r)
	inspect(r)

	// Type Assertion
	realRetriever := r.(*real.Retriever)
	fmt.Println(realRetriever.TimeOut)

	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}

	// fmt.Println(download(r))

	fmt.Println("Try a session")
	fmt.Println(session(&retriever))
}

// Type Switch
func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}
