package passhash

import (
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"

	"github.com/kwo/stringer/models"
)

// MatchPassword checks if the given password matches the user password.
func MatchPassword(user *models.User, password string) bool {
	hashedPassword, err := hex.DecodeString(user.PasswordHash)
	if err != nil {
		return false
	}
	return bcrypt.CompareHashAndPassword(hashedPassword, []byte(password)) == nil
}

// SetPassword updates the user password hash.
func SetPassword(user *models.User, password string) error {
	bhash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PasswordHash = hex.EncodeToString(bhash)
	return nil
}
