package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func apiGetTemperature(c *gin.Context) {
	var temperatures Temperature = getLastTemperature()

	c.PureJSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   temperatures,
	})
}

func apiGetTemperatures(c *gin.Context) {
	var temperatures []Temperature = getAllTemperatures()

	c.PureJSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"count":  len(temperatures),
		"data":   temperatures,
	})
}

func apiStoreTemperature(c *gin.Context) {
	var temperature Temperature
	c.BindJSON(&temperature)

	insertTemperature(temperature)

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
