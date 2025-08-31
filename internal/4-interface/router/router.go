package router

import (
	"ShittakaKeijiban_Back/internal/4-interface/handler"
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter は新しいルーターを作成します。
func NewRouter(userHandler *handler.UserHandler) http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/login", userHandler.Login).Methods("POST")
	return r
}
