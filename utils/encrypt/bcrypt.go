package encrypt

import (
	"encoding/base64"
	"mcblog/utils/errno"

	"golang.org/x/crypto/bcrypt"
)

func EncryptPw(pw string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		return "", errno.New(errno.ERROR, err)
	}
	return base64.RawStdEncoding.EncodeToString(bytes), err
}
