package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/cloudwego/biz-demo/gomall/app/user/biz/dal/mysql"
	"github.com/cloudwego/biz-demo/gomall/app/user/biz/model"
	user "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResq, err error) {
	// Finish your business logic.
	fmt.Printf(req.Email)
	fmt.Println(req.Password)
	fmt.Println(req.PasswordConfirm)
	if req.Email == "" || req.Password == "" || req.PasswordConfirm == "" {
		return nil, errors.New("password or email is empty")
	}
	if req.Password != req.PasswordConfirm {
		return nil, errors.New("password not match")
	}
	passwardHashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	newUser := &model.User{
		Email:          req.Email,
		PasswordHashed: string(passwardHashed),
	}

	err = model.Create(mysql.DB, newUser)
	if err != nil {
		return nil, err
	}
	return &user.RegisterResq{UserId: int32(newUser.ID)}, err
}
