package web

import "github.com/go-chi/chi"

func (s *Server) registerHandlers(r *chi.Mux) {
	r.Get("/error", s.loginPage)
	r.Get("/login", s.loginPage)
	r.Get("/api/login", s.apiLogin)

	r.With(s.requireLogin, s.refreshLoginSession).Group(func(r chi.Router) {
		r.Get("/", s.homePage)
	})
}
