package controllers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func ResponseError(c *gin.Context, e error) {
	log.Error().Msgf("erro when handle [%s]: %w", c.Request.URL.Path, e)

	msg := "internal error"
	if os.Getenv("IS_DEV") == "true" {
		msg = e.Error()
	}

	c.JSON(http.StatusInternalServerError, gin.H{"message": msg})
}
