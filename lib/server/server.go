package server

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/dgplug/dexer/lib/conf"
	"github.com/dgplug/dexer/lib/indexer"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var templates []string

func init() {
	templates = []string{
		"ui/index.html",
		"ui/layout/header.html",
		"ui/layout/footer.html",
		"ui/layout/search.html",
	}
}

// Server Data Structure for holding the configuration and logger
type Server struct {
	conf.Configuration
}

// RootHandler is the controller responsible for the frontend
func (s *Server) RootHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(templates...)
	s.Must(err, "Template Parsed Successfully")
	t.ExecuteTemplate(w, "index", nil)
}

// SearchFile is the controller that helps with indexing the file
func (s *Server) SearchFile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	searchResult := indexer.Search(s.IndexFilename, params["query"])
	json.NewEncoder(w).Encode(searchResult.Hits)
}

// Start function starts the server
func (s *Server) Start() {
	s.Must(nil, "Serving on "+s.Port)
	router := mux.NewRouter()
	router.HandleFunc("/", s.RootHandler)
	router.HandleFunc("/search/{query}", s.SearchFile).Methods("GET")
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./ui/"))))
	s.Must(http.ListenAndServe(s.Port, handlers.LoggingHandler(s.GetLogger(), router)), "")
}

// NewServer function creates a new server and return a pointer to it
func NewServer(c conf.Configuration) *Server {
	temp := Server{c}
	return &temp
}
