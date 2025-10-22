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

var invoiceCollection *mongo.Collection = database.OpenCollection(database.Client, "invoice")
var invoiceValidate = validator.New()

func GetInvoices() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		cursor, err := invoiceCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching invoices"})
			return
		}
		var invoices []bson.M
		if err := cursor.All(ctx, &invoices); err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error reading invoices"})
			return
		}
		c.JSON(http.StatusOK, invoices)
	}
}

func GetInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		invoiceId := c.Param("invoice_id")
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var invoice models.Invoice
		err := invoiceCollection.FindOne(ctx, bson.M{"invoice_id": invoiceId}).Decode(&invoice)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "invoice not found"})
			return
		}
		c.JSON(http.StatusOK, invoice)
	}
}

func CreateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var inv models.Invoice
		if err := c.BindJSON(&inv); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := invoiceValidate.Struct(inv); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		inv.ID = primitive.NewObjectID()
		inv.Invoice_id = inv.ID.Hex()
		inv.Created_at = time.Now()
		inv.Updated_at = time.Now()

		_, err := invoiceCollection.InsertOne(ctx, inv)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create invoice"})
			return
		}
		c.JSON(http.StatusCreated, inv)
	}
}

func UpdateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		invoiceId := c.Param("invoice_id")
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var payload models.Invoice
		if err := c.BindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var updateObj primitive.D
		if payload.Payment_method != nil {
			updateObj = append(updateObj, bson.E{Key: "payment_method", Value: payload.Payment_method})
		}
		if payload.Payment_status != nil {
			updateObj = append(updateObj, bson.E{Key: "payment_status", Value: payload.Payment_status})
		}
		if !payload.Payment_due_date.IsZero() {
			updateObj = append(updateObj, bson.E{Key: "payment_due_date", Value: payload.Payment_due_date})
		}

		payload.Updated_at = time.Now()
		updateObj = append(updateObj, bson.E{Key: "updated_at", Value: payload.Updated_at})

		upsert := true
		filter := bson.M{"invoice_id": invoiceId}
		opt := options.UpdateOptions{Upsert: &upsert}

		result, err := invoiceCollection.UpdateOne(ctx, filter, bson.D{{Key: "$set", Value: updateObj}}, &opt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "update failed"})
			return
		}
		c.JSON(http.StatusOK, result)
	}
}
