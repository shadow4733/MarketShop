package handler

import (
	"net/http"
	"user-service/internal/dto/request"
	"user-service/internal/dto/response"
	"user-service/internal/handler/helpers"
	"user-service/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
}

// NewUserHandler создает новый экземпляр AuthHandler
// @Summary Создает обработчик аутентификации
// @Description Инициализирует обработчик с сервисом аутентификации
func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// CreateNewUser обрабатывает запрос на регистрацию пользователя
// @Summary Регистрация пользователя
// @Description Создает нового пользователя и возвращает JWT токен
// @Tags users
// @Accept json
// @Produce json
// @Param registerRequest body request.RegisterRequest true "Данные для регистрации"
// @Success 200 {object} response.RegisterResponse "Успешная регистрация"
// @Failure 400 {object} response.ErrorResponse "Неверные данные запроса"
// @Failure 409 {object} response.ErrorResponse "Пользователь уже существует"
// @Failure 500 {object} response.ErrorResponse "Внутренняя ошибка сервера"
// @Router /user/create [post]
func (h *UserHandler) CreateNewUser(c *gin.Context) {
	var req request.Create

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Error:   "Неверные данные запроса",
			Message: err.Error(),
		})
		return
	}

	result, err := h.userService.Create(c.Request.Context(), &req)
	if err != nil {
		c.JSON(helpers.GetStatusCode(err), response.ErrorResponse{
			Error:   "Ошибка регистрации",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
