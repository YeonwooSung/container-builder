package routers

import (
	"github.com/gofiber/fiber/v2"
)

type fn func()

var build_v1 fn = nil
var build_v2 fn = nil

func get_handler_for_build_v1_api(c *fiber.Ctx) error {
	// check if build_v1 is set
	if build_v1 == nil {
		println("build_v1 is not set")
		return c.SendStatus(fiber.StatusNotFound)
	}

	// call build_v1
	go build_v1()

	//TODO cookie parsing, session storing, etc
	return c.Render("main", fiber.Map{
		"Title": "dashpad",
	})
}

func post_handler_for_build_v1_api(c *fiber.Ctx) error {
	// check if build_v1 is set
	if build_v1 == nil {
		println("build_v1 is not set")
		return c.SendStatus(fiber.StatusNotFound)
	}

	// call build_v1
	go build_v1()

	//TODO cookie parsing, session storing, etc
	return c.Render("main", fiber.Map{
		"Title": "dashpad",
	})
}

func handler_for_build_v2_api(c *fiber.Ctx) error {
	// check if build_v1 is set
	if build_v2 == nil {
		println("build_v1 is not set")
		return c.SendStatus(fiber.StatusNotFound)
	}

	// call build_v1
	go build_v2()

	//TODO cookie parsing, session storing, etc
	return c.Render("main", fiber.Map{
		"Title": "dashpad",
	})
}

func AddRoutersForBuildV1Api(router fiber.Router, routine_func fn) {
	build_v1 = routine_func
	router.Get("/build", get_handler_for_build_v1_api)
	router.Post("/build", post_handler_for_build_v1_api)
}

func AddRoutersForBuildV2Api(router fiber.Router, routine_func fn) {
	build_v2 = routine_func
	router.Post("/build", handler_for_build_v2_api)
}
