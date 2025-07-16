variable "project_id" {
  description = "GCP Project ID"
  type        = string
}

variable "region" {
  description = "GCP Region"
  type        = string
  default     = "us-central1"
}

variable "zone" {
  description = "GCP Zone"
  type        = string
  default     = "us-central1-a"
}

variable "image_url" {
  description = "Container image URL"
  type        = string
}

variable "api_token" {
  description = "API token for external service"
  type        = string
  sensitive   = true
}
