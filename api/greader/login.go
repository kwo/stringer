package greader

import (
	"fmt"
	"log"
	"net/http"
)

type LoginAuthenticator interface {
	Authenticate(username, password string) (string, bool, error)
}

func login(loginAuthenticator LoginAuthenticator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if err := r.ParseForm(); err != nil {
			log.Printf("login: cannot parse form: %s\n", err)
			sendLoginResponseError(w, "", authReasonUnknown)
			return
		}

		username := r.Form.Get(formkeyUsername)
		password := r.Form.Get(formkeyPassword)

		token, ok, err := loginAuthenticator.Authenticate(username, password)
		if err != nil {
			log.Printf("login error: %s\n", err)
			sendLoginResponseError(w, "", authReasonUnknown)
			return
		} else if !ok {
			log.Printf("login failure: %s\n", username)
			sendLoginResponseError(w, "", authReasonBad)
			return
		}

		log.Printf("login success: %s\n", username)
		sendLoginResponse(w, token)
	}
}

func sendLoginResponse(w http.ResponseWriter, token string) {
	w.Header().Set(hContentType, mimetypeTextNoCharset)
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, "SID=%s\n", token)
	_, _ = fmt.Fprintf(w, "LSID=%s\n", token)
	_, _ = fmt.Fprintf(w, "Auth=%s\n", token)
}

func sendLoginResponseError(w http.ResponseWriter, url, reason string) {
	w.Header().Set(hContentType, mimetypeTextNoCharset)
	w.WriteHeader(http.StatusForbidden)
	_, _ = fmt.Fprintf(w, "Url=%s\n", url)
	_, _ = fmt.Fprintf(w, "Error=%s\n", reason)
}
