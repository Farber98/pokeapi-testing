package controller

import (
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
)

// respondwithJSON writes json response format.
func respondwithJSON(c *gin.Context, code int, payload interface{}) {
	// Sets headers to response.
	c.Header("Content-Type", "application/json")
	c.Writer.WriteHeader(code)

	// Composes JSON response.
	if err := json.NewEncoder(c.Writer).Encode(payload); err != nil {
		log.Fatal(err)
	}
}
