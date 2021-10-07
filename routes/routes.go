package routes

import (
	controllers "beelogiq/notes/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/", welcome)
	router.GET("/note", controllers.GetAllNotes)
	router.POST("/note", controllers.CreateNote)
	router.GET("/note/:note_id", controllers.GetSingleNote)
	router.PUT("/note/:note_id", controllers.EditNote)
	router.DELETE("/note/:note_id", controllers.DeleteNote)
	router.NoRoute(notFound)
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome To API",
	})
	return
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  404,
		"message": "Route Not Found",
	})
	return
}
