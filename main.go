package main

import (
	"github.com/gofiber/fiber/v2"
	"kredit-plus/apply"
	"kredit-plus/internal/config"
	container "kredit-plus/src/shared/di"
)

func main() {
	app := fiber.New(
		fiber.Config{
			AppName: config.GetConfig().Server.Rest.Host,
		})
	//app.Use(
	//	recover.New(),
	//	logger.New(),
	//	cors.New(),
	//)
	v2 := app.Group("/v1")
	ctn := container.NewContainer()

	apply.Apply(v2, ctn)
	err := app.Listen(":" + string(config.GetConfig().Server.Rest.Port))
	if err != nil {
		panic(err)
	}

}
