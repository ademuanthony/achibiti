package web

import (
	"context"
	"errors"
	"fmt"
	"net/http"
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

func currentUserCtx(r *http.Request) (*userData, error) {
	userData, ok := r.Context().Value(ctxCurrentUser).(userData)
	if !ok {
		log.Trace("current user not set")
		return nil, errors.New("current user not set")
	}
	return &userData, nil
}

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