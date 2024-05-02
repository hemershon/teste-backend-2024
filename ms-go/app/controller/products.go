package controller

import (
	"log"
	"ms-go/app/helpers"
	"ms-go/app/models"
	"ms-go/app/services/products"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func handleError(c *gin.Context, err error) {
	switch e := err.(type) {
	case *helpers.GenericError:
		c.JSON(e.Code, gin.H{"message": e.Error()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
}

func IndexProducts(c *gin.Context) {
	log.Println("Received request to list all products")

	all, err := products.ListAll()
	if err != nil {
		log.Printf("Error listing products: %s", err.Error())
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": all})
}

func ShowProducts(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}

	product, err := products.Details(models.Product{ID: id})
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

func CreateProducts(c *gin.Context) {
	var params models.Product

	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Name and price are required"})
		return
	}

	product, err := products.Create(params, true)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": product})
}

func UpdateProducts(c *gin.Context) {
	var params models.Product

	if err := c.BindJSON(&params); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}

	params.ID = id

	product, err := products.Update(params, true)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}
