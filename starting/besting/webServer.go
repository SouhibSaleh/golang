package besting

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type MyHandler struct{}

var mux sync.Mutex
var count int

func (h MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	mux.Lock()
	count++
	mux.Unlock()
	obj := SimpleResponse{
		Id:     0,
		UserId: count,
		Body:   "Hello World!",
		Title:  "Hello World!",
	}

	jsonData, err := json.Marshal(obj)
	if err != nil {
		fmt.Println(err)
		return
	}
	formErr := r.ParseForm()
	if formErr != nil {
		fmt.Println(formErr)
		return
	}
	x := r.Form.Get("x")
	y := r.Form.Get("y")
	fmt.Println(x, y)

	w.Header().Set("Content-Type", "application/json")
	w.Write((jsonData))

}

func serverListner() {
	http.Handle("/test", MyHandler{})

	//http.HandleFunc("/", funcHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))

}

func funcHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
}
