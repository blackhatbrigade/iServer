package main

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoClient struct {
	client *mongo.Client
	ctx    context.Context
}

func NewClient(ctx context.Context, uri string) *mongoClient {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	return &mongoClient{client: client, ctx: ctx}
}
