package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	mw "github.com/blyndusk/salika-pagination/internal/middlewares"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	headersOk := handlers.AllowedHeaders([]string{"*"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	r.HandleFunc("/movies", Movies).Methods("GET")
	r.HandleFunc("/count_pages", Counter).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), handlers.CORS(originsOk, headersOk, methodsOk)(r)))
}

func Movies(w http.ResponseWriter, r *http.Request) {
	reqbody := struct {
		Limit   uint   `json:"limit"`
		Offset  uint   `json:"offset"`
		OrderBy string `json:"order_by"`
		Asc     string `json:"asc"`
	}{
		0,
		0,
		"",
		"",
	}
	err := json.NewDecoder(r.Body).Decode(&reqbody)
	if err != nil {
		reqbody.Limit = 15
		reqbody.Offset = 0
	}

	switch reqbody.OrderBy {
	case "":
		reqbody.OrderBy = "film_table.title"
	case "title":
		reqbody.OrderBy = "film_table.title"
	case "category":
		reqbody.OrderBy = "film_table.category"
	case "rental":
		reqbody.OrderBy = "total_rental"
	default:
		json.NewEncoder(w).Encode("wrong order_by key")
		return
	}

	switch reqbody.Asc {
	case "":
		reqbody.Asc = "ASC"
	case "asc":
		reqbody.Asc = "ASC"
	case "desc":
		reqbody.Asc = "DESC"
	default:
		json.NewEncoder(w).Encode("wrong asc key")
		return
	}

	movies, err := mw.GetMoviesWithPages(reqbody.Asc, reqbody.OrderBy, reqbody.Limit, reqbody.Offset)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	json.NewEncoder(w).Encode(movies)
}

func Counter(w http.ResponseWriter, r *http.Request) {
	reqbody := struct {
		Limit int `json:"limit"`
	}{
		0,
	}
	err := json.NewDecoder(r.Body).Decode(&reqbody)
	if err != nil {
		log.Println(err.Error())
		reqbody.Limit = 15
	}

	count, err := mw.CountPages(reqbody.Limit)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	}
	json.NewEncoder(w).Encode(count)
}
