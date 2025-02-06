package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"test-project/internal/app"
	"test-project/internal/config"
	"test-project/internal/utils"
	"time"

	flog "github.com/gofiber/fiber/v2/log"
)

// @title			Test Project
// @version		1.0
// @description	This is a test project
// @termsOfService	http://swagger.io/terms/
//
// @host			localhost:8080
// @BasePath		/api/v1
func main() {
	pathToConfig := flag.String("config", "./config/dev.yaml", "path go config file")
	flag.Parse()

	config, err := config.Load(*pathToConfig)
	if err != nil {
		log.Fatalln(err.Error())
	}

	flog.SetLevel(utils.GetFlogLevel(config.LogLevel))

	flog.Info(utils.GetFlogLevel(config.LogLevel))

	flog.Info(*config)

	TestProject, err := app.NewApp(*config)
	if err != nil {
		flog.Fatal(err.Error())
	}

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go TestProject.Run()

	_ = <-c
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	TestProject.Stop(ctx)
}
