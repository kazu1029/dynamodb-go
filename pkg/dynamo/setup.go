package dynamo

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/awslabs/goformation"
)

type EndpointResolver struct{}

func (e EndpointResolver) ResolveEndpoint(region string, options dynamodb.EndpointResolverOptions) (aws.Endpoint, error) {
	return aws.Endpoint{URL: "http://localhost:8000"}, nil
}

func localDynamoDB() *dynamodb.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("ap-northeast-1"),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider("local", "local", "local")),
	)
	if err != nil {
		log.Fatal("could not setup db connection")
	}

	db := dynamodb.NewFromConfig(cfg, dynamodb.WithEndpointResolver(EndpointResolver{}))

	ctx, cancel := context.WithTimeout(context.Background(), 1500*time.Millisecond)
	defer cancel()
	_, err = db.ListTables(ctx, nil)
	if err != nil {
		log.Fatal("make sure DynamoDB local runs on port :8000", err)
	}
	return db
}

func SetupTable(ctx context.Context, tableName, path string) (*dynamodb.Client, func()) {
	db := localDynamoDB()
	tmpl, err := goformation.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	table, err := tmpl.GetAWSDynamoDBTableWithName(tableName)
	if err != nil {
		log.Fatal(err)
	}
	input := FromCloudFormationToCreateInput(*table)
	_, err = db.CreateTable(ctx, &input)
	if err != nil {
		log.Fatal(err)
	}
	return db, func() {
		_, _ = db.DeleteTable(ctx, &dynamodb.DeleteTableInput{TableName: aws.String(tableName)})
	}
}

func SetupTableWithStream(ctx context.Context, tableName, path string) (*dynamodb.Client, func()) {
	db := localDynamoDB()
	tmpl, err := goformation.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	table, err := tmpl.GetAWSDynamoDBTableWithName(tableName)
	if err != nil {
		log.Fatal(err)
	}
	input := FromCloudFormationToCreateInputWithStream(*table)
	_, err = db.CreateTable(ctx, &input)
	if err != nil {
		log.Fatal(err)
	}
	return db, func() {
		_, _ = db.DeleteTable(ctx, &dynamodb.DeleteTableInput{TableName: aws.String(tableName)})
	}
}
