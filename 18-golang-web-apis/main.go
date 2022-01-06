package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

/// model for courses - file

type Course struct {
	CourseId string `json:"courseid"`
	CourseName string `json:"coursename"`
	Price int `json:"price"`
	Author Author `json:"author"`
}

type Author struct{
	Fullname string `json:"fullname"`
	Website string `json:"website"`
}


var courses []Course 

func main(){

	fmt.Println("go routing api")
	router := mux.NewRouter()
	router.HandleFunc("/",getAllcourses).Methods("GET")
	router.HandleFunc("/add",addNewCourse).Methods("POST")
	router.HandleFunc("/add",errorMessage).Methods("GET")
	http.ListenAndServe(":5000",router)
}

func getAllcourses(w http.ResponseWriter, r *http.Request){
	
	// content := Course{"test","test",10,Author{"test","Test"}}
	json.NewEncoder(w).Encode(courses)
}

func addNewCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println(r.URL.Query())
	newCourse := Course{"test","test",10,Author{"test","test"}}
	courses = append(courses,newCourse)
	json.NewEncoder(w).Encode(courses)
}


func errorMessage(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>please use post method</h1>"))
}