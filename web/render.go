package web

import (
	"encoding/json"
	"net"
	"net/http"
	"strings"
)

func (s *Server) render(tplName string, data map[string]interface{}, w http.ResponseWriter, r *http.Request) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	if data == nil {
		data = map[string]interface{}{}
	}

	data["currentUser"], _ = currentUserCtx(r)

	if tpl, ok := s.templates[tplName]; ok {
		err := tpl.Execute(w, data)
		if err == nil {
			return
		}
		// Filter out broken pipe (user pressed "stop") errors
		if _, ok := err.(*net.OpError); ok {
			if strings.Contains(err.Error(), "broken pipe") {
				return
			}
		}
		log.Errorf("Error executing template: %s", err.Error())
		return
	}

	log.Errorf("Template %s is not registered", tplName)
}

func (s *Server) renderError(errorMessage string, w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"error": errorMessage,
	}
	s.render("error.html", data, w, r)
}

func (s *Server) renderErrorJSON(errorMessage string, res http.ResponseWriter) {
	data := map[string]interface{}{
		"error": errorMessage,
	}

	s.renderJSON(data, res)
}

func (s *Server) renderJSON(data interface{}, res http.ResponseWriter) {
	d, err := json.Marshal(data)
	if err != nil {
		log.Errorf("Error marshalling data: %s", err.Error())
	}

	res.Header().Set("Content-Type", "application/json")
	res.Write(d)
}

// writeJSONBytes prepares the headers for pre-encoded JSON and writes the JSON
// bytes.
func (s *Server) renderJSONBytes(data []byte, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, err := w.Write(data)
	if err != nil {
		log.Warnf("ResponseWriter.Write error: %v", err)
	}
}
