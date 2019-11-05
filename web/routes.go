package web

import (
	"github.com/go-chi/chi"
)

func (s *Server) registerHandlers(r *chi.Mux) {
	r.Get("/error", s.loginPage)
	r.Get("/login", s.loginPage)

	r.Group(func(r chi.Router) {
		r.Post("/api/auth/login", s.apiLogin)

		r.With(s.jwtAuthentication).Group(func(r chi.Router) {
			r.Get("/", s.homePage) // todo home should redirect to client side home

			//Auth
			r.Post("/api/auth/refresh-token", s.refreshToken)

			//API
			// image
			r.Post("/api/images", s.uploadFile)

			// departments
			r.Get("/api/departments", s.departments)
			r.Post("/api/departments", s.createDepartment)

			// employee types
			r.Get("/api/employee-types", s.employeeTypes)
			r.Post("/api/employee-types", s.createEmployeeType)
			r.Put("/api/employee-types", s.updateEmployeeType)

			//employees
			r.Get("/api/employees", s.employees)
			r.Post("/api/employees", s.createEmployee)
		})
	})
}
