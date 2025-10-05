package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rahel-yab/Restuarant-Management/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var orderCollection *mongo.Collection = database.OpenCollection(database.Client, "order")

func GetOrders() gin.HandlerFunc {
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		
		result , err := orderCollection.Find(context.TODO(), bson.M{})
		defer cancel()
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error" : "error occured while listing order items"})

		}
		var allOrders []bson.M
		if err = result.All(ctx, &allOrders); err != nil{
			
			c.JSON(http.StatusInternalServerError, gin.H{"error":"error occured while listing food items"})
			return
		}
		c.JSON(http.StatusOK, allOrders)
	}

}

func GetOrder() gin.HandlerFunc {
	return func(c *gin.Context){
		
	}
}

func CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context){
		
	}
}

func UpdateOrder() gin.HandlerFunc {
	return func(c *gin.Context){
		
	}
}
