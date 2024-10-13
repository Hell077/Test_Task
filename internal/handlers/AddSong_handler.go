package handlers

import (
	"fmt"
	"github.com/Hell077/Test_Task/internal/database"
	"github.com/gofiber/fiber/v2"
)

type Song struct {
	Group string `json:"group"`
	Song  string `json:"song"`
}

// AddSong godoc
// @Summary Add a new song
// @Description Add a new song to the library
// @Accept json
// @Produce json
// @Param song body Song true "Song info"
// @Success 201 {object} string "success"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /songs/add [post]
func AddSong(c *fiber.Ctx) error {
	var newSong Song
	defer database.Close()
	if err := c.BodyParser(&newSong); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	errorChan := make(chan error)
	go func() {
		database.Connect()
		_, err := database.DB.Exec("INSERT INTO songs (group_name, song_name) VALUES ($1, $2)", newSong.Group, newSong.Song)
		if err != nil {
			errorChan <- err
			return
		}
		fmt.Printf("successfully adding a song %s,%s \n", newSong.Song, newSong.Group)
		errorChan <- nil
	}()
	err := <-errorChan
	if err != nil {
		return err
	}
	return nil
}
