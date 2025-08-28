package routes

import (
	"encoding/json"
	"net/http"

	"github.com/AuraReaper/go-url-shortner/api/database"
	"github.com/gin-gonic/gin"
)

type TagRequest struct {
	ShortID string `json:"shortID"`
	Tag string `json:"tag"`
}

func AddTag(c *gin.Context) {
	var tagRequest TagRequest

	if err := c.ShouldBind(&tagRequest); err != nil {
		c.JSON(http.StatusBadRequest , gin.H{
			"error" : "Inavlid Request Body",
		})
		return
	}

	shortID := tagRequest.ShortID
	tag := tagRequest.Tag

	r := database.CreateClient(0)
	defer r.Close()

	val, err := r.Get(database.Ctx, shortID).Result()

	if err != nil {
		c.JSON(http.StatusNotFound , gin.H {
			"error" : "Data not found for the provided ShortID",
		})
		return
	}

	var data map[string]interface{}
	if err := json.Unmarshal([]byte(val), &data); err != nil  {
		// id the data is not a jkson object assume it a string
		data = make(map[string]interface{})
		data["data"] = val
	}

	// check if tag field already strings and is it a slice of strings
	var tags []string
	if exsistingTags, ok := data["tags"].([]interface{}); ok {
		for _, t := range exsistingTags {
			if strTag, ok := t.(string); ok {
				tags = append(tags, strTag)
			}
		}
	}

	//check for duplicate tags
	for _, exsistingTags := range tags {
		if exsistingTags == tag { 
			c.JSON(http.StatusBadRequest, gin.H{
				"error" : "Tag already exsists",
			})
			return
		}
	}

	// add new tag to tag slice
	tags = append(tags, tag)
	data["tags"] = tags

	//marshal the data
	updateData, err := json.Marshal(data)

	if err != nil {
		c.JSON(http.StatusInternalServerError , gin.H{
			"error" : "Failed to Marshal updated data",
		})
		return
	}

	err = r.Set(database.Ctx, shortID, updateData, 0).Err()

	if err != nil {
		c.JSON(http.StatusInternalServerError , gin.H{
			"error" : "failed to update the database",
		})
		return
	}

	c.JSON(http.StatusOK , data)
}