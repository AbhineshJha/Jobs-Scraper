package handlers

import (
	"net/http"
	"scrapper/services"

	"github.com/gin-gonic/gin"
)

// GetPostingsHandler handles the /get-postings route.
func AtlassianHandler(c *gin.Context) {
	// Call fetchPostings to retrieve the data.
	postings, err := services.AtlassianScrapper()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Return the list of Posting structs as a JSON response.
	c.JSON(http.StatusOK, postings)
}
