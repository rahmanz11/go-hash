package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rahmanz11/go-hash/src/middleware"
)

// Hash request struct for request body
type HashRequest struct {
	Data      []string        `json:"data"`
	Algorithm middleware.Algo `json:"algorithm"`
}

// Defining struct for response
type HashResponse struct {
	Hash string `json:"hash"`
}

// Create hash by run this function
func CreateHash(context *gin.Context) {
	var req HashRequest

	// Binding request body json to request body struct
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create the hash
	result := middleware.CreateHash(req.Data, req.Algorithm)
	if result == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong"})
		return
	}

	// Set result to create response
	var response HashResponse
	response.Hash = result

	// Creating http response
	context.JSON(http.StatusCreated, response)
}
