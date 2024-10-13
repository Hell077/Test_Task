package handlers

import (
	"database/sql"
	"github.com/Hell077/Test_Task/internal/database"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
)

// GetSongs godoc
// @Summary Get a song by ID
// @Description Get song details by ID
// @Produce json
// @Param id path int true "Song ID"
// @Success 200 {object} SongDetail
// @Failure 404 {object} string "Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /songs/{id} [get]

func GetSongs(c *fiber.Ctx) error {
	database.Connect()
	defer database.Close()
	// Извлекаем параметр id из URL
	idParam := c.Params("id")

	// Преобразуем строку в целое число
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid song ID"})
	}

	db := database.DB // используем существующее соединение

	// Выполняем SQL-запрос для получения песни по ID
	row := db.QueryRow("SELECT id, group_name, song_name, release_date, text FROM songs WHERE id = $1", id)

	var groupName, songName string
	var releaseDate sql.NullString
	var text sql.NullString

	err = row.Scan(&id, &groupName, &songName, &releaseDate, &text)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(404).JSON(fiber.Map{"error": "Song not found"})
		}
		log.Printf("Error fetching song data: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch song data"})
	}

	// Создаем ответ с учётом возможных NULL значений
	response := fiber.Map{
		"id":           id,
		"group_name":   groupName,
		"song_name":    songName,
		"release_date": releaseDate.String, // Если это NULL, просто вернется пустая строка
		"text":         text.String,        // Если это NULL, просто вернется пустая строка
	}

	// Устанавливаем правильные значения для release_date и text
	if !releaseDate.Valid {
		response["release_date"] = nil
	}
	if !text.Valid {
		response["text"] = nil
	}

	return c.JSON(response)
}
