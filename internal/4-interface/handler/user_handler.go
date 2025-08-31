package handler

import (
	"ShittakaKeijiban_Back/internal/2-usecase"
	"ShittakaKeijiban_Back/internal/4-interface/dto"
	"encoding/json"
	"net/http"
)

// UserHandler はユーザー関連のリクエストを処理します。
type UserHandler struct {
	UserUsecase usecase.UserUsecase
}

// NewUserHandler は UserHandler のコンストラクタです。
func NewUserHandler(userUsecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{UserUsecase: userUsecase}
}

// Login はログイン処理を行います。
func (uh *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := uh.UserUsecase.Login(req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
