package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	api "github.com/litvinovmitch11/avito-merch-store/internal/generated"
	"github.com/litvinovmitch11/avito-merch-store/internal/handlers"
	"github.com/litvinovmitch11/avito-merch-store/internal/services/auth"
)

var _ api.ServerInterface = (*Server)(nil)

type Server struct {
	PostApiAuthHandler handlers.PostApiAuthHandler
}

func NewServer() Server {
	return Server{
		PostApiAuthHandler: handlers.PostApiAuthHandler{
			AuthService: auth.Service{},
		},
	}
}

// Аутентификация и получение JWT-токена. При первой аутентификации пользователь создается автоматически.
// (POST /api/auth)
func (s Server) PostApiAuth(w http.ResponseWriter, r *http.Request) {
	var request api.AuthRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		return
	}

	entity := postApiAuthRequestToEntity(request)
	response, err := s.PostApiAuthHandler.PostApiAuth(entity)
	if err != nil {
		return
	}

	serverResponse := postApiAuthEntityToResponse(response)

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(serverResponse)
}

// Купить предмет за монеты.
// (GET /api/buy/{item})
func (s Server) GetApiBuyItem(w http.ResponseWriter, r *http.Request, item string) {
	fmt.Println("GetApiBuyItem")
}

// Получить информацию о монетах, инвентаре и истории транзакций.
// (GET /api/info)
func (s Server) GetApiInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetApiBuyItem")
}

// Отправить монеты другому пользователю.
// (POST /api/sendCoin)
func (s Server) PostApiSendCoin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PostApiSendCoin")
}
