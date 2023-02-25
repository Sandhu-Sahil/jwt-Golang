package main

import (
	"context"
	"fmt"
	"jwt-golang/controllers"
	"jwt-golang/routes"
	"jwt-golang/services"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server      *gin.Engine
	us          services.UserService
	uc          controllers.UserController
	rs          routes.RouterService
	ctx         context.Context
	userc       *mongo.Collection
	mongoclient *mongo.Client
	err         error
	// validate    *validator.Validate
)

func init() {
	ctx = context.TODO()

	mongoconn := options.Client().ApplyURI("mongodb+srv://Sandhu-Sahil:ssssandhu26@cluster0.yfxftsl.mongodb.net/?retryWrites=true&w=majority")
	mongoclient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal("error while connecting with mongo", err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("error while trying to ping mongo", err)
	}

	fmt.Println("mongo connection established")

	userc = mongoclient.Database("jwt-golang").Collection("users")
	us = services.NewUserService(userc, ctx)
	uc = controllers.New(us)
	rs = routes.NewRouterService(uc)
	server = gin.Default()

	// validate = validator.New()
}

func main() {
	defer mongoclient.Disconnect(ctx)

	basepath := server.Group("/api-v1")
	rs.RegisterLoginRoutes(basepath)

	jwtpath := server.Group("/jwt")
	rs.RegisterJwtCheckRoutes(jwtpath)

	log.Fatal(server.Run(":8080"))
}
