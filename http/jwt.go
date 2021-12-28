package http

import (
	"encoding/json"
	"errors"
	"net/url"
	"strings"

	"gateway/configuration"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

var secretKey []byte = []byte("the most secreet text")

const secretIdentifier = "x-ident"

func createJwtStrategy(endpoint *configuration.Endpoint) fiber.Handler {
	return func(c *fiber.Ctx) error {

		if endpoint.Jwt == "" {
			return c.Next()
		}

		jwtStr := c.Cookies(endpoint.Jwt)
		if len(jwtStr) == 0 {
			return errors.New("no authenticated yet")
		}

		// Remove Escape, becaues fasthttp not parsing the query
		jwtStr, _ = url.QueryUnescape(jwtStr)

		// Remove if Sign Token
		if jwtStr[0:2] == "s:" {
			jwtStr = jwtStr[2:]
		}

		// Check if you get more token there
		if split := strings.Split(jwtStr, "."); len(split) > 3 {
			jwtStr = strings.Join(split[:3], ".")
		}

		claims := &jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(jwtStr, claims, func(t *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})
		if err != nil {
			return err
		}

		if !token.Valid {
			return errors.New("token is not valid")
		}

		if (*claims)[secretIdentifier] == nil {
			return errors.New("token identity is not valid")
		}

		for key, v := range *claims {
			if val, ok := v.(string); ok {
				c.Request().Header.Set("x-user-"+key, val)
			} else {
				val, _ := json.Marshal(v)
				c.Request().Header.Set("x-user-"+key, string(val))
			}
		}

		return c.Next()
	}
}
