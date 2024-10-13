package api

import (
	"github.com/gofiber/fiber/v2"
	"test/internal/handlers"
)

func Api(app *fiber.App) {
	// Получение списка песен с фильтрацией и пагинацией
	app.Get("/songs", handlers.GetSongs)

	// Получение текста песни с пагинацией по куплетам
	app.Get("/song/:id/lyrics", handlers.GetSongLyrics)

	// Добавление новой песни
	app.Post("/song/add", handlers.AddSong)

	// Удаление песни
	app.Delete("/song/:id/delete", handlers.DeleteSong)

	// Изменение данных песни
	app.Put("/song/:id/update", handlers.UpdateSong)
}
