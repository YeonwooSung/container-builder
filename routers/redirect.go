package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/redirect/v2"
)

func DefineRedirectRules(app *fiber.App) {
	// 나중에 v1에서 v2로 넘어가면 "/v1/*": "/v2/*"
	app.Use(redirect.New(redirect.Config{
		Rules: map[string]string{
			"/redirects": "/",
		},
		StatusCode: 301,
	}))
}

func handler_for_redirect_api(c *fiber.Ctx) error {
	//TODO cookie parsing, session storing, etc
	return c.Redirect("/")
}

func AddRoutersForRedirect(router fiber.Router) {
	router.Get("/redirect", handler_for_redirect_api)
}
