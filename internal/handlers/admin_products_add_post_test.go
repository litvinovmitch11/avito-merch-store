package handlers_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/litvinovmitch11/avito-merch-store/internal/entities"
	"github.com/litvinovmitch11/avito-merch-store/internal/handlers"
	mock_auth "github.com/litvinovmitch11/avito-merch-store/mocks/services/auth"
	mock_jwt "github.com/litvinovmitch11/avito-merch-store/mocks/services/jwt"
	mock_products "github.com/litvinovmitch11/avito-merch-store/mocks/services/products"
)

func TestPostAdminProductsAdd_SimpleTest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authService := mock_auth.NewMockAuthService(ctrl)
	JWTService := mock_jwt.NewMockJWTService(ctrl)
	productsService := mock_products.NewMockProductsService(ctrl)

	postAdminProductsAddHandler := handlers.PostAdminProductsAddHandler{
		AuthService:     authService,
		ProductsService: productsService,
		JWTService:      JWTService,
	}

	token := "token"
	userAuth := entities.UserAuth{
		Username: "Username",
		Password: "Password",
	}
	product := entities.Product{
		Id:    "Id",
		Title: "Title",
		Price: 100,
	}
	expectedProductID := "ProductID"

	gomock.InOrder(
		JWTService.EXPECT().ParseToken(token).Return(userAuth, nil),
		authService.EXPECT().AuthorizeUser(userAuth).Return("", nil),
		productsService.EXPECT().AddProduct(product).Return(expectedProductID, nil),
	)

	productID, err := postAdminProductsAddHandler.PostAdminProductsAdd(token, product)
	if err != nil {
		t.Errorf("PostAdminProductsAdd fail: %v", err)
	}

	if productID != expectedProductID {
		t.Errorf("expected: %s, real: %s", expectedProductID, productID)
	}
}
