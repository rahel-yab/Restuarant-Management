package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rahel-yab/Restuarant-Management/database"
	"github.com/rahel-yab/Restuarant-Management/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var orderItemCollection *mongo.Collection = database.OpenCollection(database.Client, "order_item")
var orderItemValidate = validator.New()

func GetOrderItems() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		cursor, err := orderItemCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching order items"})
			return
		}
		var items []bson.M
		if err := cursor.All(ctx, &items); err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error reading order items"})
			return
		}
		c.JSON(http.StatusOK, items)
	}
}

func GetOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("order_item_id")
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var item models.OrderItem
		err := orderItemCollection.FindOne(ctx, bson.M{"order_item_id": id}).Decode(&item)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "order item not found"})
			return
		}
		c.JSON(http.StatusOK, item)
	}
}

func CreateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var item models.OrderItem
		if err := c.BindJSON(&item); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := orderItemValidate.Struct(item); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		item.ID = primitive.NewObjectID()
		item.Order_item_id = item.ID.Hex()
		item.Created_at = time.Now()
		item.Updated_at = time.Now()

		_, err := orderItemCollection.InsertOne(ctx, item)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create order item"})
			return
		}
		c.JSON(http.StatusCreated, item)
	}
}

func UpdateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("order_item_id")
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var payload models.OrderItem
		if err := c.BindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var updateObj primitive.D
		if payload.Quantity != nil {
			updateObj = append(updateObj, bson.E{Key: "quantity", Value: payload.Quantity})
		}
		if payload.Unit_price != nil {
			updateObj = append(updateObj, bson.E{Key: "unit_price", Value: payload.Unit_price})
		}
		if payload.Order_id != nil {
			updateObj = append(updateObj, bson.E{Key: "order_id", Value: payload.Order_id})
		}
		payload.Updated_at = time.Now()
		updateObj = append(updateObj, bson.E{Key: "updated_at", Value: payload.Updated_at})

		upsert := true
		filter := bson.M{"order_item_id": id}
		opt := options.UpdateOptions{Upsert: &upsert}

		result, err := orderItemCollection.UpdateOne(ctx, filter, bson.D{{Key: "$set", Value: updateObj}}, &opt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "update failed"})
			return
		}
		c.JSON(http.StatusOK, result)
	}
}

func GetOrderItemsByOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		orderId := c.Param("order_id")
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		cursor, err := orderItemCollection.Find(ctx, bson.M{"order_id": orderId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching order items"})
			return
		}
		var items []bson.M
		if err := cursor.All(ctx, &items); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error reading order items"})
			return
		}
		c.JSON(http.StatusOK, items)
	}
}
