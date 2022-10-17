package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rwiteshbera/authentication-go-gin/database"
	"github.com/rwiteshbera/authentication-go-gin/helpers"
	"github.com/rwiteshbera/authentication-go-gin/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

// var validate = validator.New()

func HashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 16)
	if err != nil {
		log.Panic(err)
	}
	return string(hashedPassword)
}

func VerifyPassword(userPassword string, hashedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(userPassword))

	check := true
	msg := ""

	if err != nil {
		msg = "password is incorrect"
		check = false
	}
	return check, msg
}

func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User

		err := c.BindJSON(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

		countEmail, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking the email"})
			return
		}

		if countEmail > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "this email is already exists"})
			return
		}

		password := HashPassword(user.Password)
		user.Password = password

		user.UserId = primitive.NewObjectID().Hex()
		user.CreatedAt, _ = time.Parse(time.RFC1123, time.Now().Format(time.RFC1123))

		result, insertErr := userCollection.InsertOne(ctx, user) // Insert data in db
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": insertErr.Error()})
		}
		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		var savedUser models.User

		err := c.BindJSON(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		userErr := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&savedUser)
		defer cancel()
		if userErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No user found"})
			return
		}

		isPasswordValid, msg := VerifyPassword(user.Password, savedUser.Password)

		if !isPasswordValid {
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		token, err := helpers.GenerateToken(savedUser.Email, savedUser.FirstName, savedUser.LastName, savedUser.UserId)
		if err != nil {
			log.Fatal(err)
			return
		}

		userError := userCollection.FindOne(ctx, bson.M{"user_id": savedUser.UserId}).Decode(&savedUser)
		if userErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": userError.Error()})
		}
		defer cancel()
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}
