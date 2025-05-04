package middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/vkhangstack/dlt/internal/adapters/handler"
	"github.com/vkhangstack/dlt/internal/adapters/repository"
	"github.com/vkhangstack/dlt/internal/adapters/utils"
	"github.com/vkhangstack/dlt/internal/core/domain/enum"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			handler.HandleError(c, enum.TokenNotFoundError, errors.New("token is required"))
			c.Abort()
			return
		}

		apiCfg, err := repository.LoadAPIConfig()
		if err != nil {
			handler.HandleError(c, enum.TokenNotFoundError, errors.New("token is required"))
			c.Abort()
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		if tokenString == "" {
			handler.HandleError(c, enum.TokenNotFoundError, errors.New("token is required"))
			c.Abort()
			return
		}

		userID, err := utils.ValidateAccessToken(c.Request.Header.Get("Authorization"), apiCfg.JWTSecret)

		if err != nil {
			handler.HandleError(c, enum.Unauthorized, err)
			c.Abort()
			return
		}
		if userID == "" {
			handler.HandleError(c, enum.Unauthorized, errors.New("token is required"))
			c.Abort()
			return
		}

		c.Set("userId", userID)
		c.Next()
	}
}
