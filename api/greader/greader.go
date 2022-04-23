package greader

import (
	"github.com/go-chi/chi/v5"
)

type ContextKey string

const ContextKeySessionID ContextKey = "sessionID"

const (
	authReasonBad     = "BadAuthentication" // login
	authReasonUnknown = "Unknown"           // login
	authTypeGoogle    = "GoogleLogin auth="
	formkeyUsername   = "Email"  // login
	formkeyPassword   = "Passwd" // login
	hAuthorization    = "Authorization"
	hContentType      = "Content-Type"
	mimetypeJSON      = "application/json"
	// mimetypeText          = "text/plain; charset=utf-8"
	mimetypeTextNoCharset = "text/plain"
)

type Provider interface {
	LoginAuthenticator
	TokenAuthenticator
	SubscriptionListProvider
	ItemListProvider
}

func New(provider Provider) *chi.Mux {
	mux := chi.NewMux()
	mux.Post("/accounts/ClientLogin", login(provider))

	mux.Route("/reader/api/0", func(r chi.Router) {
		r.Use(authenticate(provider))
		r.Get("/user-info", userinfo())
		r.Get("/subscription/list", subscriptionList(provider))
		r.Get("/stream/items/ids", itemsIDs(provider))
	})

	return mux
}
