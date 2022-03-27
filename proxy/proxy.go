package proxy

import (
	"encoding/json"
	"fmt"
	"gateway/configuration"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"github.com/gofiber/fiber/v2/utils"
)

const INTERNAL_SERVER = 500

// proxify
func proxify(c *fiber.Ctx, endpoint *configuration.Endpoint, targetIndex int, meta *Meta) error {
	// Check all the target routes
	reqUrl := string(c.Request().URI().Path())
	reqQuery := c.Request().URI().QueryString()

	target := endpoint.Targets[targetIndex]
	urlTarget := target.URLTarget
	urlDest := endpoint.Endpoint

	// Check Params Route
	reqSegments := strings.Split(reqUrl, "/")
	for _, val := range endpoint.ParamsIndex {
		oldSegment := endpoint.Segments[val]
		newSegment := reqSegments[val]

		urlTarget = strings.ReplaceAll(urlTarget, oldSegment, newSegment)
		urlDest = strings.Replace(urlDest, oldSegment, newSegment, 1)
	}

	// Check Params Sequence
	if targetIndex > 0 {

		// Set Last Body
		var body map[string]interface{}
		if endpoint.Merge {
			body = meta.body
		}

		// Change Sequential Parameters
		targetSeg := strings.Split(urlTarget, "/")
		for _, val := range targetSeg {
			if strings.HasPrefix(val, "$$") {
				if body == nil {
					json.Unmarshal(c.Response().Body(), &body)
				}

				bodyVal := body[val[2:]]
				if newVal, ok := bodyVal.(string); ok {
					urlTarget = strings.ReplaceAll(urlTarget, val, newVal)
				}
			}
		}
	}

	// Check Star
	if target.IsStar {
		urlTarget = utils.TrimRight(urlTarget, '*')
		if len(reqUrl) >= len(urlDest) {
			urlTarget += string(reqUrl[len(urlDest)-1:])
		}
	}

	// Set Address Target
	addr := target.HostTarget + urlTarget
	if len(reqQuery) > 0 {
		addr += "?" + string(reqQuery)
	}

	return proxy.Do(c, addr)
}

// CreateProxyStrategy
func CreateProxyStrategy(endpoint *configuration.Endpoint) fiber.Handler {
	return func(c *fiber.Ctx) error {

		fmt.Println(c.Hostname())

		count := len(endpoint.Targets)
		if count == 1 {
			return proxify(c, endpoint, 0, nil)
		}

		meta := Meta{}
		for i := 0; i < count; i++ {
			err := proxify(c, endpoint, i, &meta)
			statusCode := c.Response().StatusCode()
			if err != nil || statusCode == INTERNAL_SERVER {
				c.Response().SetStatusCode(200)
				return err
			}

			if endpoint.Merge {
				meta.appendBodyBytes(c.Response().Body(), endpoint.DeepMerge)
			}
		}

		if endpoint.Merge {
			c.Response().SetBody(meta.bodyToBytes())
		}

		return nil
	}
}
