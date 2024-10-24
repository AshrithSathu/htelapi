package db

import (
	"github.com/AshrithSathu/htelapi/types"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserStore interface {
	GetUserByID(string) (*types.User, error)
}

type MongoUserStore struct {
	client *mongo.Client
}

// func NewMongoUserStore(c *mongo.Client) UserStore {
// 	return &MongoUserStore{
// 		client : client
// 	}
// }

type PostgresUserStore struct {
}
