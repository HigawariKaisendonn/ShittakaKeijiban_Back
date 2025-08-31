package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

// Env は環境変数を保持します。
type Env struct {
    SupabaseURL        string
    SupabaseServiceRoleKey string
    SupabaseJWTKey     string
    Port               string
    AppEnv             string
    TestJWT            string
}

// NewEnv は環境変数を読み込んでEnv構造体を返します。
func NewEnv() *Env {
    // .envファイルから環境変数を読み込む
    log.Println("Loading .env file")
    err := godotenv.Overload()
    if err != nil {
        log.Println("Error loading .env file, using environment variables")
    }
    log.Printf("Port from env: %s", os.Getenv("PORT"))

    return &Env{
        SupabaseURL:        os.Getenv("SUPABASE_URL"),
        SupabaseServiceRoleKey: os.Getenv("SUPABASE_SERVICE_ROLE_KEY"),
        SupabaseJWTKey:     os.Getenv("SUPABASE_JWT_KEY"),
        Port:               os.Getenv("PORT"),
        AppEnv:             os.Getenv("APP_ENV"),
        TestJWT:            os.Getenv("TESTJWT"),
    }
}
