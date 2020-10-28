terraform {
  backend "s3" {
    bucket = "new-architectures-devops"
    key    = "terraform/new-architectures.tfstate"
    region = "us-west-2"
    profile = "uniandes"
  }
}
