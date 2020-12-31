package module

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/labstack/gommon/log"
)

type DynamoCli struct {
	sdkClient *dynamodb.DynamoDB
}

// New returns initialized *S4.
func NewDynamoCli(sess *session.Session) (*DynamoCli, error) {
	svc := dynamodb.New(sess, aws.NewConfig().WithRegion(DefaultAwsRegion))
	return &DynamoCli{
		sdkClient: svc,
	}, nil
}

func (c *DynamoCli) ExportToS3(dynamoTableArn, s3BukcetName, kmsArn string) (string, error) {
	const exportIDOrderNum = 3
	input := &dynamodb.ExportTableToPointInTimeInput{
		ExportFormat: aws.String("DYNAMODB_JSON"),
		S3Bucket:     aws.String(s3BukcetName),
		TableArn:     aws.String(dynamoTableArn),
	}
	if kmsArn == "" {
		input.S3SseAlgorithm = aws.String("AES256")
		input.S3SseKmsKeyId = nil
	}
	input.S3SseAlgorithm = aws.String("KMS")
	input.S3SseKmsKeyId = aws.String(kmsArn)
	output, err := c.sdkClient.ExportTableToPointInTime(input)
	if err != nil {
		log.Error(err)
		return "", err
	}

	return strings.Split(aws.StringValue(output.ExportDescription.ExportArn), "/")[exportIDOrderNum], nil
}
