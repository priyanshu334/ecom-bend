package app

import (
	"github.com/gofiber/fiber/v3"
	"github.com/priyanshu334/tw-bend/internal/config"
	"github.com/priyanshu334/tw-bend/internal/db"
	"github.com/priyanshu334/tw-bend/internal/logger"
	"github.com/priyanshu334/tw-bend/internal/router"
)

func Start() error {
	config.Load()
	logger.Init(config.Cfg.AppEnv)
	if err := db.Connect(config.Cfg.Database); err != nil {
		return err
	}
	app := fiber.New()

	router.Setup(app)
	logger.Log.Info("server started")
	return app.Listen(":" + config.Cfg.AppPort)

}
