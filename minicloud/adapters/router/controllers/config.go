package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"lucares.github.com/minicloud/minicloud/domain/use_cases/config"
)

func WasConfiguredHandler(c *gin.Context) {
	uc, err := config.NewWasConfiguredUseCase(c.Request.Context())
	if err != nil {
		ResponseError(c, err)
		return
	}

	ok, err := uc.Execute(c.Request.Context())
	if err != nil {
		ResponseError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"wasConfigured": ok,
	})
}
