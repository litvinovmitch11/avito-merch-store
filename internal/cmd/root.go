package cmd

import (
	"fmt"

	api "github.com/litvinovmitch11/avito-merch-store/internal/generated"
)

func Run() {
	_ = api.AuthRequest{}
	fmt.Println("Start!")
}
