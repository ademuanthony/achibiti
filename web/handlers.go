package web

import (
	"net/http"
)


// /home
func (s *Server) homePage(res http.ResponseWriter, req *http.Request) {
	
	data := map[string]interface{}{
		"mempoolCount": 2,
		"blocksCount":  3,
		"votesCount":   2,
		"powCount":     6,
		"vspCount":     1,
		"exchangeTick": 5,
	}

	s.render("home.html", data, res)
}
