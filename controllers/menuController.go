package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rahel-yab/Restuarant-Management/database"
	"github.com/rahel-yab/Restuarant-Management/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var menuCollection *mongo.Collection = database.OpenCollection(database.Client, "menu")

func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context){
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		result , err := menuCollection.Find(context.TODO(), bson.M{})
		defer cancel()
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":"error occured while listing the menu items"})
		}

		var allMenus []bson.M
		if err = result.All(ctx, &allMenus); err != nil{
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allMenus)
	}
}

func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context){
		var ctx , cancel = context.WithTimeout(context.Background(), 100*time.Second)
		menuId := c.Param("menu_id")
		var menu models.Menu

		err := foodCollection.FindOne(ctx, bson.M{"menu_id":menuId}).Decode(&menu)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":"error occured while fetching the menu item"})
		}
		c.JSON(http.StatusOK, menu)
	}
}

func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context){
		var menu  models.Menu
		var ctx , cancel = context.WithTimeout(context.Background(), 100*time.Second)
		if err := c.BindJSON(&menu); err != nil{
			c.JSON(http.StatusBadRequest , gin.H{"error": err.Error()})
			return
		}

		validateionErr := validate.Struct(menu)
		if validateionErr != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error":validateionErr.Error()})
			return
		}

		menu.Created_at , _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		menu.Updated_at , _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		menu.ID = primitive.NewObjectID()
		menu.Menu_id = menu.ID.Hex()

		result , insertErr := menuCollection.InsertOne(ctx,menu)
		if insertErr != nil{
			msg := fmt.Sprintf("menu item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error" : msg})
			return 
		}

		defer cancel()
		c.JSON(http.StatusCreated, result)
		defer cancel()
	}
}

func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context){
		var ctx , cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var menu models.Menu

		if err := c.BindJSON(&menu); err != nil{
			c.JSON(http.StatusBadRequest , gin.H{"error": err.Error()})
			return
		}

		menuId := c.Param("menu_id")
		filter := bson.M{"menu_id":menuId}

		var updateObj primitive.D
		if menu.Start_date != nil && menu.End_date != nil{
			if !inTimeSpan(*menu.Start_date, *menu.End_date, time.Now()){
				msg := "Kindly retype the time"
				c.JSON(http.StatusInternalServerError, gin.H{"error":msg})
				defer cancel()
				return
			}

		updateObj = append(updateObj, bson.E{"start_date", menu.Start_date})
		updateObj = append(updateObj, bson.E{"end_date", menu.End_date})
		
		if menu.Name != ""{
			updateObj = append(updateObj, bson.E{"name", menu.Name})
		}
		if menu.Category != ""{
			updateObj = append(updateObj, bson.E{"category", menu.Category})
		}

		menu.Updated_at , _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{"updated_at", menu.Updated_at})
		upsert := true
		opt := options.UpdateOptions{
			Upsert : &upsert,

		}
		result , err := menuCollection.UpdateOne(
		ctx,
		filter,
		bson.D{
			{"$set", updateObj},
		},
		&opt,
	
	)
	if err != nil{
		msg := "menu update failed"
		c.JSON(http.StatusInternalServerError, gin.H{"error" : msg})
		return
	}

	defer cancel()

	c.JSON(http.StatusOK, result)
}

}
}