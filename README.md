# What is this
- Dynamodb export to s3 with Athena(Terraform ＋ Serverless Framework)
- REF: https://fukubaka0825.hatenablog.com/entry/2020/12/31/172315

![dynamo-export](https://user-images.githubusercontent.com/43064247/103404661-cc05a900-4b97-11eb-8e27-bd6ba98bb040.png)

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
