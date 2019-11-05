package web

import (
	"context"
	"fmt"
	"github.com/ademuanthony/achibiti/acl/handler"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
	"time"

	"github.com/gofrs/uuid"
)

type key int

const (
	ctxCurrentUser key = iota

	sessionCookieName string = "session_token"
)

var tokenExpiryTime = 15 * time.Minute

func (s *Server) requireLogin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// We can obtain the session token from the requests cookies, which come with every request
		c, err := r.Cookie(sessionCookieName)
		if err != nil {
			// If the cookie is not set, return an unauthorized status
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		sessionToken := c.Value

		// We then get the name of the user from our cache, where we set the session token
		loginData, err := s.cache.Do("GET", sessionToken)
		if err != nil {
			// If there is an error fetching from cache, return an internal server error status
			http.Redirect(w, r, "/error", http.StatusSeeOther)
			return
		}
		if loginData == nil {
			// If the session token is not present in cache, return an unauthorized error
			http.Redirect(w, r, "/error", http.StatusSeeOther)
			return
		}

		ctx := context.WithValue(r.Context(), ctxCurrentUser, loginData)

		next.ServeHTTP(w, r.WithContext(ctx))

	})
}

func (s *Server) jwtAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		notAuth := []string{"/api/user/login"} //List of endpoints that doesn't require auth
		requestPath := r.URL.Path //current request path

		//check if request does not need authentication, serve the request if it doesn't need it
		for _, value := range notAuth {

			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		tokenHeader := r.Header.Get("Authorization") //Grab the token from the header

		if tokenHeader == "" { //Token is missing, returns with error code 403 Unauthorized
			s.renderErrorJSON("Missing auth token", w)
			return
		}

		splitted := strings.Split(tokenHeader, " ") //The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
		if len(splitted) != 2 {
			s.renderErrorJSON("Invalid/Malformed auth token", w)
			return
		}

		tokenPart := splitted[1] //Grab the token part, what we are truly interested in
		claims := &handler.Claims{}

		token, err := jwt.ParseWithClaims(tokenPart, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("234dfsgk593jffjdh9ekjdsfjk43089432kjkfjfadj4390fdjk3490dgskljgdsk2390gshgfddfhjk2398-glsjl"), nil
		})

		if err != nil { //Malformed token, returns with http code 403 as usual
			s.renderErrorJSON("Malformed authentication token", w)
			return
		}

		if !token.Valid { //Token is invalid, maybe not signed on this server
			s.renderErrorJSON("Token is not valid.", w)
			return
		}

		userData := userData{
			Username:    claims.Username,
			Role:        claims.Role,
			Token:       tokenPart,
		}
		//Everything went well, proceed with the request and set the caller to the user retrieved from the parsed token
		ctx := context.WithValue(r.Context(), ctxCurrentUser, userData)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r) //proceed in the middleware chain!
	})
}

func currentUserCtx(r *http.Request) *userData {
	userData, ok := r.Context().Value(ctxCurrentUser).(userData)
	if !ok {
		log.Trace("current user not set")
		return nil
	}
	return &userData
}

// used this for refresh token and stop frontend from making request for fresh
func (s *Server) refreshLoginSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie(sessionCookieName)
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		sessionToken := c.Value

		response, err := s.cache.Do("GET", sessionToken)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if response == nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// (END) The code uptil this point is the same as the first part of the `Welcome` route

		// Now, create a new session token for the current user
		newSessionToken, err := uuid.NewV4()
		if err != nil {
			log.Trace(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, err = s.cache.Do("SETEX", newSessionToken.String(), tokenExpiryTime, fmt.Sprintf("%s",response))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Delete the older session token
		_, err = s.cache.Do("DEL", sessionToken)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Set the new token as the users `session_token` cookie
		http.SetCookie(w, &http.Cookie{
			Name:    sessionCookieName,
			Value:   newSessionToken.String(),
			Expires: time.Now().Add(tokenExpiryTime),
		})

		next.ServeHTTP(w, r)
	})
}