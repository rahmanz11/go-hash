package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rahmanz11/go-hash/src/middleware"
	"github.com/rahmanz11/go-hash/src/payloads"
)

// Todo struct for request body
type HashRequest struct {
	Data      payloads.Data `json:"data"`
	Algorithm string        `json:"algorithm"`
}

// Defining struct for response
type HashResponse struct {
	Hash string `json:"hash"`
}

// Create todo data to database by run this function
func CreateHash(context *gin.Context) {
	var req HashRequest

	// Binding request body json to request body struct
	if err := context.ShouldBindJSON(&req.Data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, _ := json.Marshal(req)
	fmt.Println(string(res))
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
