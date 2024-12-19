package controllers

import (
	"fmt"
	"go-simple-MVC/app/helpers"
	"go-simple-MVC/app/models"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type WalletCreate struct {
	Name            string  `json:"name" binding:"required"`
	Type            string  `json:"type" binding:"required"`
	Balance         float64 `json:"balance"`
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

	walletID := c.Param("id")
	fmt.Print(walletID)
	
	err := idb.DB.Unscoped().First(&wallet, "id = ?",walletID).Error
	if err != nil {
		c.JSON(http.StatusNotFound, helpers.ErrorResponse("Data not found", "404"))
		return
	}

	result = helpers.SuccessResponse(wallet, "Success Get Detail Wallet")
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) CreateWallet(c *gin.Context) {
	var (
		wallet models.Wallets
		result gin.H
		createWalletBind WalletCreate
	)

	// validate input
	if err := c.ShouldBindJSON(&createWalletBind); err != nil {
		c.JSON(http.StatusBadRequest, helpers.ErrorResponse(err.Error(), "400"))
		return
	}

	name := c.PostForm("name")
	tipe := c.PostForm("type")
	balance := c.PostForm("balance")
	wallet.Name = name
	wallet.Type = tipe

	balanceFloat, err := strconv.ParseFloat(balance, 64)
	if err != nil {
		wallet.Balance = 0
	} else {
		wallet.Balance = balanceFloat
	}
	
	wallet.Virtual_Account = fmt.Sprintf("%03d-%04d-%04d", rand.Intn(1000), rand.Intn(10000), rand.Intn(10000))
	randomString := fmt.Sprintf("%03d", rand.Intn(1000))
	wallet.Tag_Name = randomString + "_" + name
	// wallet.Key_Phrase = randomString + "_" + name

	err = idb.DB.Create(&wallet).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ErrorResponse(err.Error(), "503"))
		return
	}
	result = gin.H{
		"result": wallet,
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) UpdateWallet(c *gin.Context) {
		id := c.Query("id")
		name := c.PostForm("name")
		tipe := c.PostForm("tipe")

		var (
			wallet    models.Wallets
			newWallet models.Wallets
			result    gin.H
		)

		err := idb.DB.First(&wallet, id).Error
		if err != nil {
			c.JSON(http.StatusNotFound, helpers.ErrorResponse("Data not found", "404"))
			return
		}

		newWallet.Name = name
		newWallet.Type = tipe
		err = idb.DB.Model(&wallet).Updates(newWallet).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, helpers.ServerErrorResponse(err.Error()))
			return
		}

		c.JSON(http.StatusOK, result)
}

func (idb *InDB) DeleteWallet(c *gin.Context) {
	var (
		wallet models.Wallets
		result gin.H
	)

	id := c.Param("id")
	err := idb.DB.First(&wallet, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, helpers.ErrorResponse("Data not found", "404"))
		return
	}

	err = idb.DB.Delete(&wallet, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.ServerErrorResponse(err.Error()))
		return
	} 
	
	c.JSON(http.StatusOK, result)
}