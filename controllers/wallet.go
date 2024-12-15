package controllers

import (
	"go-simple-MVC/structs"
	"net/http"

	"go-simple-MVC/helpers"

	"github.com/gin-gonic/gin"
)

type WalletDTO struct {
	ID              uint    `json:"id"`
	Name            string  `json:"name"`
	Type            string  `json:"type"`
	Balance         float64 `json:"balance"`
	Key_Phrase      string  `json:"key_phrase"`
	User_ID         uint    `json:"user_id"`
	Virtual_Account string  `json:"virtual_account"`
	Tag_Name        string  `json:"tag_name"`
}

func (idb *InDB) GetAllWallet(c *gin.Context) {

	// limit := c.DefaultQuery("limit", "5")
	// page := c.DefaultQuery("page", "0")
	// cari := c.DefaultQuery("cari", "")
	
	var (
		result gin.H
	)

	var wallet []structs.Wallets

	idb.DB.Unscoped().Find(&wallet)
	if len(wallet) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} 

	result = helpers.SuccessResponse(wallet, "Success Get All Wallet")

	// result = gin.H{
	// 	"result": wallet,
	// 	"count":  len(wallet),
	// }



	c.JSON(http.StatusOK, result)
}

func (idb *InDB) GetDetailWallet(c *gin.Context) {
	var (
		person []structs.Person
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

func (idb *InDB) CreateWallet(c *gin.Context) {
	var (
		person structs.Person
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

func (idb *InDB) UpdateWallet(c *gin.Context) {
		id := c.Query("id")
		first_name := c.PostForm("first_name")
		last_name := c.PostForm("last_name")

		var (
			person    structs.Person
			newPerson structs.Person
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

func (idb *InDB) DeleteWallet(c *gin.Context) {
	var (
		person structs.Person
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