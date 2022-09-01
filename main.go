package main

import (
	"context"
	"fmt"
	"gethired/controller"
	"gethired/services"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server         *gin.Engine
	userservice    services.UserService
	usercontroller controller.UserController
	ctx            context.Context
	usercollection *mongo.Collection
	mongoclient    *mongo.Client
	err            error
)

func init() {
	ctx = context.TODO()
	mongocon := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoclient, err := mongo.Connect(ctx, mongocon)

	if err != nil {
		log.Fatal(err)
	}

	err = mongoclient.Ping(ctx, readpref.Primary())

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("mongo connection establish")

	usercollection = (mongoclient.Database("userdb")).Collection("users")
	userservice = services.NewUserService(usercollection, ctx)
	server = gin.Default()
}

func main() {

	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	// if err != nil {
	// 	println(err)
	// 	panic(err)
	// }

	defer mongoclient.Disconnect(ctx)
	basepath := server.Group("/v1")
	usercontroller.RegisterUserRoutes(basepath)

	log.Fatal(server.Run(":9090"))

	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	// r.Run()
}
