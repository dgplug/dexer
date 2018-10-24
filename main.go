package main

import (
	"encoding/json"
	"flag"
	"fmt"
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

// IndexFile is the controller that helps with indexing the file
func IndexFile(w http.ResponseWriter, r *http.Request) {
	err := indexer.NewIndex(config, lg)
	json.NewEncoder(w).Encode(err)
	return
}

// SearchFile is the controller that helps with indexing the file
func SearchFile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	searchResult := indexer.Search(config.IndexFilename, params["query"])
	json.NewEncoder(w).Encode(searchResult.Hits)
	return
}

func main() {
	flag.Parse()
	config = conf.NewConfig(*configFlag)
	lg = logger.NewLogger("logfile")
	fmt.Println("Refreshing the index")
	err := indexer.NewIndex(config)
	lg.Must(err)
	fmt.Printf("Serving on %v \n", config.Port)
	router := mux.NewRouter()
	router.HandleFunc("/index", IndexFile).Methods("GET")
	router.HandleFunc("/search/{query}", SearchFile).Methods("GET")
	loggedRouter := handlers.LoggingHandler(lg, router)
	log.Fatal(http.ListenAndServe(config.Port, loggedRouter))
}
