package middleware

import (
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"gitlab.com/zaen/gateway/configuration"
)

const (
	secretKey = "the most secreet text"
)

type claimHelper struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

func CreateJwtStrategy(endpoint *configuration.Endpoint) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		if endpoint.Jwt == "" {
			return nil
		}

		jwtStr := c.Cookies(endpoint.Jwt)
		if len(jwtStr) == 0 {
			c.Set("error", "No cookie "+endpoint.Jwt+" found!")
			return nil
		}

		// Check if there is 's:' in first character
		// if exist remove it
		if jwtStr[0:2] == "s:" {
			jwtStr = jwtStr[2:]
		}

		// Check if you get more token there
		if split := strings.Split(jwtStr, "."); len(split) > 3 {
			jwtStr = strings.Join(split[:3], ".")
		}

		claims := &claimHelper{}
		token, err := jwt.ParseWithClaims(jwtStr, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

		if err != nil {
			c.Set("error", err.Error())
			return nil
		}

		if !token.Valid {
			c.Set("error", "token tidak valid")
			return nil
		}

		c.Set("X-User", claims.ID)
		c.Set("X-User-ExpiredAt", string(rune(claims.ExpiresAt)))

		return nil
	}
}
