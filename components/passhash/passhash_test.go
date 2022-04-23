package passhash_test

import (
	"testing"

	"github.com/kwo/stringer/components/passhash"
	"github.com/kwo/stringer/models"
)

func TestPasshash(t *testing.T) {

	const testpassword = "world"

	u := &models.User{}
	if err := passhash.SetPassword(u, testpassword); err != nil {
		t.Fatal(err)
	}

	t.Log(u.PasswordHash)

	if !passhash.MatchPassword(u, testpassword) {
		t.Error("passwords do not match")
	}

}
