package request

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func GetRequest[T any](c *gin.Context) (T, bool) {
	var request T
	//var myMap map[string]interface{}
	//logrus.Info("запрос (пустой интерфейс): ", c.BindJSON(&myMap))
	logrus.Info("gin context: ", c)
	if err := c.BindJSON(&request); err != nil {
		logrus.Info("запрос: ", request)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return request, false
	}
	logrus.Info("запрос: ", request)

	return request, true
}
