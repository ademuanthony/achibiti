package web

import "github.com/go-chi/chi"

func (s *Server) registerHandlers(r *chi.Mux) {
	r.Get("/", s.homePage)

	// r.With(syncDataType).Get("/api/sync/{dataType}", s.sync)
	// r.With(chartTypeCtx).Get("/api/charts/{charttype}", s.chartTypeData)
}
