package server

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/litvinovmitch11/avito-merch-store/internal/connections/postgresql"
	"github.com/litvinovmitch11/avito-merch-store/internal/generated/api"
	"github.com/litvinovmitch11/avito-merch-store/internal/handlers"
	authrepo "github.com/litvinovmitch11/avito-merch-store/internal/repositories/auth"
	productsrepo "github.com/litvinovmitch11/avito-merch-store/internal/repositories/products"
	storagerepo "github.com/litvinovmitch11/avito-merch-store/internal/repositories/storage"
	authservice "github.com/litvinovmitch11/avito-merch-store/internal/services/auth"
	jwtservice "github.com/litvinovmitch11/avito-merch-store/internal/services/jwt"
	productsservice "github.com/litvinovmitch11/avito-merch-store/internal/services/products"
	storageservice "github.com/litvinovmitch11/avito-merch-store/internal/services/storage"

	"github.com/rs/zerolog"
)

var _ api.ServerInterface = (*Server)(nil)

type Server struct {
	Logger *zerolog.Logger

	PostApiAuthHandler          *handlers.PostApiAuthHandler
	GetApiBuyItemHandler        *handlers.GetApiBuyItemHandler
	GetApiInfoHandler           *handlers.GetApiInfoHandler
	PostApiSendCoinHandler      *handlers.PostApiSendCoinHandler
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
	storageRepository := storagerepo.Repository{
		PostgresqlConnection: &postgresqlConnection,
	}

	// services init
	authService := authservice.Service{
		AuthRepository: &authRepository,
	}
	JWTService := jwtservice.Service{}
	productsService := productsservice.Service{
		ProductsRepository: &productsRepository,
	}
	storageService := storageservice.Service{
		AuthRepository:     &authRepository,
		ProductsRepository: &productsRepository,
		StorageRepository:  &storageRepository,
	}

	// handlers init
	postApiAuthHandler := handlers.PostApiAuthHandler{
		AuthService: &authService,
		JWTService:  &JWTService,
	}
	getApiBuyItemHandler := handlers.GetApiBuyItemHandler{
		AuthService:    &authService,
		JWTService:     &JWTService,
		StorageService: &storageService,
	}
	getApiInfoHandler := handlers.GetApiInfoHandler{
		AuthService:    &authService,
		JWTService:     &JWTService,
		StorageService: &storageService,
	}
	postApiSendCoinHandler := handlers.PostApiSendCoinHandler{
		AuthService:    &authService,
		JWTService:     &JWTService,
		StorageService: &storageService,
	}
	postAdminProductsAddHandler := handlers.PostAdminProductsAddHandler{
		AuthService:     &authService,
		ProductsService: &productsService,
		JWTService:      &JWTService,
	}

	return Server{
		Logger: logger,

		PostApiAuthHandler:          &postApiAuthHandler,
		GetApiBuyItemHandler:        &getApiBuyItemHandler,
		GetApiInfoHandler:           &getApiInfoHandler,
		PostApiSendCoinHandler:      &postApiSendCoinHandler,
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

	entity, err := postApiAuthRequestToEntity(request)
	if err != nil {
		s.Logger.Error().Err(err).Msg("fail while process PostApiAuth")

		w.WriteHeader(http.StatusBadRequest)
		msg := "bad request"

		_ = json.NewEncoder(w).Encode(api.ErrorResponse{
			Errors: &msg,
		})

		return
	}

	response, err := s.PostApiAuthHandler.PostApiAuth(entity)
	if err != nil {
		s.Logger.Error().Err(err).Msg("fail while process PostApiAuth")

		msg := ""
		if errors.Is(err, jwtservice.ErrInvalidToken) || errors.Is(err, authservice.ErrUnauthorized) {
			w.WriteHeader(http.StatusUnauthorized)
			msg = "unauthorize"
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			msg = "internal error"
		}

		_ = json.NewEncoder(w).Encode(api.ErrorResponse{
			Errors: &msg,
		})

		return
	}

	serverResponse := postApiAuthEntityToResponse(response)

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(serverResponse)
}

// Купить предмет за монеты.
// (GET /api/buy/{item})
func (s Server) GetApiBuyItem(w http.ResponseWriter, r *http.Request, item string) {
	token := r.Header.Get("Authorization")
	err := s.GetApiBuyItemHandler.GetApiBuyItem(token, item)
	if err != nil {
		s.Logger.Error().Err(err).Msg("fail while process GetApiBuyItem")

		msg := ""
		if errors.Is(err, jwtservice.ErrInvalidToken) || errors.Is(err, authservice.ErrUnauthorized) {
			w.WriteHeader(http.StatusUnauthorized)
			msg = "unauthorize"
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			msg = "internal error"
		}

		_ = json.NewEncoder(w).Encode(api.ErrorResponse{
			Errors: &msg,
		})

		return
	}

	w.WriteHeader(http.StatusOK)
}

// Получить информацию о монетах, инвентаре и истории транзакций.
// (GET /api/info)
func (s Server) GetApiInfo(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	response, err := s.GetApiInfoHandler.GetApiInfo(token)
	if err != nil {
		s.Logger.Error().Err(err).Msg("fail while process GetApiInfo")

		msg := ""
		if errors.Is(err, jwtservice.ErrInvalidToken) || errors.Is(err, authservice.ErrUnauthorized) {
			w.WriteHeader(http.StatusUnauthorized)
			msg = "unauthorize"
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			msg = "internal error"
		}

		_ = json.NewEncoder(w).Encode(api.ErrorResponse{
			Errors: &msg,
		})

		return
	}

	serverResponse := getApiInfoEntityToResponse(response)

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(serverResponse)
}

// Отправить монеты другому пользователю.
// (POST /api/sendCoin)
func (s Server) PostApiSendCoin(w http.ResponseWriter, r *http.Request) {
	var request api.SendCoinRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		s.Logger.Error().Err(err).Msg("fail before process PostApiSendCoin")
		return
	}

	token := r.Header.Get("Authorization")
	entity, err := postSendCoinRequestToEntity(request)
	if err != nil {
		s.Logger.Error().Err(err).Msg("fail while process PostApiSendCoin")

		w.WriteHeader(http.StatusBadRequest)
		msg := "bad request"

		_ = json.NewEncoder(w).Encode(api.ErrorResponse{
			Errors: &msg,
		})

		return
	}

	err = s.PostApiSendCoinHandler.PostApiSendCoin(token, entity)
	if err != nil {
		s.Logger.Error().Err(err).Msg("fail while process PostApiSendCoin")

		msg := ""
		if errors.Is(err, jwtservice.ErrInvalidToken) || errors.Is(err, authservice.ErrUnauthorized) {
			w.WriteHeader(http.StatusUnauthorized)
			msg = "unauthorize"
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			msg = "internal error"
		}

		_ = json.NewEncoder(w).Encode(api.ErrorResponse{
			Errors: &msg,
		})

		return
	}

	w.WriteHeader(http.StatusOK)
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

	token := r.Header.Get("Authorization")
	entity, err := postAdminProductsAddToEntity(request)
	if err != nil {
		s.Logger.Error().Err(err).Msg("fail while process PostAdminProductsAdd")

		w.WriteHeader(http.StatusBadRequest)
		msg := "bad request"

		_ = json.NewEncoder(w).Encode(api.ErrorResponse{
			Errors: &msg,
		})

		return
	}

	response, err := s.PostAdminProductsAddHandler.PostAdminProductsAdd(token, entity)
	if err != nil {
		s.Logger.Error().Err(err).Msg("fail while process PostAdminProductsAdd")

		msg := ""
		if errors.Is(err, jwtservice.ErrInvalidToken) || errors.Is(err, authservice.ErrUnauthorized) {
			w.WriteHeader(http.StatusUnauthorized)
			msg = "unauthorize"
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			msg = "internal error"
		}

		_ = json.NewEncoder(w).Encode(api.ErrorResponse{
			Errors: &msg,
		})

		return
	}

	serverResponse := postAdminProductsAddEntityToResponse(response)

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(serverResponse)
}
