package main

import (
	"ShittakaKeijiban_Back/internal/2-usecase"
	"ShittakaKeijiban_Back/internal/3-infrastructure/config"
	"ShittakaKeijiban_Back/internal/3-infrastructure/database"
	"ShittakaKeijiban_Back/internal/4-interface/handler"
	"ShittakaKeijiban_Back/internal/4-interface/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// 設定を読み込む
	cfg := config.NewEnv()

	// 依存関係を注入する
	userRepository, err := database.NewSupabaseUserRepository(cfg)
	if err != nil {
		log.Fatal(err)
	}

	userUsecase := usecase.NewUserInteractor(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)

	// ルーターを設定する
	router := router.NewRouter(userHandler)

	// サーバーを起動する
	port := cfg.Port
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), router); err != nil {
		log.Fatal(err)
	}
}
