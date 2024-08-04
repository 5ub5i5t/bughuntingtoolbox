package controller

import (
	"5ub5i5t/bughuntingtoolbox/database"
	"5ub5i5t/bughuntingtoolbox/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddDomain(context *gin.Context) {
	context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	context.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	context.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	context.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

	var input model.Domain
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	savedEntry, err := input.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": savedEntry})
}

func GetDomainById(c *gin.Context) {
	var domain model.Domain
	id := c.Param("id")

	result := database.Database.First(&domain, id)
	if result.Error != nil {
		fmt.Println(result.Error)
	} else {
		fmt.Printf("ID: %d, Target: %s, Domain: %s\n", domain.ID, domain.Target, domain.Domain)
	}

	c.JSON(http.StatusOK, domain)

	//c.JSON(200, gin.H{
	//	"Item ID": id,
	//})
	//return domain, nil
}

func GetDomains(c *gin.Context) {
	var domains []model.Domain

	result := database.Database.Find(&domains)
	if result.Error != nil {
		fmt.Println(result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, domains)
}
