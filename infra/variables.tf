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

variable "api_image_url" {
  description = "Container image URL for API service"
  type        = string
}

variable "web_image_url" {
  description = "Container image URL for web service"
  type        = string
}

variable "api_token" {
  description = "API token for external service"
  type        = string
  sensitive   = true
}
