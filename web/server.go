package web

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"text/template"

	"github.com/go-chi/chi"
)

// DataQuery is and interface for database operations
type DataQuery interface {
	
}

// Server represents the application server
type Server struct {
	templates    map[string]*template.Template
	lock         sync.RWMutex
	db           DataQuery
}

// StartHTTPServer is the entry point for the http server
func StartHTTPServer(httpHost, httpPort string, db DataQuery) {
	server := &Server{
		templates:    map[string]*template.Template{},
		db:           db,
	}

	router := chi.NewRouter()
	workDir, _ := os.Getwd()

	filesDir := filepath.Join(workDir, "web/public/dist")
	FileServer(router, "/static", http.Dir(filesDir))
	server.registerHandlers(router)

	// load templates
	server.loadTemplates()

	address := net.JoinHostPort(httpHost, httpPort)

	log.Infof("Starting http server on %s", address)
	err := http.ListenAndServe(address, router)
	if err != nil {
		log.Error("Error starting web server")
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// FileServer creates HTTP routes for static file server
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
}
