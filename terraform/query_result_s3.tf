resource "aws_s3_bucket" "sample_query_result" {
  bucket        = "sample-result-bucket"
  acl           = "log-delivery-write"
  force_destroy = "false"

  versioning {
    enabled = true
  }

  lifecycle_rule {
    enabled = true
    expiration {
      days = 90
    }
    noncurrent_version_expiration {
      days = 1
    }
  }

  server_side_encryption_configuration {
    rule {
      apply_server_side_encryption_by_default {
        kms_master_key_id = ""
        sse_algorithm     = "aws:kms"
      }
    }
  }
}


# enforce public access block on bucket
# reapply on next tf apply if disabled
resource "aws_s3_bucket_public_access_block" "sample_query_result" {
  bucket = aws_s3_bucket.sample_query_result.bucket

  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}