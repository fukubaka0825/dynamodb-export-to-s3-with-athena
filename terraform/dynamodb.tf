resource "aws_dynamodb_table" "sample" {
  name         = "sample_dynamodb_table"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "id"

  attribute {
    name = "id"
    type = "S"
  }

  attribute {
    name = "name"
    type = "S"
  }

  point_in_time_recovery {
    enabled = true
  }
}
