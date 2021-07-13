package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/kazu1029/dynamodb-go/pkg/dynamo"
)

type Order struct {
	ID        string `dynamodbav:"id"`
	Price     int    `dynamodbav:"price"`
	IsShipped bool   `dynamodbav:"is_shipped"`
}

func main() {
	ctx := context.Background()
	tableName := "OrdersTable"
	db, cleanup := dynamo.SetupTable(ctx, tableName, "./putget/template.yml")
	defer cleanup()

	order := Order{
		ID:        "12-34",
		Price:     22,
		IsShipped: false,
	}

	avs, err := attributevalue.MarshalMap(order)
	if err != nil {
		panic(err)
	}

	pOut, err := db.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      avs,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("pOut: %#v\n", pOut)

	gOut, err := db.GetItem(ctx, &dynamodb.GetItemInput{
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{
				Value: "12-34",
			},
		},
		TableName: aws.String(tableName),
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("gOut: %#v\n", gOut)
}
