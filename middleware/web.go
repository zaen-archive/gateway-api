package middleware

import (
	. "gateway/configuration"

	"github.com/gofiber/fiber/v2"
)

// createWebStrategy :
func createWebStrategy(webs []Web) func(*fiber.Ctx) error {
	return nil
	// return func(c *gin.Context) {
	// 	length := len(webs)

	// 	for i := 0; i < length; i++ {
	// 		web := webs[i]
	// 		reqURI := c.Request.RequestURI

	// 		// If length of request path less than web alias
	// 		// Meaning it's not the path for web
	// 		if len(reqURI) < len(web.Alias) {
	// 			continue
	// 		}

	// 		// Check if base or beginning path is same
	// 		if web.Alias != reqURI[:len(web.Alias)] {
	// 			continue
	// 		}

	// 		// Remove beginning path
	// 		uri := reqURI[len(web.Alias):]
	// 		c.Request.RequestURI = uri
	// 		c.Request.URL.Path = uri
	// 		pathExt := filepath.Ext(uri)

	// 		// If doesn't have extension
	// 		// Mean will serve the index.html
	// 		var webPath string
	// 		if runtime.GOOS == "windows" {
	// 			webPath = web.Windowspath
	// 		} else {
	// 			webPath = web.Path
	// 		}

	// 		if pathExt == "" {
	// 			http.ServeFile(c.Writer, c.Request, path.Join(webPath, "index.html"))
	// 		} else if utils.ArrayContain(web.Extensions, pathExt[1:]) {
	// 			http.FileServer(http.Dir(webPath)).ServeHTTP(c.Writer, c.Request)
	// 		}

	// 		return
	// 	}

	// 	c.Next()
	// }
}
