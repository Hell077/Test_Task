package api

import (
	"github.com/Hell077/Test_Task/internal/handlers"
	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func Api(app *fiber.App) {
	// Получение списка песен с фильтрацией и пагинацией
	app.Get("/songs/:id", handlers.GetSongs)

	// Получение текста песни с пагинацией по куплетам
	app.Get("/songs/:id/lyrics", handlers.GetSongLyrics)

	// Добавление новой песни
	app.Post("/song/add", handlers.AddSong)

	// Удаление песни
	app.Delete("/song/:id/delete", handlers.DeleteSong)

	// Изменение данных песни
	app.Put("/songs/:id/update", handlers.UpdateSong)

	app.Get("/swagger/*", fiberSwagger.WrapHandler)
}
