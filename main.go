package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rdooley/dogs/dogs"
	"net/http"
)

type DogReq struct {
	ID int `uri:"id" binding:"required"`
}

type NewDogReq struct {
	Name    string `json:"name" binding:"required"`
	Owner   string `json:"owner" binding:"required"`
	Details string `json:"details" binding:"required"`
}

func main() {
	router := gin.Default()

	// Get all dogs
	router.GET("/dogs", func(c *gin.Context) {
		c.JSON(http.StatusOK, dogs.LoadDogs())
	})

	// Create a dog
	router.POST("/dogs", func(c *gin.Context) {
		var req NewDogReq
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		dog, err := dogs.NewDog(req.Name, req.Owner, req.Details)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"ID": dog.ID})
	})

	// Get a specific dog
	router.GET("/dogs/:id", func(c *gin.Context) {
		var req DogReq
		if err := c.ShouldBindUri(&req); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}
		dog, found := dogs.LoadDog(req.ID)
		if !found {
			c.String(http.StatusNotFound, fmt.Sprintf("Dog %d not found", req))
		} else {
			c.JSON(http.StatusOK, dog)
		}
	})

	// Update a specific dog
	router.PUT("/dogs/:id", func(c *gin.Context) {
		c.String(http.StatusOK, "Dogs")
		// ID := c.Param("id")
	})

	// Delete a dog
	router.DELETE("/dogs/:id", func(c *gin.Context) {
		c.String(http.StatusOK, "Dogs")
		// ID := c.Param("id")
	})

	// Run the thing on 8080
	router.Run(":8080")
}
