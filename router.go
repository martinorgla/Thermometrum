package main

import (
	"github.com/gin-gonic/contrib/static"
	_ "github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"net/http"
	"time"
)

func setupRouter() {
	// Set the router as the default one shipped with Gin
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(cors.Default())

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./client/build", true)))

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
			c.Request.Context().Done()
		})

		api.GET("/temperature", apiGetTemperature)
		api.GET("/temperatures", apiGetTemperatures)
		api.POST("/temperature", apiStoreTemperature)
	}

	// Start and run the server
	router.Run(":8001")

	srv := &http.Server{
		Addr:         ":8001",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	srv.ListenAndServe()
}
