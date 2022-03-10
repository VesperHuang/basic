package main

import (
	"basic/retriever/mock"
	"basic/retriever/real"
	"fmt"
)

const url = "http://www.imoc.com"

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string,
		form map[string]string) string
}

func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(url,
		map[string]string{
			"name":   "ccmouse",
			"course": "golang",
		})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked imooc.com",
	})
	return s.Get(url)

}

func inspect(r Retriever) {

	fmt.Println("Inspectiong", r)
	fmt.Printf("%T %v\n", r, r)
	fmt.Print("Type switch:")

	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
	fmt.Println()
}

func main() {
	var r Retriever
	retriver := mock.Retriever{
		"this is a fake imooc.com"}
	r = &retriver
	inspect(r)

	// r = &real.Retriever{
	// 	UserAgent: "Mozilla/5.0",
	// 	TimeOut:   time.Minute,
	// }
	// inspect(r)

	// if mockRetriever, ok := r.(*mock.Retriever); ok {
	// 	fmt.Println(mockRetriever.Contents)
	// } else {
	// 	fmt.Println("not a mock retriever")
	// }

	// realRetriever := r.(*real.Retriever)
	// fmt.Println(realRetriever.TimeOut)

	fmt.Println("Try a session")
	fmt.Println(session(&retriver))

	//fmt.Println(download(r))
}
