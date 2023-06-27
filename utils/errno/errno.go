package errno

import (
	"fmt"
)

type err struct {
	code   int
	msg    string
	errord error
}

func (err err) Error() string {
	return fmt.Sprintf("ERR-code: %d, msg: %s, error: %v", err.code, err.msg, err.errord)
}

func New(code int, errord error) error {
	return err{
		code:   code,
		msg:    codemsg[code],
		errord: errord,
	}
}

func GetCode(errord error) int {
	if errord == nil {
		return SUCCESS
	}
	if e, ok := errord.(err); ok {
		return e.code
	}
	return -1
}

func GetMsg(errord error) string {
	if errord == nil {
		return codemsg[SUCCESS]
	}
	if e, ok := errord.(err); ok {
		return e.msg
	}
	return ""
}
