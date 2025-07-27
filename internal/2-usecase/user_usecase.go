// userのユースケースを定義する

package usecase

import (
	"ShittakaKeijiban_Back/internal/domain" // ドメイン層（エンティティ・リポジトリIF）をインポート
)

// UserInteractorはユーザー関連のビジネスロジックを担当する構造体
type UserInteractor struct {
	UserRepo domain.UserRepository // ユーザーリポジトリのインターフェース
}

// RegisterUserは新規ユーザー登録処理を行うユースケース
func (uc *UserInteractor) RegisterUser(user *domain.User) error {
	return nil
}
