package test

import (
	"fmt"
	"net/http"
)

type dollars float32
type database map[string]dollars

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "%s", "we")
}
func (db database) home(w http.ResponseWriter, req *http.Request) {

	//_, err := fmt.Fprintf(w, "%s", "home")
	//w.Write([]byte("home"))
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//p := Point{0, 1}
	w.Header().Set("content-type", "text/html")
	//w.Header().Set("Content-Type", "application/json")
	//byts, errM := json.Marshal(p)
	//if errM != nil {
	//	w.WriteHeader(http.StatusInternalServerError)
	//	return
	//}
	//w.Write(byts)
	w.Write([]byte("<h1>Hello World!</h1>"))
	
}

func HttpPlayground() {

	db := database{"shoes": 30.0, "shirt": 15.0}
	mux := http.NewServeMux()
	mux.Handle("/home", http.HandlerFunc(db.home))
	//mux.HandleFunc("/", db.home)
	err := http.ListenAndServe("localhost:8080", mux)

	if err != nil {
		fmt.Println(err)
		return
	}
}
