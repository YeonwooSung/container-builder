package routers

import (
	"github.com/gofiber/fiber/v2"
)

func handler_for_build_api(c *fiber.Ctx) error {
	//TODO cookie parsing, session storing, etc
	return c.Render("main", fiber.Map{
		"Title": "dashpad",
	})
}

func AddRoutersForBuildApi(router fiber.Router) {
	router.Get("/build", handler_for_build_api)
}
