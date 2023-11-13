package delivery

import (
	uc "main/usecase"
)

// @title TCRA API
// @version 1.0
// @description TCRA back server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8080
// @BasePath  /api

type Handler struct {
	usecase uc.UsecaseInterface
}

func NewHandler(uc uc.UsecaseInterface) *Handler {
	return &Handler{
		usecase: uc,
	}
}
