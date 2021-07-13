package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/kazu1029/dynamodb-go/pkg/dynamo"
)

type item struct {
	Directory string `dynamodbav:"directory"`
	Filename  string `dynamodbav:"filename"`
	Size      string `dynamodbav:"size"`
}

func main() {
	ctx := context.Background()
	tableName := "FileSystemTable"
	db, cleanup := dynamo.SetupTable(ctx, tableName, "./compositeprimarykeys/template.yml")
	defer cleanup()

	insert(ctx, db, tableName)

	fmt.Printf("\n=========================== Start Simple Get ============================\n")
	simpleGet(ctx, db, tableName)
	fmt.Printf("=========================== End Simple Get ============================\n")

	fmt.Printf("\n=========================== Start Get All ============================\n")
	getAll(ctx, db, tableName)
	fmt.Printf("=========================== End Get All ============================\n")

	fmt.Printf("\n=========================== Start Get All Reports Before 2019 ============================\n")
	getReportsBefore2019(ctx, db, tableName)
	fmt.Printf("=========================== End Get All Reports Before 2019 =============================\n")

}

func simpleGet(ctx context.Context, db *dynamodb.Client, tableName string) {
	out, err := db.GetItem(ctx, &dynamodb.GetItemInput{
		Key: map[string]types.AttributeValue{
			"directory": &types.AttributeValueMemberS{Value: "finances"},
			"filename":  &types.AttributeValueMemberS{Value: "report2020.pdf"},
		},
		TableName: aws.String(tableName),
	})
	if err != nil {
		panic(err)
	}

	var i item
	err = attributevalue.UnmarshalMap(out.Item, &i)
	if err != nil {
		panic(err)
	}
	fmt.Printf("item: %#v\n", i)
}

func getAll(ctx context.Context, db *dynamodb.Client, tableName string) {
	expr, err := expression.NewBuilder().
		WithKeyCondition(
			expression.KeyEqual(expression.Key("directory"), expression.Value("finances")),
		).Build()
	if err != nil {
		panic(err)
	}

	out, err := db.Query(ctx, &dynamodb.QueryInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		TableName:                 aws.String(tableName),
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

func getReportsBefore2019(ctx context.Context, db *dynamodb.Client, tableName string) {
	expr, err := expression.NewBuilder().
		WithKeyCondition(
			expression.KeyAnd(
				expression.KeyEqual(expression.Key("directory"), expression.Value("finances")),
				expression.KeyLessThan(expression.Key("filename"), expression.Value("report2019")),
			),
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

func insert(ctx context.Context, db *dynamodb.Client, tableName string) {
	item1 := item{Directory: "finances", Filename: "report2017.pdf", Size: "1MB"}
	item2 := item{Directory: "finances", Filename: "report2018.pdf", Size: "1MB"}
	item3 := item{Directory: "finances", Filename: "report2019.pdf", Size: "1MB"}
	item4 := item{Directory: "finances", Filename: "report2020.pdf", Size: "2MB"}
	item5 := item{Directory: "fun", Filename: "game01.pdf", Size: "1MB"}

	for _, item := range []item{item1, item2, item3, item4, item5} {
		attrs, _ := attributevalue.MarshalMap(&item)
		_, _ = db.PutItem(ctx, &dynamodb.PutItemInput{
			TableName: aws.String(tableName),
			Item:      attrs,
		})
	}
}
