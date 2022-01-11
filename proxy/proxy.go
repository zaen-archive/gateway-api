package proxy

import (
	"encoding/json"
	"gateway/configuration"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/valyala/fasthttp"
)

// type Meta struct {
// 	body    map[string]interface{}
// 	headers []byte
// }

// func (meta *Meta) bodyToBytes() []byte {
// 	val, _ := json.Marshal(meta.body)

// 	return val
// }

// func (meta *Meta) appendBodyBytes(m []byte) error {
// 	var v map[string]interface{}
// 	err := json.Unmarshal(m, &v)
// 	if err != nil {
// 		return err
// 	}

// 	meta.appendBody(v)

// 	return nil
// }

// func (meta *Meta) appendBody(m map[string]interface{}) {
// 	if meta.body == nil {
// 		meta.body = m
// 		return
// 	}

// 	for k, v := range m {
// 		meta.body[k] = v
// 	}
// }

// func (meta *Meta) appendHeader(val []byte) {
// 	meta.headers = append(meta.headers, val...)
// }

var client = &fasthttp.Client{
	NoDefaultUserAgentHeader: true,
	DisablePathNormalizing:   true,
}

func createProxy(req *fasthttp.Request, res *fasthttp.Response, endpoint string, target *configuration.EndpointTarget) error {
	// Check all the target routes
	reqUrl := string(req.URI().Path())
	reqQuery := req.URI().QueryString()

	urlTarget := target.URLTarget
	urlDest := endpoint

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
	addr := target.HostTarget + urlTarget
	if len(reqQuery) > 0 {
		// TODO: Finish Query String Cumulative
		// Check if there are some passing value into query parameter
		// if index > 0 {
		// }
		addr += "?" + string(reqQuery)
	}

	return client.Do(req, res)
}

// CreateProxyStrategy
func CreateProxyStrategy(endpoint *configuration.Endpoint) fiber.Handler {
	return func(c *fiber.Ctx) error {

		if len(endpoint.Targets) == 1 {
			return createProxy(c.Request(), c.Response(), endpoint.Endpoint, &endpoint.Targets[0])
		}

		var resBody map[string]interface{}
		for i := 0; i < len(endpoint.Targets); i++ {
			res := c.Response()
			target := &endpoint.Targets[i]

			res.ResetBody()

			err := createProxy(c.Request(), res, endpoint.Endpoint, target)
			if err != nil {
				return err
			}

			b := map[string]interface{}{}
			json.Unmarshal(res.Body(), &b)

			if resBody == nil {
				resBody = b
			} else {
				for key, val := range b {
					resBody[key] = val
				}
			}
		}

		val, _ := json.Marshal(resBody)
		c.Write(val)

		return nil
	}
}
