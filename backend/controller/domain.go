package controller

import (
	"5ub5i5t/bughuntingtoolbox/database"
	"5ub5i5t/bughuntingtoolbox/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
}

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

func UpdateDomainById(c *gin.Context) {
	fmt.Printf("UpdateDomainById here...")
	var updateDomain model.Domain
	id := c.Param("id")

	if err := c.ShouldBindJSON(&updateDomain); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.Database.Model(&updateDomain).Where("id = ?", id).Updates(updateDomain)
	if result.RowsAffected == 0 {
		//return model.Domain{}, errors.New("Domain data not update.")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Domain data not updated."})
		return
	}
	//return updateDomain, nil
	c.JSON(http.StatusOK, nil)
}

func DeleteDomainById(c *gin.Context) {
	//var domain model.Domain
	id := c.Param("id")

	result := database.Database.Delete(&model.Domain{}, id)

	if result.Error != nil {
		fmt.Printf("Delete error...")
		//c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No record found with that id."})
		c.JSON(http.StatusOK, gin.H{"status": "fail", "message": "No record found with that id."})
		return
	}

	fmt.Printf("Delete success!")
	c.JSON(http.StatusNoContent, nil)
	//c.JSON(http.StatusOK, nil)
}
