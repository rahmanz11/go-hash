package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rahmanz11/go-hash/src/middleware"
	"github.com/rahmanz11/go-hash/src/payloads"
)

// Hash request struct for request body
type HashRequest struct {
	Data      payloads.Data `json:"data"`
	Algorithm string        `json:"algorithm"`
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

	// Matching data payloads struct with data request struct
	data := payloads.Data{}
	data.Value = req.Data.Value

	// Create the hash
	result := middleware.CreateHash(&data, req.Algorithm)
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
