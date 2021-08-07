package main

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodbstreams"
)

type Sample struct {
	PK string `dynamodbav:"pk"`
	SK string `dynamodbav:"sk"`
}

func main() {
	// ctx := context.Background()
	// db, cleanup := dynamo.SetupTableWithStream(ctx, "SampleTable", "./stream/template.yml")
	// defer cleanup()

	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion("ap-northeast-1"),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider("local", "local", "local")),
	)
	if err != nil {
		panic(err)
	}

	dynamo := dynamodb.NewFromConfig(cfg)
	dynamoStream := dynamodbstreams.NewFromConfig(cfg)
}
