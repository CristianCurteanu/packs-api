package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/CristiCurteanu/pack-api/internal/common/env"
	"github.com/CristiCurteanu/pack-api/internal/common/packager"
	"github.com/CristiCurteanu/pack-api/internal/common/storage"
	"github.com/CristiCurteanu/pack-api/internal/packsapi"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	apiV1 := app.Group("/api/v1")

	fileArchiver := storage.NewFile[packager.PackSizes](os.Getenv("PACKS_DATA_PATH"))
	packs := fileArchiver.LoadOrDefault(packager.DefaultPackSizes)
	pkgr := packager.NewPackager(packs)

	packsapi.RegisterHandlers(apiV1, pkgr)

	go func() {
		if err := app.Listen(env.GetEnvOrDefault("API_PORT", ":8089")); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	log.Println("Gracefully shutting down...")
	app.Shutdown()

	log.Println("Running cleanup tasks...")

	// Cleanup tasks
	fileArchiver.Save(pkgr.List())

	log.Println("Fiber was successful shutdown.")
}
