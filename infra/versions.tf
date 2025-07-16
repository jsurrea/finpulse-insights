terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 4.40"
    }
  }
  required_version = ">= 1.3"
}

provider "google" {
  project = var.project_id
  region  = var.region
}
