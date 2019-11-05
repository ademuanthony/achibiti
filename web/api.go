package web

import (
	"encoding/json"
	"fmt"
	acl "github.com/ademuanthony/achibiti/acl/proto/acl"
	hr "github.com/ademuanthony/achibiti/hr/proto/hr"
	"github.com/gofrs/uuid"
	"net/http"
	"strconv"
	"time"
)

// auth/login
func (s *Server) apiLogin(w http.ResponseWriter, r *http.Request)  {
	if err := r.ParseForm(); err != nil {
		s.renderErrorJSON("error in parsing request", w)
		return
	}

	var loginRequest acl.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		log.Trace(err)
		s.renderErrorJSON("error in reading request body", w)
		return
	}

	loginResp, err := s.accountService.Login(r.Context(), &loginRequest)

	if err != nil {
		s.renderErrorJSON(err.Error(), w)
		return
	}

	s.storeAndSendUserSession(loginResp, w)
}

// auth/refresh-token
func (s *Server) refreshToken(w http.ResponseWriter, r *http.Request) {
	userData := currentUserCtx(r)
	if userData == nil {
		s.renderErrorJSON("Please login to continue", w)
		return
	}

	loginResp, err := s.accountService.RefreshToken(r.Context(), &acl.RefreshTokenRequest{Username: userData.Username})
	if err != nil {
		s.renderErrorJSON(err.Error(), w)
		return
	}

	s.storeAndSendUserSession(loginResp, w)
}

func (s *Server) storeAndSendUserSession(loginResp *acl.LoginResponse, w http.ResponseWriter) {
	// Create a new random session token
	sessionToken, err := uuid.NewV4()
	if err != nil {
		s.renderErrorJSON("Internal server error. Please try again later", w)
		return
	}

	loginData := userData{
		Username:    loginResp.Username,
		Email:       loginResp.Email,
		PhoneNumber: loginResp.PhoneNumber,
		Name:        loginResp.Name,
		Role:        loginResp.Role,
		Token:       loginResp.Token,
	}

	// Set the token in the cache, along with the user whom it represents
	// The token has an expiry time of 120 seconds
	_, err = s.cache.Do("SETEX", sessionToken.String(), tokenExpiryTime.Seconds(), loginData)
	if err != nil {
		// If there is an error in setting the cache, return an internal server error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Finally, we set the client cookie for "session_token" as the session token we just generated
	// we also set an expiry time of 120 seconds, the same as the cache
	http.SetCookie(w, &http.Cookie {
		Name:    sessionCookieName,
		Value:   sessionToken.String(),
		Expires: time.Now().Add(tokenExpiryTime),
	})

	var data = map[string]interface{}{
		"user": loginResp,
	}

	s.renderJSON(data, w)
}

// departments
func (s *Server) createDepartment(w http.ResponseWriter, r *http.Request) {
	var request hr.CreateDepartmentRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Trace(err)
		s.renderErrorJSON("cannot create department, error in reading request body", w)
		return
	}

	response, err := s.hrService.CreateDepartment(r.Context(), &request)
	if err != nil {
		s.renderErrorJSON(fmt.Sprintf("cannot create department, %s", err.Error()), w)
		return
	}

	s.renderJSON(map[string]string{"id": response.GetId()}, w)
}

func (s *Server) departments(w http.ResponseWriter, r *http.Request) {
	log.Trace("getting departments")
	if err := r.ParseForm(); err != nil {
		log.Trace(err)
		s.renderErrorJSON("error is parsing request", w)
		return
	}

	skipCount, _ := strconv.Atoi(r.FormValue("skip_count"))
	resultCount, _ := strconv.Atoi(r.FormValue("result_count"))

	request := &hr.DepartmentsRequest{
		SkipCount:            int32(skipCount),
		MaxResultCount:       int32(resultCount),
	}

	response, err := s.hrService.Departments(r.Context(), request)
	if err != nil {
		log.Trace(err)
		s.renderErrorJSON(fmt.Sprintf("cannot fetch departments, %s", err.Error()), w)
		return
	}

	data := map[string]interface{} {
		"total_count": response.GetTotalCount(),
		"items": response.GetDepartments(),
	}

	s.renderJSON(data, w)
}

// employee type
func (s *Server) createEmployeeType(w http.ResponseWriter, r *http.Request) {
	var request hr.CreateEmployeeTypeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Trace(err)
		s.renderErrorJSON("cannot create employee type, error in reading request body", w)
		return
	}

	response, err := s.hrService.CreateEmployeeType(r.Context(), &request)
	if err != nil {
		s.renderErrorJSON(fmt.Sprintf("cannot create employee, %s", err.Error()), w)
		return
	}

	s.renderJSON(map[string]string{"id": response.GetEmployeeTypeId()}, w)
}

func (s *Server) employeeTypes(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Trace(err)
		s.renderErrorJSON("error is parsing request", w)
		return
	}

	skipCount, _ := strconv.Atoi(r.FormValue("skip_count"))
	resultCount, _ := strconv.Atoi(r.FormValue("result_count"))

	request := &hr.EmployeeTypesRequest {
		SkipCount:            int32(skipCount),
		MaxResultCount:       int32(resultCount),
	}

	response, err := s.hrService.EmployeeTypes(r.Context(), request)
	if err != nil {
		log.Trace(err)
		s.renderErrorJSON(fmt.Sprintf("cannot fetch employee types, %s", err.Error()), w)
		return
	}

	data := map[string]interface{} {
		"total_count": response.GetTotalCount(),
		"items": response.GetEmployeeTypes(),
	}

	s.renderJSON(data, w)
}

func (s *Server) updateEmployeeType(w http.ResponseWriter, r *http.Request) {
	var request hr.UpdateEmployeeTypeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Trace(err)
		s.renderErrorJSON("cannot update employee type, error in reading request body", w)
		return
	}

	_, err := s.hrService.UpdateEmployeeType(r.Context(), &request)
	if err != nil {
		s.renderErrorJSON(fmt.Sprintf("cannot update employee, %s", err.Error()), w)
		return
	}

	s.renderJSON(map[string]string{"id": request.GetId()}, w)
}

// employee
func (s *Server) createEmployee(w http.ResponseWriter, r *http.Request) {
	var createEmployeeRequest hr.CreateEmployeeRequest
	if err := json.NewDecoder(r.Body).Decode(&createEmployeeRequest); err != nil {
		s.renderErrorJSON("cannot create employee, error in reading request body " + err.Error(), w)
		return
	}

	createEmployeeResponse, err := s.hrService.CreateEmployee(r.Context(), &createEmployeeRequest)
	if err != nil {
		s.renderErrorJSON(err.Error(), w)
		return
	}

	s.renderJSON(map[string]string{"id": createEmployeeResponse.GetEmployeeId()}, w)
}

func (s *Server) employees(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Trace(err)
		s.renderErrorJSON("error is parsing request", w)
		return
	}

	skipCount, _ := strconv.Atoi(r.FormValue("skip_count"))
	resultCount, _ := strconv.Atoi(r.FormValue("result_count"))

	request := &hr.EmployeesRequest {
		DepartmentId: r.FormValue("department_id"),
		EmployeeTypeId: r.FormValue("employee_type_id"),
		SkipCount:            int32(skipCount),
		MaxResultCount:       int32(resultCount),
	}

	response, err := s.hrService.Employees(r.Context(), request)
	if err != nil {
		log.Trace(err)
		s.renderErrorJSON(fmt.Sprintf("cannot fetch employees, %s", err.Error()), w)
		return
	}

	data := map[string]interface{} {
		"total_count": response.GetTotalCount(),
		"items": response.GetEmployees(),
		"user_data": currentUserCtx(r),
	}

	s.renderJSON(data, w)
}