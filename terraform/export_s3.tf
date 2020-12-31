resource "aws_s3_bucket" "sample_export" {
  bucket        = "sample-export"
  acl           = "private"
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
        kms_master_key_id = "arn:aws:kms:ap-northeast-1:99999999999999:key/xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
        sse_algorithm     = "aws:kms"
      }
    }
  }
}


# enforce public access block on bucket
# reapply on next tf apply if disabled
resource "aws_s3_bucket_public_access_block" "sample_export" {
  bucket = aws_s3_bucket.sample_export.bucket

  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}