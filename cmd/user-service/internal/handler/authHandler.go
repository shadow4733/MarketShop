package handler

import (
	"net/http"
	"user-service/internal/dto/request"
	"user-service/internal/dto/response"
	"user-service/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *service.AuthService
}

// NewAuthHandler создает новый экземпляр AuthHandler
// @Summary Создает обработчик аутентификации
// @Description Инициализирует обработчик с сервисом аутентификации
func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// Register обрабатывает запрос на регистрацию пользователя
// @Summary Регистрация пользователя
// @Description Создает нового пользователя и возвращает JWT токен
// @Tags auth
// @Accept json
// @Produce json
// @Param registerRequest body request.RegisterRequest true "Данные для регистрации"
// @Success 200 {object} response.RegisterResponse "Успешная регистрация"
// @Failure 400 {object} response.ErrorResponse "Неверные данные запроса"
// @Failure 409 {object} response.ErrorResponse "Пользователь уже существует"
// @Failure 500 {object} response.ErrorResponse "Внутренняя ошибка сервера"
// @Router /auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var req request.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Error:   "Неверные данные запроса",
			Message: err.Error(),
		})
		return
	}

	result, err := h.authService.RegisterUser(req)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "пользователь с таким email уже существует" ||
			err.Error() == "пользователь с таким username уже существует" {
			statusCode = http.StatusConflict
		}

		c.JSON(statusCode, response.ErrorResponse{
			Error:   "Ошибка регистрации",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
