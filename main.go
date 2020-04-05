package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type task struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Content string `json:"content"`
}

type allTasks []task

var tasks = allTasks {
	{
		Id: 1,
		Name: "pepito",
		Content: "Some Content",
	},
}

func indexRoute(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to my api")
}

func getTasks(w http.ResponseWriter, r *http.Request)  {
	enc := json.NewEncoder(w)
	d := map[string]int{"apple": 5, "lettuce": 7}
	fmt.Println(d)
	enc.Encode(tasks)
}

func main()  {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/tasks", getTasks)
	log.Fatal(http.ListenAndServe(":3400", router))
}


















