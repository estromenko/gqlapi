package server

import (
	"context"
	"fmt"
	"gqlapi/utils"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func (s *Server) baseMiddleware(endpoint http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		s.logger.Info(r.UserAgent() + ". " + r.Method + ". " + time.Now().String())

		endpoint(w, r)
	}
}

func (s *Server) authenticationMiddleware(endpoint http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenHeader := r.Header.Get("Authorization")

		splitedToken := strings.Split(tokenHeader, " ")

		if len(splitedToken) != 2 {
			ctx := context.WithValue(r.Context(), utils.ContextValue("tokenError"), "Not authorized")
			endpoint(w, r.WithContext(ctx))
			return
		}

		token, err := jwt.Parse(splitedToken[1], func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method, got " + t.Header["alg"].(string))
			}
			return []byte(s.config.User.JWTSecret), nil
		})

		if err != nil {
			ctx := context.WithValue(r.Context(), utils.ContextValue("tokenError"), err.Error())
			endpoint(w, r.WithContext(ctx))
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		id := fmt.Sprintf("%v", claims["id"])

		user, err := s.db.User().FindByID(id)
		if err != nil {
			ctx := context.WithValue(r.Context(), utils.ContextValue("tokenError"), err.Error())
			endpoint(w, r.WithContext(ctx))
			return
		}

		ctx := context.WithValue(r.Context(), utils.ContextValue("user"), user)
		endpoint(w, r.WithContext(ctx))
	}
}
