package middleware

import (
	"let-you-cook/domain/dto"
	"let-you-cook/utils/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")

		if token == "" {
			c.JSON(http.StatusUnauthorized, dto.Resp{
				Status:  http.StatusUnauthorized,
				Message: "unauthorized",
				Data:    nil,
			})
			c.Abort()
			return
		}

		token = strings.TrimPrefix(token, "Bearer ")

		user, err := jwt.ParseToken(token)

		if err != nil {
			c.JSON(http.StatusUnauthorized, dto.Resp{
				Status:  http.StatusUnauthorized,
				Message: "unauthorized",
				Data:    nil,
			})
			c.Abort()
			return
		}

		c.Set("user_id", user.Id)
		c.Set("email", user.Email)
		c.Set("user", user)
		c.Next()
	}
}
