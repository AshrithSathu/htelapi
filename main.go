package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/AshrithSathu/htelapi/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const dburi = "mongodb://localhost:27017"
const dbname = "htelapi"
const userColl = "users"

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().
		ApplyURI(dburi))
	if err != nil {
		fmt.Println("Error connecting to MongoDB")
		log.Fatal("error in connection: ", err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal("MongoDB not reachable: ", err)
	}
	ctxBackground := context.Background()
	coll := client.Database(dbname).Collection(userColl)

	user := types.User{
		FirstName: "James",
		LastName:  "Hello",
	}
	_, err = coll.InsertOne(ctxBackground, user)
	if err != nil {
		log.Fatal("Error inserting user: ", err)
	}

	var james types.User
	if err = coll.FindOne(ctxBackground, bson.M{}).Decode(&james); err != nil {
		log.Fatal("Error finding user: ", err)
	}

	fmt.Println(james)

	listenAddr := flag.String("listenAddr", ":4000", "server listen address")
	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	// userHandler := api.NewUserHandler()
	// apiv1.Get("/user", api.HandleGetUsers)
	// apiv1.Get("/user:id", api.HandleGetUser)
	app.Listen(*listenAddr)
}

func handleFoo(c *fiber.Ctx) error {
	return c.JSON(map[string]string{
		"message": "Foo",
	})
}

func handleUser(c *fiber.Ctx) error {
	return c.JSON(map[string]string{
		"user": "James Foo",
	})
}
