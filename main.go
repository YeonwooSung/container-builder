package main

import (
	"flag"
	"fmt"
	"log"

	"container-builder/routers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html"
)

var port = flag.Int("p", 8080, "서버가 Listen할 port 번호를 입력해주세요.")

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	// Initialize standard Go html template engine
	engine := html.New("./public/views", ".html")
	addr := fmt.Sprintf(":%d", *port)
	// use the fiber view engine for rendering engine
	app := fiber.New(fiber.Config{
		Views:        engine,
		ErrorHandler: HttpServerErrorHandler,
	})
	// use recover for error handling
	app.Use(recover.New())

	//-----------------------------------------------
	// Middleware
	app.Use(func(c *fiber.Ctx) error {
		// Set some security headers:
		c.Set("X-XSS-Protection", "1; mode=block")
		c.Set("X-Content-Type-Options", "nosniff")
		c.Set("X-Download-Options", "noopen")
		c.Set("Strict-Transport-Security", "max-age=5184000")
		c.Set("X-Frame-Options", "SAMEORIGIN")
		c.Set("X-DNS-Prefetch-Control", "off")

		// Go to next middleware:
		return c.Next()
	})
	//-----------------------------------------------

	app.Get("/", func(c *fiber.Ctx) error {
		//TODO cookie parsing, session storing, etc
		return c.Render("main", fiber.Map{
			"Title": "go-server",
		})
	})

	api := app.Group("/api")                          // /api
	v1 := api.Group("/v1", func(c *fiber.Ctx) error { // middleware for /api/v1
		c.Set("Version", "v1")
		return c.Next()
	})
	v2 := api.Group("/v2", func(c *fiber.Ctx) error { // middleware for /api/v2
		c.Set("Version", "v2")
		return c.Next()
	})

	//-------------------------------------------------
	// routing

	// routing for API v1
	routers.AddRoutersForBuildV1Api(v1, BuildDocker) /* "/build" api */
	routers.AddRoutersForRedirect(v1)                /* "/redirect" api */

	// routing for API v2
	routers.AddRoutersForBuildV2Api(v2, BuildKubernetes) /* "/build" api */

	// define redirect rules
	routers.DefineRedirectRules(app)
	//-------------------------------------------------
	log.Printf("Server is listening %d", *port)
	log.Fatal(app.Listen(addr))
}

func HttpServerErrorHandler(ctx *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError

	// Retrieve the custom status code if it's an fiber.*Error
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}
	// print error log for debugging
	log.Printf("Error code: %d - err.Error(): %v\n", code, err.Error())

	// Send custom error page
	err = ctx.Status(code).SendFile(fmt.Sprintf("views/%d.html", code))
	if err != nil {
		// In case the SendFile fails
		return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	// Return from handler
	return nil
}
