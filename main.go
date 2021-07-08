package main

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/kazu1029/dynamodb-go/models"
)

func main() {
	var (
		dynamoDBRegion   = os.Getenv("AWS_REGION")
		disableSSL       = false
		dynamoDBEndpoint = os.Getenv("DYNAMO_ENDPOINT")
	)

	if dynamoDBRegion == "" {
		dynamoDBRegion = "ap-northeast-1"
	}

	sess, err := session.NewSession()
	if err != nil {
		panic(err)
	}

	db := dynamo.New(sess, &aws.Config{
		Region:     aws.String(dynamoDBRegion),
		Endpoint:   aws.String(dynamoDBEndpoint),
		DisableSSL: aws.Bool(disableSSL),
	})

	table := db.Table("User")
	user := models.User{
		Email:          "email_1@example.com",
		OrganizationID: "OrganizationID_1",
		CouponCode:     "CouponCode__1",
		UserName:       "UserName_1",
	}

	if err := table.Put(user).Run(); err != nil {
		panic(err)
	}
}
