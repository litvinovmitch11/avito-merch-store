package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/litvinovmitch11/avito-merch-store/internal/connections/postgresql"
	"github.com/litvinovmitch11/avito-merch-store/internal/generated/api"
	"github.com/litvinovmitch11/avito-merch-store/internal/handlers"
	authrepo "github.com/litvinovmitch11/avito-merch-store/internal/repositories/auth"
	productsrepo "github.com/litvinovmitch11/avito-merch-store/internal/repositories/products"
	authservice "github.com/litvinovmitch11/avito-merch-store/internal/services/auth"
	productsservice "github.com/litvinovmitch11/avito-merch-store/internal/services/products"
	"github.com/rs/zerolog"
)

var _ api.ServerInterface = (*Server)(nil)

type Server struct {
	Logger                      *zerolog.Logger
	PostApiAuthHandler          *handlers.PostApiAuthHandler
	PostAdminProductsAddHandler *handlers.PostAdminProductsAddHandler
}

func NewServer(
	logger *zerolog.Logger,
) Server {
	// connections init
	postgresqlConnection := postgresql.Connection{}

	// repositories init
	authRepository := authrepo.Repository{
		PostgresqlConnection: &postgresqlConnection,
	}
	productsRepository := productsrepo.Repository{
		PostgresqlConnection: &postgresqlConnection,
	}

	// services init
	authService := authservice.Service{
		AuthRepository: &authRepository,
	}
	productsService := productsservice.Service{
		ProductsRepository: &productsRepository,
	}

	// handlers init
	postApiAuthHandler := handlers.PostApiAuthHandler{
		AuthService: &authService,
	}
	postAdminProductsAddHandler := handlers.PostAdminProductsAddHandler{
		ProductsService: &productsService,
	}

	return Server{
		Logger:                      logger,
		PostApiAuthHandler:          &postApiAuthHandler,
		PostAdminProductsAddHandler: &postAdminProductsAddHandler,
	}
}

// Аутентификация и получение JWT-токена. При первой аутентификации пользователь создается автоматически.
// (POST /api/auth)
func (s Server) PostApiAuth(w http.ResponseWriter, r *http.Request) {
	var request api.AuthRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		s.Logger.Error().Err(err).Msg("fail before process PostApiAuth")
		return
	}

	entity := postApiAuthRequestToEntity(request)
	response, err := s.PostApiAuthHandler.PostApiAuth(entity)
	if err != nil {
		s.Logger.Error().Err(err).Msg("fail while process PostApiAuth")
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

// Добавление нового мерча.
// (POST /admin/products/add)
func (s Server) PostAdminProductsAdd(w http.ResponseWriter, r *http.Request) {
	var request api.ProductAddRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		s.Logger.Error().Err(err).Msg("fail before process PostAdminProductsAdd")
		return
	}

	entity := postAdminProductsAddToEntity(request)
	response, err := s.PostAdminProductsAddHandler.PostAdminProductsAdd(entity)
	if err != nil {
		s.Logger.Error().Err(err).Msg("fail while process PostAdminProductsAdd")
		return
	}

	serverResponse := postAdminProductsAddEntityToResponse(response)

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(serverResponse)
}
