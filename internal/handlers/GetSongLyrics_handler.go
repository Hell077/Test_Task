package handlers

import (
	"github.com/Hell077/Test_Task/internal/database"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
)

// GetSongLyrics godoc
// @Summary Get song lyrics by ID
// @Description Get lyrics of the song by ID
// @Produce json
// @Param id path int true "Song ID"
// @Param verse query int false "Verse number"
// @Success 200 {string} string "Song Lyrics"
// @Failure 404 {object} string "Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /songs/{id}/lyrics [get]
func GetSongLyrics(c *fiber.Ctx) error {
	database.Connect()
	defer database.Close()
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid song ID"})
	}

	db := database.DB

	rows, err := db.Query(`SELECT verse_number, text FROM song_lyrics WHERE song_id = $1 ORDER BY verse_number`, id)
	if err != nil {
		log.Printf("Error fetching song lyrics: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch song lyrics"})
	}
	defer rows.Close()

	var lyrics []map[string]interface{}

	for rows.Next() {
		var verseNumber int
		var lyricText string
		if err := rows.Scan(&verseNumber, &lyricText); err != nil {
			log.Printf("Error scanning lyrics: %v", err)
			return c.Status(500).JSON(fiber.Map{"error": "Failed to read lyrics"})
		}
		lyric := map[string]interface{}{
			"verse_number": verseNumber,
			"text":         lyricText,
		}
		lyrics = append(lyrics, lyric)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error after reading lyrics: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": "Error while processing lyrics"})
	}

	return c.JSON(fiber.Map{"lyrics": lyrics})
}
