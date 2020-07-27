package session

import (
	"github.com/Yamiyo/account/db"
	"github.com/Yamiyo/account/glob"
	"github.com/Yamiyo/account/models/dto"
)

func init() {
	glob.AutoRegister(&sessionService{})
}

type sessionService struct {
	db    db.Databases
	cache db.Databases
}

func (ss *sessionService) Login(user *dto.UserDTO) {

}

func (ss *sessionService) Logout(user *dto.UserDTO) {

}

func (ss *sessionService) Register(user *dto.UserDTO) {

}

func (ss *sessionService) CheckRole(user *dto.UserDTO) Role {

}
