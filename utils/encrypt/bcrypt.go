package encrypt

import (
	"mcblog/utils/errno"

	"golang.org/x/crypto/bcrypt"
)

func EncryptPw(pw string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		return "", errno.New(errno.ERROR, err)
	}
	return string(bytes), err
}
