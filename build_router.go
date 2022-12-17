package main

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
)

func generate_build_spec() *BuildSpec {
	buildSpec := BuildSpec{}

	//TODO
	// if err := json.Unmarshal(body, &buildSpec); err != nil {
	// 	log.Printf("json unmarshal has failed: %s\n", err)
	// 	return nil
	// }

	if buildSpec.BuildArgStr != "" {
		if err := json.Unmarshal([]byte(buildSpec.BuildArgStr), &buildSpec.BuildArg); err != nil {
			log.Printf("docker argument unmarshal has failed : %s\n", err)
			log.Println(buildSpec.BuildArgStr)
			return nil
		}
	}

	log.Printf("buildSpec: %#+v\n", buildSpec)

	return &buildSpec
}

func get_handler_for_build_v1_api(c *fiber.Ctx) error {
	// call BuildDocker
	go BuildDocker(generate_build_spec())

	//TODO cookie parsing, session storing, etc
	return c.Render("main", fiber.Map{
		"Title": "dashpad",
	})
}

func post_handler_for_build_v1_api(c *fiber.Ctx) error {
	// call BuildDocker
	go BuildDocker(generate_build_spec())

	//TODO cookie parsing, session storing, etc
	return c.Render("main", fiber.Map{
		"Title": "dashpad",
	})
}

func handler_for_build_v2_api(c *fiber.Ctx) error {
	// call BuildKubernetes
	go BuildKubernetes()

	//TODO cookie parsing, session storing, etc
	return c.Render("main", fiber.Map{
		"Title": "dashpad",
	})
}

func AddRoutersForBuildV1Api(router fiber.Router) {
	router.Get("/build", get_handler_for_build_v1_api)
	router.Post("/build", post_handler_for_build_v1_api)
}

func AddRoutersForBuildV2Api(router fiber.Router) {
	router.Post("/build", handler_for_build_v2_api)
}
