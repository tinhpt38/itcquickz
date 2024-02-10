package auth

import (
	model "changeme/app/model"
	"errors"
)

type LoginI interface {
	Login(idnum string) (user model.User, err error)
}

type LoginService struct {
}

func (ls *LoginService) Login(idnum string) (user model.User, err error) {
	if idnum == "v" {
		user = model.User{
		}

	} else {
		err = errors.New("không tìm thấy thí sinh")
	}

	return
}
