# What is this
- dynamodb export to s3 with athena 
## Variables
|name | type| is_required|description|
| ---- | ---- | ---- | ---- |
|cron(schedule)|String|true| |
|cron(enabled)|Bool|true| |
|DynamoDB Table Arn|String|true| |
|S3 Bukcet Name|String|true| |
|Kms Key ARN|String|false| |
|Athena enabled|Bool|false| |
|Athena query_result_output_bucket_name|String|false| |
|Athena database_name|String|false| |
|Athena table_name|String|false| |
|Athena table_schema|String|false| |
