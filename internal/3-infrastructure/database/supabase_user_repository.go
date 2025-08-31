package database

import (
	"ShittakaKeijiban_Back/internal/1-domain/entity"
	"ShittakaKeijiban_Back/internal/3-infrastructure/config"

	"github.com/supabase-community/gotrue-go/types"
	"github.com/supabase-community/supabase-go"
)

// SupabaseUserRepository は Supabase と連携するユーザーリポジトリの実装です。
type SupabaseUserRepository struct {
	client *supabase.Client
}

// NewSupabaseUserRepository は SupabaseUserRepository のコンストラクタです。
func NewSupabaseUserRepository(cfg *config.Env) (*SupabaseUserRepository, error) {
	client, err := supabase.NewClient(cfg.SupabaseURL, cfg.SupabaseServiceRoleKey, nil)
	if err != nil {
		return nil, err
	}
	return &SupabaseUserRepository{client: client}, nil
}

// Register は Supabase で新規ユーザー登録を行います。
func (r *SupabaseUserRepository) Register(user *entity.User) error {
	_, err := r.client.Auth.Signup(types.SignupRequest{
		Email:    user.Email,
		Password: user.Password,
	})
	return err
}

// Login は Supabase でログインを行います。
func (r *SupabaseUserRepository) Login(email, password string) (string, error) {
	resp, err := r.client.Auth.Token(types.TokenRequest{
		GrantType: "password",
		Email:     email,
		Password:  password,
	})
	if err != nil {
		return "", err
	}
	return resp.AccessToken, nil
}