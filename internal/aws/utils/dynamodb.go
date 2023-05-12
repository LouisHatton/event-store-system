package utils

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type AWSObject = map[string]types.AttributeValue

func StringKey(name, value string) AWSObject {
	return AWSObject{
		name: StringValue(value),
	}
}

func StringValue(value string) *types.AttributeValueMemberS {
	return &types.AttributeValueMemberS{Value: value}
}

func JsonMarshalOptions(opt *attributevalue.EncoderOptions) {
	opt.TagKey = "json"
}

func JsonUnmarshalOptions(opt *attributevalue.DecoderOptions) {
	opt.TagKey = "json"
}
