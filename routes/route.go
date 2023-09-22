package routes

import (
	"database/sql"
	"janio-my-backend/api"
	"janio-my-backend/models"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var db *sql.DB

func SetupRoutes(d *sql.DB) {
	db = d
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"},
		AllowMethods:     []string{"POST, GET, OPTIONS, PUT, DELETE, UPDATE"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "Cache-Control"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	router.GET("my/orders", GetOrdersMYProvider)
	router.POST("my/order/update", PostStatus)

	router.Run("localhost:9883")
}

func GetOrdersMYProvider(c *gin.Context) {
	orders, err := api.GetOrdersMYProvider(db)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Orders not found"})
	}
	c.IndentedJSON(http.StatusOK, orders)
}

func PostStatus(c *gin.Context) {
	var status models.MYOrderStatus

	c.BindJSON(&status)

	orderStatus, errS := api.PostStatusMYProvider(db, status)
	if errS != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": errS.Error()})
	}
	c.IndentedJSON(http.StatusOK, orderStatus)
}
