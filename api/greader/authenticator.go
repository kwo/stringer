package greader

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/kwo/stringer/models"
)

type TokenAuthenticator interface {
	// lookup userinfo given token returning userID, username, created time and any error
	AuthenticateToken(token string) (*models.User, error)
}

func authenticate(tokenAuthenticator TokenAuthenticator) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		f := func(w http.ResponseWriter, r *http.Request) {

			ctx := r.Context()

			token := getTokenFromRequest(r)
			if len(token) == 0 {
				log.Println("auth token: missing")
				http.Error(w, "no token", http.StatusUnauthorized)
				return
			}

			user, err := tokenAuthenticator.AuthenticateToken(token)
			if err != nil {
				log.Printf("auth token: cannot get user: %s\n", err)
				http.Error(w, "", http.StatusUnauthorized)
				return
			}

			ctx = context.WithValue(ctx, ContextKeySessionID, user)
			next.ServeHTTP(w, r.WithContext(ctx))

		}
		return http.HandlerFunc(f)
	}
}

func getTokenFromRequest(r *http.Request) string {
	if value := r.Header.Get(hAuthorization); len(value) != 0 {
		if strings.HasPrefix(value, authTypeGoogle) {
			return strings.TrimSpace(strings.TrimPrefix(value, authTypeGoogle))
		}
	}
	return ""
}

func getUserFromContext(ctx context.Context) *models.User {
	if value := ctx.Value(ContextKeySessionID); value != nil {
		if user, ok := value.(*models.User); ok {
			return user
		}
	}
	return nil
}
