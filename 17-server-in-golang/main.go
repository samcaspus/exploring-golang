package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)


func main(){

	fmt.Println("server in golang using gorilla mux")
	router := mux.NewRouter()
	router.HandleFunc("/",HomePage).Methods("GET")
	http.ListenAndServe(":5000",router)

}


func HomePage(w http.ResponseWriter, r *http.Request){

	w.Write([]byte("<h1>hello world</h1>"))
}