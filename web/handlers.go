package web

import (
	"net/http"
	"time"

	accounts "github.com/ademuanthony/achibiti/accounts/proto/accounts"
	"github.com/gofrs/uuid"
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

func (s *Server) apiLogin(w http.ResponseWriter, r *http.Request)  {
	if err := r.ParseForm(); err != nil {
		s.renderErrorJSON("error in parsing request", w)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	loginResp, err := s.accountService.Login(r.Context(), &accounts.LoginRequest{
		Username:             username,
		Password:             password,
	})

	if err != nil {
		s.renderErrorJSON(err.Error(), w)
		return
	}

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
