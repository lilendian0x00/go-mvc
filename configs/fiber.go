package configs

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

func FiberNewConfig() *fiber.App {
	app := fiber.New(fiber.Config{
		Prefork:                      false,
		ServerHeader:                 "",
		StrictRouting:                false,
		CaseSensitive:                false,
		Immutable:                    false,
		UnescapePath:                 false,
		ETag:                         false,
		BodyLimit:                    10 * 1024 * 1024,
		Concurrency:                  256 * 1024,
		ViewsLayout:                  "",
		ReadTimeout:                  time.Second * 4,
		WriteTimeout:                 time.Second * 4,
		IdleTimeout:                  time.Second * 4,
		ReadBufferSize:               4096,
		WriteBufferSize:              4096,
		CompressedFileSuffix:         ".gz",
		ProxyHeader:                  "",
		GETOnly:                      false,
		ErrorHandler:                 nil,
		DisableKeepalive:             false,
		DisableDefaultDate:           true,
		DisableDefaultContentType:    false,
		DisableHeaderNormalizing:     false,
		DisableStartupMessage:        true,
		AppName:                      "",
		StreamRequestBody:            false,
		DisablePreParseMultipartForm: false,
		ReduceMemoryUsage:            false,
		JSONEncoder:                  nil,
		JSONDecoder:                  nil,
		Network:                      "tcp4",
		EnableTrustedProxyCheck:      false,
		TrustedProxies:               nil,
		EnablePrintRoutes:            false,
	})

	return app
}