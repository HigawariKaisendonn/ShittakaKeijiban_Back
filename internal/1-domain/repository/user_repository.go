package repository

import "ShittakaKeijiban_Back/internal/1-domain/entity"

// UserRepository はユーザーに関するデータ操作を定義するインターフェース
type UserRepository interface {
	Register(user *entity.User) error
	Login(email, password string) (string, error)
}