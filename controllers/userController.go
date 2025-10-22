package controllers

import (
    "context"
    "fmt"
    "log"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"
    "github.com/rahel-yab/Restuarant-Management/database"
    "github.com/rahel-yab/Restuarant-Management/helpers"
    "github.com/rahel-yab/Restuarant-Management/models"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var userValidate = validator.New()

func GetUsers() gin.HandlerFunc {
    return func(c *gin.Context) {
        var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
        defer cancel()

        cursor, err := userCollection.Find(ctx, bson.M{})
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching users"})
            return
        }
        var users []bson.M
        if err := cursor.All(ctx, &users); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "error reading users"})
            return
        }
        c.JSON(http.StatusOK, users)
    }
}

func GetUser() gin.HandlerFunc {
    return func(c *gin.Context) {
        userId := c.Param("user_id")
        var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
        defer cancel()

        var user models.User
        err := userCollection.FindOne(ctx, bson.M{"_id": userId}).Decode(&user)
        if err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
            return
        }
        c.JSON(http.StatusOK, user)
    }
}

func SignUp() gin.HandlerFunc {
    return func(c *gin.Context) {
        var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
        defer cancel()

        var user models.User
        if err := c.BindJSON(&user); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        if err := userValidate.Struct(user); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        // check existing email
        count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "error checking existing user"})
            return
        }
        if count > 0 {
            c.JSON(http.StatusConflict, gin.H{"error": "email already registered"})
            return
        }

        hashed := HashPassword(user.Password)
        user.Password = hashed
        user.ID = primitive.NewObjectID()
        user.Created_at = time.Now()
        user.Updated_at = time.Now()
        user.Access_token, user.Refresh_token, _ = helpers.GenerateAllTokens(user.Email, user.First_name, user.Last_name, user.ID.Hex())

        _, insertErr := userCollection.InsertOne(ctx, user)
        if insertErr != nil {
            log.Println(insertErr)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "user creation failed"})
            return
        }

        // do not return password
        user.Password = ""
        c.JSON(http.StatusCreated, user)
    }
}

func Login() gin.HandlerFunc {
    return func(c *gin.Context) {
        var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
        defer cancel()

        var input struct {
            Email    string `json:"email" binding:"required"`
            Password string `json:"password" binding:"required"`
        }
        if err := c.BindJSON(&input); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        var found models.User
        err := userCollection.FindOne(ctx, bson.M{"email": input.Email}).Decode(&found)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
            return
        }

        if !VerifyPassword(found.Password, input.Password) {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
            return
        }

        access, refresh, _ := helpers.GenerateAllTokens(found.Email, found.First_name, found.Last_name, found.ID.Hex())

        // update tokens in DB
        _, err = userCollection.UpdateOne(ctx, bson.M{"_id": found.ID}, bson.M{"$set": bson.M{"access_token": access, "refresh_token": refresh, "updated_at": time.Now()}})
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update tokens"})
            return
        }

        found.Password = ""
        found.Access_token = access
        found.Refresh_token = refresh
        c.JSON(http.StatusOK, found)
    }
}

func HashPassword(password string) string {
    hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(hash)
}

func VerifyPassword(hashedPassword string, providedPassword string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(providedPassword))
    return err == nil
}