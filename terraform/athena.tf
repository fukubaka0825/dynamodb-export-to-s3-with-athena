resource "aws_athena_database" "sample" {
  name          = "sample_db"
  bucket        = aws_s3_bucket.sample_export.bucket
  force_destroy = false
  encryption_configuration {
    encryption_option = "SSE_KMS"
    kms_key           = "arn:aws:kms:ap-northeast-1:99999999999999:key/xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  }
}

resource "aws_athena_workgroup" "sample" {
  name = "sample"

  configuration {
    enforce_workgroup_configuration    = true
    publish_cloudwatch_metrics_enabled = true
    result_configuration {
      output_location = "s3://${aws_s3_bucket.sample_query_result.bucket}/output/"
      encryption_configuration {
        encryption_option = "SSE_KMS"
        kms_key_arn       = "arn:aws:kms:ap-northeast-1:99999999999999:key/xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
      }
    }
  }
}
