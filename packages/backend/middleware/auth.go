package middleware

import (
	"errors"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var ErrInvalidUserId = errors.New("invalid user id")

const (
	authHeader = "Authorization"
	userIdCtx  = "userId"
)

func AuthRequired(c *gin.Context) {
	header := c.Request.Header.Get(authHeader)
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "Unauthorized",
		})
		return
	}

	if headerParts[0] != "Bearer" || headerParts[1] == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "Invalid auth header",
		})
		return
	}

	userId, err := verifyAccessToken(headerParts[1])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "Invalid access token",
		})
		return
	}

	c.Set(userIdCtx, userId)
}

func GetUserIdFromGinCtx(c *gin.Context) (int, error) {
	userId, ok := c.Get(userIdCtx)
	if !ok {
		return 0, ErrInvalidUserId
	}

	userIdInt, ok := userId.(int)
	if !ok {
		return 0, ErrInvalidUserId
	}

	return userIdInt, nil
}

func verifyAccessToken(tokenStr string) (int, error) {
	supplyKeyVerificationFunc := func(
		token *jwt.Token,
	) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid jwt token signing method")
		}
		return []byte(os.Getenv("ACCESS_TOKEN_SECRET")), nil
	}

	token, err := jwt.ParseWithClaims(
		tokenStr,
		&jwt.StandardClaims{},
		supplyKeyVerificationFunc,
	)
	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		return strconv.Atoi(claims.Subject)
	}
	return 0, errors.New("token is invalid")
}
