package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func HandleIndex(ctx *gin.Context) {
	data := BasicInfo{
		Email:            "dremkay71@gmail.com",
		Current_datetime: time.Now().UTC().Format(time.RFC3339),
		Github_url:       "https://github.com/imoleBytes/hng12_stage_0",
	}

	ctx.JSON(http.StatusOK, data)
}

// struct for basic info
type BasicInfo struct {
	Email            string `json:"email"`
	Current_datetime string `json:"current_datetime"`
	Github_url       string `json:"github_url"`
}
