package handlers

import (
	"github.com/Hell077/Test_Task/internal/database"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

// DeleteSong godoc
// @Summary Delete a song by ID
// @Description Delete a song from the library by ID
// @Param id path int true "Song ID"
// @Success 204 {string} string "Deleted successfully"
// @Failure 404 {object} string "Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /songs/{id} [delete]
func DeleteSong(c *fiber.Ctx) error {
	database.Connect()
	defer database.Close()
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid song ID")
	}

	var exists bool
	checkQuery := `SELECT EXISTS(SELECT 1 FROM songs WHERE id = $1)`
	err = database.DB.QueryRow(checkQuery, id).Scan(&exists)
	if err != nil || !exists {
		return c.Status(fiber.StatusNotFound).SendString("Song not found")
	}

	deleteLyricsQuery := `
        DELETE FROM song_lyrics
        WHERE song_id = $1
    `
	_, err = database.DB.Exec(deleteLyricsQuery, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error deleting song lyrics")
	}
	deleteSongQuery := `
        DELETE FROM songs
        WHERE id = $1
    `
	_, err = database.DB.Exec(deleteSongQuery, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error deleting song")
	}

	return c.Status(fiber.StatusNoContent).SendString("Song deleted successfully")
}
