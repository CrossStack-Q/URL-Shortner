package routes

import (
	"net/http"
	"time"

	"github.com/AuraReaper/go-url-shortner/api/database"
	"github.com/AuraReaper/go-url-shortner/api/models"
	"github.com/gin-gonic/gin"
)

func EditURL(c *gin.Context) {

	shortID := c.Param("shortID")
	var body models.Request

	if err := c.BindJSON(&body) ; err != nil {
		c.JSON(http.StatusBadRequest , gin.H {
			"error" : "Error in Parsing JSON",
		})
		return
	}

	r := database.CreateClient(0)
	defer r.Close()

	val, err := r.Get(database.Ctx, shortID).Result()

	// check is id exsist in db
	if err != nil  || val == "" {
		c.JSON(http.StatusNotFound , gin.H{
			"error" : "ShortID does not exsist",
		})
	}

	// update the content of the url, expiry time usingf ShortID
	err = r.Set(database.Ctx, shortID, body.URL, body.Expiry*3600*time.Second).Err()

	if err != nil {
		c.JSON(http.StatusInternalServerError , gin.H{
			"error" : "Unable to update the content",
		})
		return
	}

	c.JSON(http.StatusOK , gin.H{
		"message" : "The Content has been updated!!!",
	})
}