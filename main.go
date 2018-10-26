package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/farhaanbukhsh/file-indexer/conf"
	"github.com/farhaanbukhsh/file-indexer/indexer"
	"github.com/farhaanbukhsh/file-indexer/logger"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var configFlag = flag.String("config", "NULL", "To pass a different configuration file")

var config conf.Configuration
var lg *logger.Logger
var templates []string

func init() {
	flag.Parse()
	lg = logger.NewLogger("logfile")
	config = conf.NewConfig(*configFlag, lg)

	templates = []string{
		"ui/index.html",
		"ui/layout/header.html",
		"ui/layout/footer.html",
		"ui/layout/search.html",
		"ui/layout/new_index.html",
	}
}

// RootHandler is the controller responsible for the frontend
func RootHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(templates...)
	lg.Must(err, "Template Parsed Successfully")
	t.ExecuteTemplate(w, "index", nil)
}

// IndexFile is the controller that helps with indexing the file
func IndexFile(w http.ResponseWriter, r *http.Request) {
	err := indexer.NewIndex(config, lg)
	json.NewEncoder(w).Encode(err)
}

// SearchFile is the controller that helps with indexing the file
func SearchFile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	searchResult := indexer.Search(config.IndexFilename, params["query"])
	json.NewEncoder(w).Encode(searchResult.Hits)
}

func main() {
	fmt.Println("Refreshing the index")
	err := indexer.NewIndex(config, lg)
	lg.Must(err, "Index Succesfully Created")
	fmt.Printf("Serving on %v \n", config.Port)
	router := mux.NewRouter()
	router.HandleFunc("/", RootHandler)
	router.HandleFunc("/index", IndexFile).Methods("GET")
	router.HandleFunc("/search/{query}", SearchFile).Methods("GET")
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./ui/"))))
	log.Fatal(http.ListenAndServe(config.Port, handlers.LoggingHandler(lg, router)))
}
