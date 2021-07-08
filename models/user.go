package models

type User struct {
	Email          string `dynamo:"Email,hash"`
	OrganizationID string `dynamo:"OrganizationID"`
	CouponCode     string `dynamo:"CouponCode"`
	UserName       string `dynamo:"UserName"`
}
