package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model `json:"model"`
	Title      string `json:"title"`
	Body       string `json:"body"`
	Url        string `json:"url"`
}

// Create User Table
func CreateNoteTable(db *gorm.DB) error {
	db.AutoMigrate(&Note{})
	log.Printf("Note table created")
	return nil
}

// Initialize DB connection (to avoid too many connections)
var dbConnect *gorm.DB

func InitiateDB(db *gorm.DB) {
	dbConnect = db
}

func GetAllNotes(c *gin.Context) {
	var notes []Note
	result := dbConnect.Find(&notes)

	if result.Error != nil {
		log.Printf("Error while getting all notes, Reason: %v\n", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All Notes",
		"data":    notes,
	})
	return
}

func CreateNote(c *gin.Context) {
	var note Note
	c.BindJSON(&note)
	title := note.Title
	body := note.Body
	url := note.Url

	result := dbConnect.Create(&Note{
		Title: title,
		Body:  body,
		Url:   url,
	})
	if result.Error != nil {
		log.Printf("Error while inserting new note into db, Reason: %v\n", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Note created Successfully",
	})
	return
}

func GetSingleNote(c *gin.Context) {
	var note Note
	note_id := c.Param("note_id")
	result := dbConnect.First(&note, note_id)

	if result.Error != nil {
		log.Printf("Error while getting a single note, Reason: %v\n", result.Error)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Note not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Single Note",
		"data":    note,
	})
	return
}

func EditNote(c *gin.Context) {
	var note Note
	note_id := c.Param("note_id")
	dbConnect.First(&note, note_id)
	c.BindJSON(&note)

	result := dbConnect.Save(&note)
	if result.Error != nil {
		log.Printf("Error, Reason: %v\n", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Note Edited Successfully",
	})
	return
}

func DeleteNote(c *gin.Context) {
	note_id := c.Param("note_id")

	result := dbConnect.Delete(&Note{}, note_id)
	if result.Error != nil {
		log.Printf("Error while deleting a single note, Reason: %v\n", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Note deleted successfully",
	})
	return
}
