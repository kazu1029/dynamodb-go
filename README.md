# dynamodb-go

## Create table
```
aws dynamodb create-table --endpoint-url http://localhost:8000 \
  --table-name User \
  --attribute-definitions \
    AttributeName=Email,AttributeType=S \
    AttributeName=OrganizationID,AttributeType=S \
    AttributeName=CouponCode,AttributeType=S \
  --key-schema \
    AttributeName=Email,KeyType=HASH \
    AttributeName=OrganizationID,KeyType=RANGE \
  --provisioned-throughput \
    ReadCapacityUnits=5,WriteCapacityUnits=5 \
  --global-secondary-indexes \
    IndexName=GSI-Email-OrganizationID,KeySchema=['{AttributeName=Email,KeyType=HASH},{AttributeName=OrganizationID,KeyType=RANGE}'],Projection={ProjectionType=ALL},ProvisionedThroughput={'ReadCapacityUnits=5,WriteCapacityUnits=5'} \
    IndexName=GSI-Email-CouponCode,KeySchema=['{AttributeName=Email,KeyType=HASH},{AttributeName=CouponCode,KeyType=RANGE}'],Projection={ProjectionType=ALL},ProvisionedThroughput={'ReadCapacityUnits=5,WriteCapacityUnits=5'}
```

```
aws dynamodb create-table --endpoint-url http://localhost:8000 \
  --table-name User \
  --attribute-definitions \
    AttributeName=Email,AttributeType=S \
    AttributeName=OrganizationID,AttributeType=S \
    AttributeName=CouponCode,AttributeType=S \
  --key-schema \
    AttributeName=Email,KeyType=HASH \
    AttributeName=OrganizationID,KeyType=RANGE \
  --provisioned-throughput \
    ReadCapacityUnits=5,WriteCapacityUnits=5 \
  --global-secondary-indexes \
    IndexName=GSI-Email-OrganizationID,KeySchema=['{AttributeName=Email,KeyType=HASH},{AttributeName=OrganizationID,KeyType=RANGE}'],Projection={ProjectionType=ALL},ProvisionedThroughput={'ReadCapacityUnits=5,WriteCapacityUnits=5'} \
  --local-secondary-indexes \
    IndexName=GSI-Email-CouponCode,KeySchema=['{AttributeName=Email,KeyType=HASH},{AttributeName=CouponCode,KeyType=RANGE}'],Projection={ProjectionType=ALL}
```
