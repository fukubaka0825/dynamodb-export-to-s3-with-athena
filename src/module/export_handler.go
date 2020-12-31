package module

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/labstack/gommon/log"
)

type ExportEvent struct {
	S3BucketName                      string `json:"s3_bucket_name"`
	DynamoTableArn                    string `json:"dynamo_table_arn"`
	KmsArn                            string `json:"kms_arn"`
	AthenaEnabled                     bool   `json:"athena_enabled"`
	AthenaDatabaseName                string `json:"athena_database_name"`
	AthenaQueryResultOutputBucketName string `json:"athena_query_result_output_bucket_name"`
	AthenaTableName                   string `json:"athena_table_name"`
	AthenaTableSchema                 string `json:"athena_table_schema"`
	AthenaCreateTableFrequency        int    `json:"athena_create_table_frequency"`
}

type exportHandler struct {
	Event *ExportEvent
}

func NewExportHandler(event *ExportEvent) *exportHandler {
	return &exportHandler{
		Event: event,
	}
}

func (s *exportHandler) Run() error {
	sess := session.Must(session.NewSession())
	dynamoCli, err := NewDynamoCli(sess)
	if err != nil {
		log.Error(err)
		return err
	}
	exportID, err := dynamoCli.ExportToS3(s.Event.DynamoTableArn, s.Event.S3BucketName, s.Event.KmsArn)
	if err != nil {
		log.Error(err)
		return err
	}
	if s.Event.AthenaEnabled {
		athenaCli, err := NewAthenaCli(sess)
		if err != nil {
			log.Error(err)
			return err
		}
		newLocation := fmt.Sprintf("s3://%s/AWSDynamoDB/%s/data/", s.Event.S3BucketName, exportID)
		outputLocation := fmt.Sprintf("s3://%s/output/", s.Event.AthenaQueryResultOutputBucketName)
		//初回ならCreateTable,初回以降はtable locationだけ新規export objectに設定し直す
		if err := athenaCli.CreateTableIfNotExists(s.Event.AthenaDatabaseName, s.Event.AthenaTableName, s.Event.AthenaTableSchema, newLocation, outputLocation); err != nil {
			log.Error(err)
			return err
		}
		if err := athenaCli.ChangeLocation(s.Event.AthenaDatabaseName, s.Event.AthenaTableName, newLocation, outputLocation); err != nil {
			log.Error(err)
			return err
		}
	}
	return nil
}
