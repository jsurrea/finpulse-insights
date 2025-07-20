output "cloud_run_url" {
  description = "URL of the Cloud Run service"
  value       = google_cloud_run_service.stocks_api.status[0].url
}

output "cockroachdb_internal_ip" {
  description = "Internal IP of CockroachDB"
  value       = google_compute_instance.cockroachdb.network_interface.0.network_ip
}

output "database_url" {
  description = "Database connection URL"
  value       = "postgresql://root@${google_compute_instance.cockroachdb.network_interface.0.network_ip}:26257/defaultdb?sslmode=disable"
  sensitive   = true
}

output "stocks_web_url" {
  description = "URL del frontend"
  value       = google_cloud_run_service.stocks_web.status[0].url
}

output "domain_stocks_web" {
  description = "Dominio personalizado del frontend"
  value       = "https://finpulseinsights.com"
}

output "dns_name_servers" {
  description = "Name servers para configurar en el registrador de dominio"
  value       = google_dns_managed_zone.main_zone.name_servers
}

