package middleware

import (
	"context"
	"net/http"
	"strings"

	"go-arepas/auth"

	"github.com/julienschmidt/httprouter"
)

type contextKey string

const jwtClaimsKey contextKey = "jwtClaims"

type JWTMiddleware struct {
	jwtStrategy *auth.JWTStrategy
}

func NewJWTMiddleware(jwtStrategy *auth.JWTStrategy) *JWTMiddleware {
	return &JWTMiddleware{
		jwtStrategy: jwtStrategy,
	}
}

func (m *JWTMiddleware) ValidateJWT(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "missing authorization header", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			http.Error(w, "invalid authorization header format", http.StatusUnauthorized)
			return
		}

		token, err := m.jwtStrategy.ValidateToken(tokenString)
		if err != nil {
			http.Error(w, "invalid token: "+err.Error(), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), jwtClaimsKey, token.Claims)

		next(w, r.WithContext(ctx), ps)
	}
}
