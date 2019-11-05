package web

import (
	"net/http"
)

// /home
func (s *Server) homePage(w http.ResponseWriter, r *http.Request) {
	
	data := map[string]interface{}{
		"title": "Home",
	}

	s.render("home.html", data, w, r)
}

func (s *Server) errorPage(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"title": "Server Error",
	}

	s.render("error.html", data, w, r)
}

func (s *Server) loginPage(w http.ResponseWriter, r *http.Request)  {
	data := map[string]interface{}{
		"title": "Account Login",
	}

	s.render("login.html", data, w, r)
}
