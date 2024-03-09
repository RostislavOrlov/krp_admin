package request

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func GetRequest[T any](c *gin.Context) (T, bool) {
	var request T
	logrus.Info("запрос: ", request)

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return request, false
	}

	return request, true
}
