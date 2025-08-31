package usecase

import (
	"ShittakaKeijiban_Back/internal/1-domain/entity"
	"ShittakaKeijiban_Back/internal/1-domain/repository"
)

// UserInteractor はユーザー関連のビジネスロジックを担当します。
type UserInteractor struct {
	UserRepository repository.UserRepository
}

// NewUserInteractor は UserInteractor のコンストラクタです。
func NewUserInteractor(userRepository repository.UserRepository) UserUsecase {
	return &UserInteractor{UserRepository: userRepository}
}

// Register は新規ユーザー登録を行います。
func (ui *UserInteractor) Register(user *entity.User) error {
	return ui.UserRepository.Register(user)
}

// Login はログイン処理を行います。
func (ui *UserInteractor) Login(email, password string) (string, error) {
	return ui.UserRepository.Login(email, password)
}