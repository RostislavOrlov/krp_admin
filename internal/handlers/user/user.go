package user

import (
	"encoding/json"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"krp_admin/internal/dto"
	"krp_admin/internal/handlers/request"
	"krp_admin/internal/middleware"
	"krp_admin/internal/services"
	"net/http"
)

type UserHandler struct {
	userService *services.UserService
	engine      *gin.Engine
}

func (h *UserHandler) InitRoutes() {

	h.engine.POST("/employees", middleware.CheckAuthHeaderMiddleware(),
		middleware.CheckAccessTokenExpiresMiddleware(), h.AddUser)
	h.engine.PATCH("/employees", middleware.CheckAuthHeaderMiddleware(),
		middleware.CheckAccessTokenExpiresMiddleware(), h.EditUser)
	h.engine.DELETE("/employees", middleware.CheckAuthHeaderMiddleware(),
		middleware.CheckAccessTokenExpiresMiddleware(), h.DeleteUser)
	h.engine.GET("/employees", middleware.CheckAuthHeaderMiddleware(),
		middleware.CheckAccessTokenExpiresMiddleware(), h.ListUsers)

	h.engine.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:8081"},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Origin"},
	}))
}

func NewUserHandler(srv *services.UserService, engine *gin.Engine) (*UserHandler, error) {
	h := &UserHandler{
		userService: srv,
		engine:      engine,
	}
	h.InitRoutes()
	return h, nil
}

func (h *UserHandler) AddUser(c *gin.Context) {
	req, ok := request.GetRequest[dto.AddUserRequest](c)
	logrus.Debug(req)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "add user request error", "text": ok})
		return
	}

	client := http.Client{}

	reqRegAuthService, err := http.NewRequest("POST", "http://localhost:8080/register", nil)
	resp, err := client.Do(reqRegAuthService)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	var respBody dto.AddUserResponse
	err = json.Unmarshal(body, &respBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error", "text": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": respBody})
}

func (h *UserHandler) EditUser(c *gin.Context) {
	req, ok := request.GetRequest[dto.EditUserRequest](c)
	logrus.Debug(req)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "edit user request error", "text": ok})
		return
	}

	resp, err := h.userService.EditUser(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error (edit user)", "text": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": resp})
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	req, ok := request.GetRequest[dto.DeleteUserRequest](c)
	logrus.Debug(req)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "delete user request error", "text": ok})
		return
	}

	resp, err := h.userService.DeleteUser(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error (delete user)", "text": err.Error(), "deleted": false})
	}

	c.JSON(http.StatusOK, gin.H{"data": resp, "deleted": true})
}

func (h *UserHandler) ListUsers(c *gin.Context) {
	req, ok := request.GetRequest[dto.DeleteUserRequest](c)
	logrus.Debug(req)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "list users request error", "text": ok})
		return
	}
	users, err := h.userService.ListUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error (list users)", "text": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}
