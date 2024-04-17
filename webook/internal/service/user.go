package service

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"my-geektime-basic/webook/internal/domain"
	"my-geektime-basic/webook/internal/repository"
)

var (
	ErrDuplicateEmail        = repository.ErrDuplicateUser
	ErrInvalidUserOrPassword = errors.New("用户不存在或者密码不对")
)

type UserService interface {
	Signup(ctx context.Context, u domain.User) error
	Login(ctx context.Context, email string, password string) (domain.User, error)
	UpdateNonSensitiveInfo(ctx context.Context, user domain.User) error
	FindById(ctx context.Context, uid int64) (domain.User, error)
	FindOrCreate(ctx context.Context, phone string) (domain.User, error)
	FindOrCreateByWechat(ctx context.Context, info domain.WechatInfo) (domain.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func (svc *userService) Signup(ctx context.Context, u domain.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return svc.repo.Create(ctx, u)
}

func (svc *userService) Login(ctx context.Context, email string, password string) (domain.User, error) {
	u, err := svc.repo.FindByEmail(ctx, email)
	if err == repository.ErrUserNotFound {
		return domain.User{}, ErrInvalidUserOrPassword
	}
	if err != nil {
		return domain.User{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return domain.User{}, ErrInvalidUserOrPassword
	}
	return u, nil

}

func (svc *userService) UpdateNonSensitiveInfo(ctx context.Context, user domain.User) error {
	return svc.repo.UpdateNonZeroFields(ctx, user)
}

func (svc *userService) FindById(ctx context.Context, uid int64) (domain.User, error) {
	return svc.repo.FindById(ctx, uid)
}

func (svc *userService) FindOrCreate(ctx context.Context, phone string) (domain.User, error) {
	u, err := svc.repo.FindByPhone(ctx, phone)
	if err != repository.ErrUserNotFound {
		return u, err
	}
	err = svc.repo.Create(ctx, domain.User{
		Phone: phone,
	})
	if err != nil && err != repository.ErrDuplicateUser {
		return domain.User{}, err
	}
	return svc.repo.FindByPhone(ctx, phone)
}

func (svc *userService) FindOrCreateByWechat(ctx context.Context, wechatInfo domain.WechatInfo) (domain.User, error) {
	u, err := svc.repo.FindByWechat(ctx, wechatInfo.OpenId)
	if err != repository.ErrUserNotFound {
		return u, err
	}
	zap.L().Info("新用户", zap.Any("wechatInfo", wechatInfo))
	err = svc.repo.Create(ctx, domain.User{
		WechatInfo: wechatInfo,
	})
	if err != nil && err != repository.ErrDuplicateUser {
		return domain.User{}, err
	}
	return svc.repo.FindByWechat(ctx, wechatInfo.OpenId)
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}
