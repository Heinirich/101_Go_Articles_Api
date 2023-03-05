package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      int    `json:"Id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func main() {
	Articles = []Article{
		{Id: 1, Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		{Id: 2, Title: "Hello 2", Desc: "Article Description 2", Content: "Article Content 2"},
	}
	handleRequests()
}

func handleRequests() {

	myRouters := mux.NewRouter().StrictSlash(true)
	myRouters.HandleFunc("/", homePageHandler).Methods("GET")
	myRouters.HandleFunc("/articles", getAllArticleHandler).Methods("GET")
	myRouters.HandleFunc("/article/{id}", getOneArticleHandler).Methods("GET")
	myRouters.HandleFunc("/article", postOneArticleHandler).Methods("POST")
	myRouters.HandleFunc("/article/{id}", updateOneArticleHandler).Methods("PATCH")
	myRouters.HandleFunc("/article/{id}", deleteOneArticleHandler).Methods("DELETE")

	// http.HandleFunc("/", homePageHandler)
	// http.HandleFunc("/articles",getAllArticleHandler)
	fmt.Println("Starting Server at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", myRouters))
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func getAllArticleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: allArticleHandler")
	json.NewEncoder(w).Encode(Articles)
}
func getOneArticleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	key := vars["id"]

	for _, v := range Articles {
		if strconv.Itoa(v.Id) == key {
			json.NewEncoder(w).Encode(v)
		}
	}

}

func postOneArticleHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	// fmt.Fprintf(w,"%+v",string(reqBody))
	var article Article
	json.Unmarshal(reqBody, &article)
	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(article)

}

func deleteOneArticleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// we will need to extract the `id` of the article we
	// wish to delete
	id := vars["id"]


	// we then need to loop through all our articles
	for index, article := range Articles {
		// if our id path parameter matches one of our
		// articles
		if strconv.Itoa(article.Id) == id {
			// updates our Articles array to remove the
			// article
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}

	json.NewEncoder(w).Encode(Articles)
}

func updateOneArticleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var updatedEvent Article
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &updatedEvent)
	
	for i, article := range Articles {
		fmt.Println(strconv.Itoa(article.Id),id)
		if strconv.Itoa(article.Id) == id {
			
			article.Title = updatedEvent.Title
			article.Desc = updatedEvent.Desc
			article.Content = updatedEvent.Content
			Articles[i] = article
			json.NewEncoder(w).Encode(article)
		}
	}
}
