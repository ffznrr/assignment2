package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type Orders struct {
	OrderID      int `json:"PrimaryKey"`
	OrderedAt    time.Time
	CustomerName string
	Items        []Items
}

type Items struct {
	ItemID      int    `json:"PrimaryKey"`
	ItemCode    string `json:"ItemCode"`
	Description string `json:"Description"`
	Quantity    int    `json:"Quantity"`
}

var OrderData []Orders

func CreateOrders(c *gin.Context) {
	var orders Orders
	var Items []gin.H // Declare Items slice here

	err := c.ShouldBindJSON(&orders)
	if err != nil {
		// Handle validation errors
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on Field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return // Return early to avoid continuing with invalid data
	}

	// Process valid data
	for _, item := range orders.Items {
		Items = append(Items, gin.H{
			"ItemCode":    item.ItemCode,
			"description": item.Description,
			"quantity":    item.Quantity,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"OrderedAt":    orders.OrderedAt,
		"CustomerName": orders.CustomerName,
		"Items":        Items,
	})
}

func ReadOrders(c *gin.Context) {
	var order Orders

	c.JSON(http.StatusOK, gin.H{
		"id":           order.OrderID,
		"orderedAt":    "senin",
		"customerName": "fauzan",
		"items": []gin.H{
			{
				"itemCode":    "123",
				"description": "iphoneXR",
				"quantity":    2,
			},
		},
	})
}

func UpdateOrder(ctx *gin.Context) {
	orderIDParam := ctx.Param("orderID")

	// Convert orderIDParam from string to int
	orderID, err := strconv.Atoi(orderIDParam)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	condition := false

	var updatedOrder Orders

	err = ctx.ShouldBindJSON(&updatedOrder)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, order := range OrderData {
		if orderID == order.OrderID {
			condition = true

			// Update the order in the slice
			OrderData[i] = updatedOrder
			OrderData[i].OrderID = orderID
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "data Not Found",
			"error_message": fmt.Sprintf("order with id %v not found", orderID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("order with id %v has been successfully updated", orderID),
	})
}

func DeleteOrder(ctx *gin.Context) {
	orderIDParam := ctx.Param("OrderID")

	// Convert orderID from string to int
	orderID, err := strconv.Atoi(orderIDParam)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	condition := false

	for i, order := range OrderData {
		if orderID == order.OrderID {
			condition = true

			// Remove the order from the slice
			OrderData = append(OrderData[:i], OrderData[i+1:]...)
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "data Not Found",
			"error_message": fmt.Sprintf("order with id %v not found", orderID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("order with id %v has been successfully deleted", orderID),
	})
}
