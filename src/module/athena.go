package module

import (
	"errors"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/athena"
)

type AthenaCli struct {
	sdkClient *athena.Athena
}

// New returns initialized *S4.
func NewAthenaCli(sess *session.Session) (*AthenaCli, error) {
	svc := athena.New(sess, aws.NewConfig().WithRegion(DefaultAwsRegion))
	return &AthenaCli{
		sdkClient: svc,
	}, nil
}

func (c *AthenaCli) CreateTableIfNotExists(athenaDatabaseName, athenaTableName, athenaTableSchema, newLocation, outputLocation string) error {
	query := fmt.Sprintf(`
CREATE EXTERNAL TABLE IF NOT EXISTS %s.%s(
   item %s COMMENT 'from deserializer')
ROW FORMAT SERDE 
  'org.openx.data.jsonserde.JsonSerDe' 
STORED AS INPUTFORMAT 
  'org.apache.hadoop.mapred.TextInputFormat' 
OUTPUTFORMAT 
  'org.apache.hadoop.hive.ql.io.HiveIgnoreKeyTextOutputFormat'
LOCATION
  '%s'
TBLPROPERTIES (
  'classification'='json', 
  'compressionType'='gzip', 
  'last_modified_by'='hadoop',
  'typeOfData'='file')
`, athenaDatabaseName, athenaTableName, athenaTableSchema, newLocation)
	if err := c.executeQuery(query, outputLocation); err != nil {
		return err
	}
	return nil
}

func (c *AthenaCli) ChangeLocation(athenaDatabaseName, athenaTableName, newLocation, outputLocation string) error {
	query := fmt.Sprintf("ALTER TABLE %s.%s SET LOCATION '%s'", athenaDatabaseName, athenaTableName, newLocation)
	if err := c.executeQuery(query, outputLocation); err != nil {
		return err
	}
	return nil
}

func (c *AthenaCli) executeQuery(query, resultLocation string) error {
	resultConf := &athena.ResultConfiguration{}
	resultConf.SetOutputLocation(resultLocation)

	input := &athena.StartQueryExecutionInput{
		QueryString:         aws.String(query),
		ResultConfiguration: resultConf,
	}

	output, err := c.sdkClient.StartQueryExecution(input)
	if err != nil {
		return err
	}

	executionInput := &athena.GetQueryExecutionInput{
		QueryExecutionId: output.QueryExecutionId,
	}

	// 終わるまで待つ
	var executionOutput *athena.GetQueryExecutionOutput

	for {
		executionOutput, err = c.sdkClient.GetQueryExecution(executionInput)
		if err != nil {
			return err
		}
		// executionOutput.QueryExecution.Status.Stateは*string
		switch *executionOutput.QueryExecution.Status.State {
		case athena.QueryExecutionStateQueued, athena.QueryExecutionStateRunning:
			time.Sleep(5 * time.Second)
		case athena.QueryExecutionStateSucceeded:
			return nil
		default:
			return errors.New(executionOutput.String())
		}
	}
}
