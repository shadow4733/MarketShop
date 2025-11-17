package handler

import (
	"context"
	"github.com/google/uuid"
	"net/http"
	"user-service/internal/dto/request"
	"user-service/internal/dto/response"
	"user-service/internal/handler/helpers"
	"user-service/internal/model"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	GetByID(ctx context.Context, id uuid.UUID) (*model.User, error)
	GetByUsername(ctx context.Context, username string) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	Create(ctx context.Context, data *request.Create) (*response.User, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type UserHandler struct {
	userService UserService
}

func NewUserHandler(userService UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// CreateNewUser обрабатывает запрос на регистрацию пользователя
// @Summary Регистрация пользователя
// @Description Создает нового пользователя
// @Tags users
// @Accept json
// @Produce json
// @Param request body request.Create true "Данные для регистрации"
// @Success 201 {object} response.User "Успешная регистрация"
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
