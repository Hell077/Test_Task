package handlers

import (
	"github.com/Hell077/Test_Task/internal/database"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

// UpdateSong godoc
// @Summary Update song details
// @Description Update the details of a song and its lyrics by ID
// @Accept json
// @Produce json
// @Param id path int true "Song ID"
// @Param song body Song1 true "Song info"
// @Success 200 {string} string "Updated successfully"
// @Failure 404 {object} string "Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /songs/{id} [put]

type Song1 struct {
	GroupName string       `json:"group_name"`
	SongName  string       `json:"song_name"`
	Text      string       `json:"text"`
	Lyrics    []VerseInput `json:"lyrics"`
}

type VerseInput struct {
	VerseNumber int    `json:"verse_number"`
	Text        string `json:"text"`
}

func UpdateSong(c *fiber.Ctx) error {
	database.Connect()
	defer database.Close()
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid song ID")
	}

	var song Song1
	if err := c.BodyParser(&song); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request payload")
	}

	query := `
        UPDATE songs
        SET group_name = $1, song_name = $2, text = $3
        WHERE id = $4
    `
	_, err = database.DB.Exec(query, song.GroupName, song.SongName, song.Text, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error updating song")
	}

	for _, verse := range song.Lyrics {
		query = `
            UPDATE song_lyrics
            SET text = $1
            WHERE song_id = $2 AND verse_number = $3
        `
		_, err = database.DB.Exec(query, verse.Text, id, verse.VerseNumber)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error updating song lyrics")
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Song updated successfully",
	})
}
