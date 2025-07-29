package middleware

import (
	"fmt"
	"os"

	"github.com/LaviqueDias/api-crud-go/src/configuration/logger"
	"github.com/LaviqueDias/api-crud-go/src/configuration/rest_err"
	"github.com/LaviqueDias/api-crud-go/src/user/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
)

func RequireAuth(userService service.UserServiceInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("Authorization")
		if err != nil {
			restErr := rest_err.NewUnauthorizedError("authorization cookie not found")
			logger.Error("Cookie not found", err,
				zap.String("journey", "RequireAuth"),
			)
			c.JSON(restErr.Code, restErr)
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("SECRET")), nil
		})

		if err != nil || !token.Valid {
			restErr := rest_err.NewUnauthorizedError("invalid or expired token")
			logger.Error("Invalid token", err,
				zap.String("journey", "RequireAuth"),
			)
			c.JSON(restErr.Code, restErr)
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || claims["sub"] == nil {
			restErr := rest_err.NewUnauthorizedError("invalid token claims")
			logger.Error("Invalid claims in token", nil,
				zap.String("journey", "RequireAuth"),
			)
			c.JSON(restErr.Code, restErr)
			c.Abort()
			return
		}

		userEmail, ok := claims["sub"].(string)
		if !ok {
			restErr := rest_err.NewUnauthorizedError("invalid token subject")
			logger.Error("Invalid sub format in token", nil,
				zap.String("journey", "RequireAuth"),
			)
			c.JSON(restErr.Code, restErr)
			c.Abort()
			return
		}

		logger.Info("User email extracted from token", zap.String("email", userEmail))

		user, restErr := userService.FindUserByEmail(userEmail)
		if restErr != nil {
			logger.Error("User not found in middleware", restErr,
				zap.String("journey", "RequireAuth"),
			)
			c.JSON(restErr.Code, restErr)
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}

