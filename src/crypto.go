package elysium

import (
	_ "crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"math/rand"
)

func hash(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	sha := h.Sum(nil)
	encoded := base64.StdEncoding.EncodeToString([]byte(sha))
	return string(encoded)
}

func generateSalt() string {
	b := make([]rune, PW_SALT_LENGTH)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func getSalt(un string) (bool, string) {
	var salt string
	err := DB.QueryRow("SELECT user_salt FROM users WHERE user_name=?", un).Scan(&salt)
	if err != nil {
		return false, ""
	} else {
		return true, salt
	}
}
