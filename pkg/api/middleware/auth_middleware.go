package middleware

import (
	"camping-backend-with-go/internal/domain/entity"
	"camping-backend-with-go/internal/domain/presenter"
	"camping-backend-with-go/pkg/config"
	"camping-backend-with-go/pkg/util"
	"strings"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func Protected() fiber.Handler {

	return jwtware.New(jwtware.Config{
		ErrorHandler:   jwtError,
		SuccessHandler: successHandler,
		SigningKey:     jwtware.SigningKey{Key: []byte(config.Config("JWT_SECRET"))},
		AuthScheme:     "Bearer",
		TokenLookup:    "header:Authorization",
	})
}

func successHandler(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	db := c.Locals("db").(*gorm.DB)

	var user entity.User
	db.Where("id = ?", userId).First(&user)

	c.Locals("is_authenticated", true)
	c.Locals("request_user", user)
	return c.Next()
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "missing or malformed JWT" {
		jsonResponse := presenter.NewJsonResponse(true, err.Error(), nil)
		return c.Status(fiber.StatusBadRequest).JSON(jsonResponse)
	}

	jsonResponse := presenter.NewJsonResponse(false, "Invalid or expired JWT", nil)
	return c.Status(fiber.StatusUnauthorized).JSON(jsonResponse)
}

func RequestAuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// 기본적으로 인증되지 않은 상태로 설정
		c.Locals("is_authenticated", false)
		c.Locals("request_user", "anonymous")

		// Authorization 헤더에서 토큰 추출
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Next()
		}

		// Bearer 접두사 제거
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			// "Bearer" 접두사가 없으면 잘못된 형식
			return c.Next()
		}

		// 토큰 검증

		token, err := validateToken(tokenString)
		if err != nil {
			// 토큰이 유효하지 않으면 다음 미들웨어로 넘어감
			return c.Next()
		}

		// 토큰이 유효하면 사용자 정보 설정
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Next()
		}

		userId, ok := claims["user_id"].(float64)

		db := c.Locals("db").(*gorm.DB)

		var user entity.User
		if err := db.First(&user, int(userId)).Error; err != nil {
			return c.Next()
		}

		c.Locals("is_authenticated", true)
		c.Locals("request_user", user)

		return c.Next()

	}
}

func validateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(config.Config("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrTokenInvalidId
	}

	return token, nil
}

func IsAuthenticatedOrReadOnly() fiber.Handler {
	return func(c *fiber.Ctx) error {
		method := c.Method()

		// GET 요청은 모두 허용
		if method == fiber.MethodGet {
			return c.Next()
		}

		authHeader := c.Get("Authorization")
		if authHeader == "" || !isAuthenticated(authHeader, c) {
			jsonResponse := presenter.NewJsonResponse(true, "인증되지 않은 사용자입니다.", nil)
			return c.Status(fiber.StatusUnauthorized).JSON(jsonResponse)
		}

		return c.Next()
	}
}

func isAuthenticated(bearerToken string, context ...*fiber.Ctx) bool {

	tokenString := strings.TrimPrefix(bearerToken, "Bearer ")
	if bearerToken == tokenString {
		// Bearer 값이 없으면 error
		return false
	}

	token, err := validateToken(tokenString)
	if err != nil {
		return err == nil
	}

	c, err := util.ContextParser(context...)
	util.HandleFunc(err)

	db, ok := c.Locals("db").(*gorm.DB)
	if !ok {
		return false
	}

	userId, ok := token.Claims.(jwt.MapClaims)["user_id"].(float64)
	if !ok {
		return false
	}

	var user entity.User
	if err := db.First(&user, userId).Error; err != nil {
		return false
	}

	return true
}
