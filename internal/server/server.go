package server

import (
	"fmt"
	"net/http"

	api "github.com/litvinovmitch11/avito-merch-store/internal/generated"
)

var _ api.ServerInterface = (*Server)(nil)

type Server struct{}

func NewServer() Server {
	return Server{}
}

// Аутентификация и получение JWT-токена. При первой аутентификации пользователь создается автоматически.
// (POST /api/auth)
func (Server) PostApiAuth(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PostApiAuth")
}

// Купить предмет за монеты.
// (GET /api/buy/{item})
func (Server) GetApiBuyItem(w http.ResponseWriter, r *http.Request, item string) {
	fmt.Println("GetApiBuyItem")
}

// Получить информацию о монетах, инвентаре и истории транзакций.
// (GET /api/info)
func (Server) GetApiInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetApiBuyItem")
}

// Отправить монеты другому пользователю.
// (POST /api/sendCoin)
func (Server) PostApiSendCoin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PostApiSendCoin")
}
