// @title Songs API
// @version 1.0
// @description API for managing songs.
// @host localhost:8080
// @BasePath /

package main

import (
	_ "github.com/Hell077/Test_Task/docs"
	"github.com/Hell077/Test_Task/internal"
	"github.com/Hell077/Test_Task/internal/api"
	"github.com/Hell077/Test_Task/internal/config"
	"github.com/gofiber/fiber/v2"
	"sync"
)

func main() {
	{
		err := config.LoadEnv()
		if err != nil {
			return
		}
	}
	res := internal.CheckExist()
	if res == false {
		wg := sync.WaitGroup{}
		wg.Add(1)
		func() {
			defer wg.Done()
			err := internal.RunMigrations()
			if err != nil {
				panic(err)
			}
		}()
		wg.Wait()
	}
	app := fiber.New()
	api.Api(app)
	{
		err := app.Listen(":8000")
		if err != nil {
			return
		}
	}
}
