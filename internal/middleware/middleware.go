package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"
)

func CheckAuthHeaderMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		logrus.Info("старт работы middleware")
		authHeader := c.GetHeader("Authorization")
		logrus.Info("authorization header: ", authHeader)
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Missing Authorization header",
			})
			c.Abort()
			return
		}

		authHeaderParts := strings.Split(authHeader, " ")
		if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid Authorization header format",
			})
			c.Abort()
			return
		}

		token := authHeaderParts[1]
		fmt.Println("Token:", token)
		c.Set("access_token", token)
		c.Next()
	}
}

type TokenClaims struct {
	jwt.StandardClaims
}

func CheckAccessTokenExpiresMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.MustGet("access_token").(string)

		mySigningKey := []byte("SecretKey")

		accessTokenWithClaims, err := jwt.ParseWithClaims(accessToken, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
			return mySigningKey, nil
		})
		if err != nil {
			logrus.Warning("error parsing access token: ", err.Error())
		}

		if claimsAccess, okAccess := accessTokenWithClaims.Claims.(*TokenClaims); okAccess && accessTokenWithClaims.Valid {
			fmt.Printf("%v %v", claimsAccess.StandardClaims.ExpiresAt)
			//accessTokenClaims := jwt.StandardClaims{
			//	ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			//	IssuedAt:  time.Now().Unix(),
			//}

			if claimsAccess.ExpiresAt < time.Now().Unix() {
				c.JSON(http.StatusUnauthorized, gin.H{"message": "access token expired"})
				c.Abort()
			}
		}
		c.Next()
	}
}
