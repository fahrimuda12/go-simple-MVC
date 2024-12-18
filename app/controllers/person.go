package controllers

import (
	"go-simple-MVC/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


func (idb *InDB) GetAllPerson(c *gin.Context) {
	var (
		person []models.Person
		result gin.H
	)

	idb.DB.Find(&person)
	if len(person) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": person,
			"count":  len(person),
		}
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) GetDetailPerson(c *gin.Context) {
	var (
		person []models.Person
		result gin.H
	)

	id := c.Param("id")
	err := idb.DB.First(&person, id).Error
	if err != nil {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": person,
			"count":  1,
		}
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) CreatePerson(c *gin.Context) {
	var (
		person models.Person
		result gin.H
	)

	first_name := c.PostForm("first_name")
	last_name := c.PostForm("last_name")
	person.First_Name = first_name
	person.Last_Name = last_name
	idb.DB.Create(&person)
	result = gin.H{
		"result": person,
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) UpdatePerson(c *gin.Context) {
		id := c.Query("id")
		first_name := c.PostForm("first_name")
		last_name := c.PostForm("last_name")

		var (
			person    models.Person
			newPerson models.Person
			result    gin.H
		)

		err := idb.DB.First(&person, id).Error
		if err != nil {
			result = gin.H{
				"result": "Data not found",
			}
		}

		newPerson.First_Name = first_name
		newPerson.Last_Name = last_name
		err = idb.DB.Model(&person).Updates(newPerson).Error
		if err != nil {
			result = gin.H{
				"result": "Update failed",
			}
		} else {
			result = gin.H{
				"result": "Data updated successfully",
			}
		}
		c.JSON(http.StatusOK, result)
}

func (idb *InDB) DeletePerson(c *gin.Context) {
	var (
		person models.Person
		result gin.H
	)

	id := c.Param("id")
	err := idb.DB.First(&person, id).Error
	if err != nil {
		result = gin.H{
			"result": "Data not found",
		}
	}

	err = idb.DB.Delete(&person, id).Error
	if err != nil {
		result = gin.H{
			"result": "Delete failed",
		}
	} else {
		result = gin.H{
			"result": "Data deleted successfully",
		}
	}

	c.JSON(http.StatusOK, result)
}