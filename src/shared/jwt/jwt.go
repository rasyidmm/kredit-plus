package jwt

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"strings"
)

var (
	jwtKey = "d57e8fa5d223947de2f6433f401acf18085bfd487f0b09523c17a19b8499c39dcaaf856d2f91c07e43492193435cc866"
)

type AuthCtx struct {
	Id float64
}

func GenerateToken(userId int64) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": userId,
	})

	tokenString, err := token.SignedString([]byte(jwtKey))
	return tokenString, err
}
func Authentication(next fiber.Handler) fiber.Handler {

	return func(ctx *fiber.Ctx) error {
		authToken := ctx.Get("Authorization")
		if authToken == "" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"errors": "Unauthorized",
			})
		}

		splitToken := strings.Split(authToken, "Bearer ")
		token, err := jwt.Parse(splitToken[1], func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method :%v", token.Header["alg"])
			}

			return []byte(jwtKey), nil
		})

		if !token.Valid || err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"errors": "Unauthorized",
			})
		}

		var id float64
		destructID := token.Claims.(jwt.MapClaims)["id"]
		if destructID == nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"errors": "Unauthorized",
			})
		} else {

			id = destructID.(float64)
		}

		authCtx := AuthCtx{
			Id: id,
		}

		ctx.Locals("authCtx", authCtx)
		next(ctx)
		return nil
	}
}
