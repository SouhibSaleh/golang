package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

type simpleResponse struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func (res simpleResponse) toString() string {
	return ``
}
func fetchUrl() {
	for i := 1; i < 30; i++ {
		time1 := time.Now()
		url := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", i)
		res, err := http.Get(url)
		fmt.Println(url)
		if err != nil {
			log.Fatal(err)
			return
		}
		defer res.Body.Close()
		_, errr := io.ReadAll(res.Body)
		if errr != nil {
			log.Fatal(errr)
			return
		}
		time2 := time.Now().Sub(time1)
		fmt.Println(time2.Seconds())
	}

	//var response simpleResponse
	//err = json.Unmarshal(result, &response)
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
	//fmt.Println(string(result))
}

func fetchFromGo(url string, ch chan string) {

	time1 := time.Now()
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer res.Body.Close()
	_, errr := io.ReadAll(res.Body)
	if errr != nil {
		log.Fatal(errr)
		return
	}
	time2 := time.Now().Sub(time1)

	ch <- strconv.FormatFloat(time2.Seconds(), 'f', 10, 64)
}
func goFetch() {
	ch := make(chan string)
	for i := 0; i < 30; i++ {
		url := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", i+1)
		fmt.Println(url)
		go fetchFromGo(url, ch)
	}
	for i := 0; i < 30; i++ {
		fmt.Println(<-ch)
	}
}
