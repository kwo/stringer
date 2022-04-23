package greader

import (
	"encoding/json"
	"log"
	"net/http"
)

type UserInfo struct {
	UserID        string `json:"userId"`
	UserName      string `json:"userName"`
	UserProfileID string `json:"userProfileId"`
	UserEmail     string `json:"userEmail"`
	IsBloggerUser bool   `json:"isBloggerUser,omitempty"`
	SignupTime    int64  `json:"signupTimeSec,omitempty"`
}

func userinfo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		if user := getUserFromContext(ctx); user != nil {
			userinfo := UserInfo{
				UserID:        user.ID,
				UserProfileID: user.ID,
				UserName:      user.Name,
				UserEmail:     user.Username,
				SignupTime:    user.Created.Unix(),
			}
			data, err := json.Marshal(&userinfo)
			if err != nil {
				log.Printf("userinfo: cannot marshal response: %s", err)
				http.Error(w, "", http.StatusInternalServerError)
				return
			}
			w.Header().Set(hContentType, mimetypeJSON)
			_, _ = w.Write(data)
			return
		}

		http.Error(w, "", http.StatusUnauthorized)
	}
}
