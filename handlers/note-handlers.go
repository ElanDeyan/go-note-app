package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	models "elandeyan.github.io/go-note-app-models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateNote(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var noteEntry models.NoteCreateEntry
		if err := ctx.ShouldBindJSON(&noteEntry); err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}

		note := models.Note{
			Title:    noteEntry.Title,
			Content:  noteEntry.Content,
			IsPublic: noteEntry.IsPublic,
		}

		result := db.Create(&note)

		if result.Error != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"msg": result.Error})
			return
		}
		message := fmt.Sprintf("created %v entries", result.RowsAffected)
		ctx.IndentedJSON(http.StatusAccepted, gin.H{"msg": message})
	}
}

func ListNotes(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var notes []models.Note
		result := db.Find(&notes)

		if result.Error != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"msg": result.Error})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"notes": notes})
	}
}

func GetNoteByID(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		noteId, err := strconv.Atoi(idParam)

		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}
		var note models.Note
		result := db.First(&note, noteId)

		switch {
		case errors.Is(result.Error, gorm.ErrRecordNotFound):
			ctx.IndentedJSON(http.StatusNotFound, gin.H{"msg": "note not found"})
			return
		case result.Error != nil:
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"msg": result.Error.Error()})
			return
		}

		ctx.IndentedJSON(http.StatusOK, gin.H{"note": note})
	}
}

func UpdateNoteById(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		noteId, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}
		var updateNoteEntry models.NoteUpdateEntry
		if err := ctx.ShouldBindJSON(&updateNoteEntry); err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}

		var matchedNote models.Note
		result := db.First(&matchedNote, noteId)
		if result.Error != nil {
			ctx.IndentedJSON(http.StatusNotFound, gin.H{"msg": result.Error})
			return
		}

		if updateNoteEntry.Title != nil {
			matchedNote.Title = *updateNoteEntry.Title
		}
		if updateNoteEntry.Content != nil {
			matchedNote.Content = *updateNoteEntry.Content
		}
		if updateNoteEntry.IsPublic != nil {
			matchedNote.IsPublic = updateNoteEntry.IsPublic
		}

		db.Save(&matchedNote)
		ctx.IndentedJSON(http.StatusAccepted, gin.H{"msg": "successful note update", "note": matchedNote})
	}
}

func DeleteNoteById(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		noteId, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}

		result := db.Delete(models.Note{}, noteId)
		if result.Error != nil {
			ctx.IndentedJSON(http.StatusNotFound, gin.H{"msg": result.Error})
			return
		}

		ctx.IndentedJSON(http.StatusOK, gin.H{"msg": "successfully deleted"})
	}
}
