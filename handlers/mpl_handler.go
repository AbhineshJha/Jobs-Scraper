package handlers

import (
	"net/http"
	"scrapper/services"

	"github.com/gin-gonic/gin"
)

func MplHandler(c *gin.Context) {
	// Call fetchPostings to retrieve the data.
	postings, err := services.MplScrapper()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the list of Posting structs as a JSON response.
	c.JSON(http.StatusOK, postings)
}
