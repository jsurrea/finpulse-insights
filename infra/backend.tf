terraform {
  backend "gcs" {
    bucket = "truora-stocks-tf-backend"
    prefix = "terraform/state"
  }
}
