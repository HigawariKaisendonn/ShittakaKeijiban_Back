// インターフェース定義

package domain

// UserRepositoryはユーザーに関するデータ操作を定義するインターフェース
// 実装はDBやAPIなどに依存し、ドメイン層では具体的な保存方法を意識せずに利用できます。
type UserRepository interface {
	Register(user *User) error
}
