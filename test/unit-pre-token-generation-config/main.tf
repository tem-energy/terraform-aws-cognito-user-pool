terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "5.100.0"
    }
  }
}

provider "aws" {
  region                      = "us-east-1"
  skip_credentials_validation = true
  skip_metadata_api_check     = true
  skip_requesting_account_id  = true
}

module "test" {
  source = "../../aws-cognito-user-pool"

  name = "pre-token-generation-config-test"

  lambda_pre_token_generation_config = {
    lambda_arn     = "arn:aws:lambda:us-east-1:123456789012:function:pre-token-generation"
    lambda_version = "V2_0"
  }

  lambda_custom_email_sender = {
    lambda_arn     = "arn:aws:lambda:us-east-1:123456789012:function:custom-email-sender"
    lambda_version = "V1_0"
  }
  lambda_kms_key_arn = "arn:aws:kms:us-east-1:123456789012:key/00000000-0000-0000-0000-000000000000"
}
