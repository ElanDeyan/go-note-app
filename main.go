package main

import (
	"log"

	handlers "elandeyan.github.io/go-note-app-handlers"
	models "elandeyan.github.io/go-note-app-models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	log.Print("Starting to connect with database")

	db, err := gorm.Open(sqlite.Open("sample.db"), &gorm.Config{})
	if err != nil {
		log.Panic("unable to connect with database")
	}
	log.Print("Database connected")

	log.Print("Making auto migrate")
	db.AutoMigrate(&models.Note{})

	log.Print("Initializing gin")
	router := gin.Default()

	router.GET("/notes", handlers.ListNotes(db))
	router.GET("/notes/:id", handlers.GetNoteByID(db))
	router.POST("/notes", handlers.CreateNote(db))
	router.PATCH("/notes/:id", handlers.UpdateNoteById(db))
	router.DELETE("/notes/:id", handlers.DeleteNoteById(db))

	router.Run()

}
