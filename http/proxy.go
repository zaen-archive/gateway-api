package http

import (
	"encoding/json"
	"gateway/configuration"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/valyala/fasthttp"
)

type Meta struct {
	resp map[string]interface{}
}

func (meta *Meta) toBytes() []byte {
	val, _ := json.Marshal(meta.resp)

	return val
}

func (meta *Meta) appendBytes(m []byte) error {
	var v map[string]interface{}
	err := json.Unmarshal(m, &v)
	if err != nil {
		return err
	}

	meta.append(v)

	return nil
}

func (meta *Meta) append(m map[string]interface{}) {
	if meta.resp == nil {
		meta.resp = m
		return
	}

	for k, v := range m {
		meta.resp[k] = v
	}
}

var client = &fasthttp.Client{
	NoDefaultUserAgentHeader: true,
	DisablePathNormalizing:   true,
}

// createProxyStrategy
func createProxyStrategy(endpoint *configuration.Endpoint) fiber.Handler {
	return func(c *fiber.Ctx) error {
		meta := Meta{}

		// Check all the target routes
		for _, target := range endpoint.Targets {
			reqUrl := string(c.Request().URI().Path())
			reqQuery := c.Request().URI().QueryString()

			urlTarget := target.URLTarget
			urlDest := endpoint.Endpoint

			// Check Params
			reqSegments := strings.Split(reqUrl, "/")
			for _, val := range target.ParamsIndex {
				oldSegment := target.Segments[val]
				newSegment := reqSegments[val]

				urlTarget = strings.ReplaceAll(urlTarget, oldSegment, newSegment)
				urlDest = strings.Replace(urlDest, oldSegment, newSegment, 1)
			}

			// Check Star
			if target.IsStar {
				urlTarget = utils.TrimRight(urlTarget, '*')
				if len(reqUrl) >= len(urlDest) {
					urlTarget += string(reqUrl[len(urlDest)-1:])
				}
			}

			// Set Address Target
			// TODO: Finish Query String Cumulative
			addr := target.HostTarget + urlTarget
			if len(reqQuery) > 0 {
				// Check if there are some passing value into query parameter
				// if index > 0 {
				// }
				// addr += "?" + string(reqQuery)
			}

			// Request to Host Target
			res := fasthttp.Response{}
			req := fasthttp.Request{
				Header: c.Request().Header,
			}
			method := target.Method
			if method == "ALL" {
				method = string(c.Request().Header.Method())
			}
			req.Header.SetMethod(method)
			req.SetRequestURI(addr)

			// Request set body
			if method == fasthttp.MethodPost || method == fasthttp.MethodPut {
				req.SetBody(c.Request().Body())
			}

			if err := client.Do(&req, &res); err != nil {
				return err
			}

			meta.appendBytes(res.Body())
		}

		c.Response().Header.Del(fiber.HeaderServer)
		return nil
	}
}
