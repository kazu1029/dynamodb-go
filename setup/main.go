package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/kazu1029/dynamodb-go/pkg/dynamo"
)

func main() {
	ctx := context.Background()
	db, cleanup := dynamo.SetupTable(ctx, "PartitionKeyTable", "./pkg/dynamo/testdata/template.yml")
	defer cleanup()

	out, err := db.DescribeTable(ctx, &dynamodb.DescribeTableInput{
		TableName: aws.String("PartitionKeyTable"),
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("out: %#v\n", out)
}
