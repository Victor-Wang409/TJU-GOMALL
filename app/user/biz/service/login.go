package service

import (
	"context"
	"errors"

	"github.com/cloudwego/biz-demo/gomall/app/user/biz/dal/mysql"
	"github.com/cloudwego/biz-demo/gomall/app/user/biz/model"
	user "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginResq, err error) {
	// Finish your business logic.
	if req.Email == "" || req.Password == "" {
		return nil, errors.New("password or email is empty")
	}

	row, err := model.GetbyEmail(mysql.DB, req.Email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(row.PasswordHashed), []byte(req.Password))
	if err != nil {
		return nil, err
	}

	resp = &user.LoginResq{
		UserId: int32(row.ID),
	}

	return resp, nil
}
