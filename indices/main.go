package main

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/kazu1029/dynamodb-go/pkg/dynamo"
)

type item struct {
	Directory string    `dynamodbav:"directory"`
	Filename  string    `dynamodbav:"filename"`
	Size      string    `dynamodbav:"size"`
	CreatedAt time.Time `dynamodbav:"created_at"`
}

func main() {
	ctx := context.Background()
	tableName := "FileSystemTable"
	db, cleanup := dynamo.SetupTable(ctx, tableName, "./indices/template.yml")
	defer cleanup()

	fmt.Printf("\n========================= Start photosTakensFrom2019 ==========================\n")
	photosTakensFrom2019(ctx, db, tableName)
	fmt.Printf("========================= End photosTakensFrom2019 ==========================\n")
}

func photosTakensFrom2019(ctx context.Context, db *dynamodb.Client, tableName string) {
	expr, err := expression.NewBuilder().
		WithKeyCondition(
			expression.KeyAnd(
				expression.KeyEqual(
					expression.Key("directory"),
					expression.Value("photos"),
				),
				expression.KeyGreaterThanEqual(
					expression.Key("created_at"),
					expression.Value(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)),
				)),
		).
		Build()
	if err != nil {
		panic(err)
	}

	out, err := db.Query(ctx, &dynamodb.QueryInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		TableName:                 aws.String(tableName),
		IndexName:                 aws.String("ByCreatedAt"),
	})
	if err != nil {
		panic(err)
	}

	var items []item
	err = attributevalue.UnmarshalListOfMaps(out.Items, &items)
	if err != nil {
		panic(err)
	}

	for _, item := range items {
		fmt.Printf("item: %#v\n", item)
	}
}
