package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"gitlab.com/zaen/gateway/configuration"
)

// createProxyStrategy
func CreateProxyStrategy(endpoint *configuration.Endpoint) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		// targetPath := createTargetPath(string(c.Request().URI().Path()), string(c.Request().URI().FullURI()), endpoint.URLTarget)

		return proxy.Balancer(proxy.Config{
			Servers: []string{
				endpoint.HostTarget,
			},
			ModifyRequest: func(c *fiber.Ctx) error {
				c.Request().Header.Add("X-Origin-Host", endpoint.HostTarget)
				c.Request().Header.Add("X-Forwarded-Host", c.Hostname())

				return nil
			},
			ModifyResponse: func(c *fiber.Ctx) error {
				c.Response().Header.Del(fiber.HeaderServer)
				return nil
			},
		})(c)
	}
}

//create target path
// func createTargetPath(pathRequest, pathGin, pathTarget string) string {
// 	wildIndex := strings.Index(pathGin, "*")

// 	if wildIndex != -1 {
// 		wildPathReq := pathRequest[wildIndex:]
// 		wildPathGin := pathGin[wildIndex:]
// 		pathTargets := strings.FieldsFunc(pathTarget, func(r rune) bool {
// 			return r == '/'
// 		})

// 		for i, s := range pathTargets {
// 			if s[0] == '*' {
// 				if s == wildPathGin {
// 					pathTargets[i] = wildPathReq

// 					return "/" + strings.Join(pathTargets, "/")
// 				}

// 				panic("Wild Card is not same")
// 			}
// 		}
// 	}

// 	return pathTarget
// }
