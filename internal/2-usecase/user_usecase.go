package usecase

import "ShittakaKeijiban_Back/internal/1-domain/entity"

// UserUsecase はユーザー関連のユースケースのインターフェースです。
type UserUsecase interface {
	Register(user *entity.User) error
	Login(email, password string) (string, error)
}