package dynamodb

import (
	"context"
	"fmt"

	awsutils "github.com/LouisHatton/user-audit-saas/internal/aws/utils"
	"github.com/LouisHatton/user-audit-saas/internal/store"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"go.uber.org/zap"
)

var _ store.Store = (*DynamoDb)(nil)

type DynamoDb struct {
	log       *zap.Logger
	client    *dynamodb.Client
	TableName string
}

func New(log zap.Logger, tableName string) *DynamoDb {
	cfg := DynamoDb{
		TableName: tableName,
		log:       &log,
	}

	awsEndpoint := "http://localhost:8000"
	awsRegion := "eu-west-1"

	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		if awsEndpoint != "" {
			return aws.Endpoint{
				PartitionID:   "aws",
				URL:           awsEndpoint,
				SigningRegion: awsRegion,
			}, nil
		}

		// returning EndpointNotFoundError will allow the service to fallback to its default resolution
		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})

	awsCfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithEndpointResolverWithOptions(customResolver),
	)
	if err != nil {
		log.Fatal("Cannot load the AWS configs", zap.Error(err))
	}

	svc := dynamodb.NewFromConfig(awsCfg)
	cfg.client = svc

	return &cfg
}

func (db *DynamoDb) Get(id string, doc interface{}) error {
	input := dynamodb.GetItemInput{
		TableName: aws.String(db.TableName),
		Key:       awsutils.StringKey("id", id),
	}

	output, err := db.client.GetItem(context.TODO(), &input)
	if err != nil {
		return fmt.Errorf("failed to get item from dynamodb: %s", err.Error())
	}

	if len(output.Item) == 0 {
		return fmt.Errorf("returned item from dynamodb is empty")
	}

	err = attributevalue.UnmarshalMapWithOptions(output.Item, &doc, awsutils.JsonUnmarshalOptions)
	if err != nil {
		return fmt.Errorf("failed to unmarshal from dynamodb: %s", err.Error())
	}

	return nil
}

func (db *DynamoDb) Put(doc interface{}) error {
	input, err := attributevalue.MarshalMapWithOptions(doc, awsutils.JsonMarshalOptions)
	if err != nil {
		return fmt.Errorf("failed to marshal data for dynamodb: %s", err.Error())
	}

	_, err = db.client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(db.TableName),
		Item:      input,
	})
	if err != nil {
		return fmt.Errorf("failed to put item into dynamodb: %s", err.Error())
	}
	return nil
}
