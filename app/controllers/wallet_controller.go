package controllers

import (
	"go-simple-MVC/app/models"
	"net/http"
	"strconv"

	"go-simple-MVC/app/helpers"

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

	pagePar := c.DefaultQuery("page", "1")
	// convert page to int
	page, err := strconv.Atoi(pagePar)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ErrorResponse(err.Error(), "400"))
		return
	}
	limitPar := c.DefaultQuery("limit", "10")
	// convert limit to int
	limit, err := strconv.Atoi(limitPar)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ErrorResponse(err.Error(), "400"))
		return
	}
	
	var (
		result gin.H
		wallet []models.Wallets
		count int64
	)

	print(page)
	print(limit)
	offset := (page - 1) * limit
	idb.DB.Unscoped().Offset(offset).Limit(limit).Find(&wallet)

	idb.DB.Model(&wallet).Count(&count)

	data := map[string]interface{}{
		"data": wallet,
		"count": count,
		"page": page,
		"limit": limit,
	}

	result = helpers.SuccessResponse(data, "Success Get All Wallet")

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) GetDetailWallet(c *gin.Context) {
	var (
		wallet []models.Wallets
		result gin.H
	)

	id := c.Param("id")
	walletID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ErrorResponse("Invalid ID format", "400"))
		return
	}
	err = idb.DB.Unscoped().First(&wallet, walletID).Error
	if err != nil {
		c.JSON(http.StatusNotFound, helpers.ErrorResponse("Data not found", "404"))
		return
	}

	result = helpers.SuccessResponse(wallet, "Success Get Detail Wallet")
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) CreateWallet(c *gin.Context) {
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

func (idb *InDB) UpdateWallet(c *gin.Context) {
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

func (idb *InDB) DeleteWallet(c *gin.Context) {
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