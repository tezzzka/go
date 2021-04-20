package main

import (
	"strconv"
	// "reflect"
	"fmt"
	"log"
	"net/http"
	"Repo"
	// "strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

// func itemsHandler(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case http.MethodGet:
// 		fmt.Fprintf(w,Repo.GetAllItems(),r.URL.Path[1:])
// 	case http.MethodPost:
// 		Repo.CreateRecord()
// 	default:
// 		panic("error")
// 	}
// }

// func itemHandler(w http.ResponseWriter, r *http.Request) {
// 	// fmt.Println(r.URL)
// 	fmt.Fprintf(w, "Hi, I love %s", r.URL.Path[1:])
// 	switch r.Method {
// 	case http.MethodGet:
// 		fmt.Fprintf(w, "%s", Repo.GetById(r.URL.Path[1:]))
// 	case http.MethodDelete:
// 	case http.MethodPut:
// 	default:
// 	}
// }

func main() {
	fmt.Println("Console message:api server has started.")

	http.HandleFunc("/", handler)

	http.HandleFunc("/items",func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			fmt.Fprintf(w,Repo.GetAllItems(),r.URL.Path[1:])
		case http.MethodPost:
			Repo.CreateRecord()
		default:
			panic("error")
		}
	})

	http.HandleFunc("/items/",func(w http.ResponseWriter, r *http.Request) {
		i, _ := strconv.Atoi(r.URL.Path[8:])
		switch r.Method {
		case http.MethodGet:
			fmt.Fprintf(w,Repo.GetById(i),r.URL.Path[1:])
		case http.MethodDelete:
			Repo.DelItem(i)
		case http.MethodPut:
		default:
			fmt.Fprintf(w, "Error in request. Please, Try again.")
			panic("error")
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
