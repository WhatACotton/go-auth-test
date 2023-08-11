package handler

import (
	"net/http"
	"strconv"
	"unify/internal/database"
	"unify/internal/models"

	"github.com/gin-gonic/gin"
)

// itemlist
func PostItemList(c *gin.Context) {
	var newItem models.ItemList
	c.BindJSON(&newItem)
	c.JSON(http.StatusOK, database.PostItemList(newItem))
}

func GetItemList(c *gin.Context) {
	c.JSON(http.StatusOK, database.GetItemList())
}

func UpdateItemList(c *gin.Context) {
	ItemId := c.Query("itemid")
	InfoId := c.Query("infoid")
	database.UpdateItemList(ItemId, InfoId)
	c.JSON(http.StatusOK, "done!!")
}

func ChangeItemStatus(c *gin.Context) {
	ItemId := c.Query("itemid")
	StatusCode, _ := strconv.Atoi(c.Query("status"))
	database.ChangeItemStatus(ItemId, StatusCode)
	c.JSON(http.StatusOK, "done!!")
}

// iteminfo
func PostItemInfo(c *gin.Context) {
	var newItem models.ItemInfo
	c.BindJSON(&newItem)
	database.PostItemInfo(newItem)
	c.JSON(http.StatusOK, database.GetItemInfo(newItem.InfoId))
}

func GetItemInfo(c *gin.Context) {
	InfoId := c.Query("infoid")
	c.JSON(http.StatusOK, database.GetItemInfo(InfoId))
}

func UpdateItemInfo(c *gin.Context) {
	ItemInfo := new(models.ItemInfo)
	database.UpdateItemInfo(*ItemInfo)
	c.JSON(http.StatusOK, "done!!")
}
