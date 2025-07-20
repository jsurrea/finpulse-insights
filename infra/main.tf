# Cloud Router
resource "google_compute_router" "nat_router" {
  name    = "truora-nat-router"
  network = google_compute_network.private_network.name
  region  = var.region
}

# Cloud NAT
resource "google_compute_router_nat" "cloud_nat" {
  name                               = "truora-cloud-nat"
  router                             = google_compute_router.nat_router.name
  region                             = var.region
  nat_ip_allocate_option             = "AUTO_ONLY"
  source_subnetwork_ip_ranges_to_nat = "ALL_SUBNETWORKS_ALL_IP_RANGES"
}

# VPC Network for private communication
resource "google_compute_network" "private_network" {
  name                    = "truora-private-network"
  auto_create_subnetworks = false
}

# Subnet for our services
resource "google_compute_subnetwork" "private_subnet" {
  name          = "truora-private-subnet"
  ip_cidr_range = "10.0.0.0/16"
  network       = google_compute_network.private_network.id
  region        = var.region
}

# Firewall rule for SSH access via IAP
resource "google_compute_firewall" "ssh_iap_firewall" {
  name    = "truora-ssh-iap-firewall"
  network = google_compute_network.private_network.name

  allow {
    protocol = "tcp"
    ports    = ["22"]
  }

  source_ranges = ["35.235.240.0/20"]  # IAP IP ranges
  target_tags   = ["cockroachdb"]
}

# Firewall rule for internal communication
resource "google_compute_firewall" "internal_firewall" {
  name    = "truora-internal-firewall"
  network = google_compute_network.private_network.name

  allow {
    protocol = "tcp"
    ports    = ["26257", "8080"]
  }

  source_ranges = ["10.0.0.0/16"]
  target_tags   = ["cockroachdb"]
}

# Cloud Storage for initialization scripts
resource "google_storage_bucket" "init_scripts" {
  name          = "${var.project_id}-init-scripts"
  location      = "US-CENTRAL1"
  force_destroy = true
}

# Upload initialization scripts
resource "google_storage_bucket_object" "init_sql" {
  name   = "init_db.sql"
  bucket = google_storage_bucket.init_scripts.name
  source = "${path.module}/../admin/init_db.sql"
}

resource "google_storage_bucket_object" "load_data_script" {
  name   = "load_data.go"
  bucket = google_storage_bucket.init_scripts.name
  source = "${path.module}/../admin/load_data.go"
}

# CockroachDB instance
resource "google_compute_instance" "cockroachdb" {
  name         = "cockroachdb-instance"
  machine_type = "e2-standard-2"
  zone         = var.zone

  tags = ["cockroachdb"]

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-12"
      size  = 50
    }
  }

  network_interface {
    network    = google_compute_network.private_network.id
    subnetwork = google_compute_subnetwork.private_subnet.id
    
    # No external IP - private only
  }

  service_account {
    scopes = ["cloud-platform"]
  }

  # Inline startup script - no external dependencies
  metadata_startup_script = <<-EOF
    #!/bin/bash

    # Wait for network to be ready
    until ping -c1 8.8.8.8 >/dev/null 2>&1 && curl -s https://www.google.com >/dev/null 2>&1; do
        log "Network not ready. Retrying in 5s..."
        sleep 5
    done

    log "Internet connectivity established"

    
    # Enable strict error handling
    set -euo pipefail
    
    # Setup logging
    LOG_FILE="/var/log/cockroach-init.log"
    exec > >(tee -a "$LOG_FILE")
    exec 2>&1
    
    echo "=== Starting CockroachDB initialization at $(date) ==="
    
    # Function to log with timestamp
    log() {
        echo "[$(date '+%Y-%m-%d %H:%M:%S')] $1"
    }
    
    # Function to retry command
    retry() {
        local max_attempts=$1
        local delay=$2
        local command="$${@:3}"
        local attempt=1
        
        while [ $attempt -le $max_attempts ]; do
            log "Attempt $attempt/$max_attempts: $command"
            if eval "$command"; then
                log "Command succeeded on attempt $attempt"
                return 0
            fi
            
            if [ $attempt -lt $max_attempts ]; then
                log "Command failed, retrying in $delay seconds..."
                sleep $delay
            fi
            
            ((attempt++))
        done
        
        log "ERROR: Command failed after $max_attempts attempts"
        return 1
    }
    
    # Update system packages
    log "Updating system packages..."
    apt-get update -y
    
    # Install Docker
    log "Installing Docker..."
    apt-get install -y docker.io curl
    
    # Start and enable Docker service
    log "Starting Docker service..."
    systemctl start docker
    systemctl enable docker
    
    # Wait for Docker to be ready
    log "Waiting for Docker to be ready..."
    retry 10 3 "docker version"
    
    # Start CockroachDB container
    log "Starting CockroachDB container..."
    docker run -d --name cockroach-single \
        --restart=unless-stopped \
        -p 26257:26257 -p 8080:8080 \
        -v cockroach-data:/cockroach/cockroach-data \
        cockroachdb/cockroach:latest start-single-node --insecure
    
    # Wait for container to be running
    log "Waiting for CockroachDB container to be running..."
    retry 30 5 "docker ps | grep cockroach-single | grep -q 'Up'"
    
    # Wait for CockroachDB to be ready for connections
    log "Waiting for CockroachDB to accept connections..."
    retry 60 5 "docker exec cockroach-single ./cockroach sql --insecure --execute='SELECT 1;'"

    # Define new working directory
    WORKDIR="/opt/cockroach-init"
    mkdir -p "$WORKDIR"
    chown "$(whoami)":"$(whoami)" "$WORKDIR"  # Si se requiere, para permisos
    
    # Download initialization files
    log "Downloading initialization files..."
    retry 5 10 "gsutil cp gs://${google_storage_bucket.init_scripts.name}/init_db.sql $WORKDIR/init_db.sql"
    retry 5 10 "gsutil cp gs://${google_storage_bucket.init_scripts.name}/load_data.go $WORKDIR/load_data.go"
    
    # Execute DDL script
    log "Executing DDL script..."
    cat "$WORKDIR/init_db.sql" | tee -a "$LOG_FILE"

    log "Ejecutando DDL con docker exec usando database=defaultdb..."
    if ! docker exec -i cockroach-single ./cockroach sql --insecure --database=defaultdb < "$WORKDIR/init_db.sql" 2>&1 | tee -a "$LOG_FILE"; then
      log "ERROR: Falló la ejecución del SQL. Verifica sintaxis o base de datos correcta."
      exit 1
    fi


    # Install Go for data loading
    log "Installing Go..."
    apt-get install -y golang-go
    
    # Setup Go environment and load data
    log "Setting up Go environment for data loading..."
    export HOME="/opt/cockroach-init"
    export GOMODCACHE="/opt/cockroach-init/go-mod-cache"
    mkdir -p "$GOMODCACHE"
    export GOPATH="/opt/cockroach-init/go"
    mkdir -p "$GOPATH"
    cd $WORKDIR
    
    # Set environment variables
    export API_TOKEN="${var.api_token}"
    export DATABASE_URL="postgresql://root@localhost:26257/defaultdb?sslmode=disable"
    export GOPROXY=direct
    
    # Initialize Go module
    log "Initializing Go module..."
    go mod init data-loader
    
    # Download dependencies (with retries)
    log "Downloading Go dependencies..."
    retry 5 10 "go mod tidy"
    
    # Test database connection before loading data
    log "Testing database connection..."
    retry 10 5 "docker exec cockroach-single ./cockroach sql --insecure --execute='SHOW DATABASES;'"
    
    # Load data with timeout
    log "Loading data..."
    timeout 600 go run load_data.go || {
        log "ERROR: Data loading failed or timed out"
        exit 1
    }
    
    # Verify data was loaded
    log "Verifying data was loaded..."
    record_count=$(docker exec cockroach-single ./cockroach sql --insecure --execute="SELECT COUNT(*) FROM stock_recommendations;" --format=csv | tail -1)
    log "Data verification successful! Total records: $record_count"
    
    # Final health check
    log "Performing final health check..."
    docker exec cockroach-single ./cockroach sql --insecure --execute="SELECT 'CockroachDB is healthy' AS status;"
    
    log "=== CockroachDB initialization completed successfully at $(date) ==="
    
    # Create success marker
    touch /var/log/cockroach-init-success
    
    log "Initialization complete. CockroachDB is ready for connections."
    
  EOF

  depends_on = [
    google_storage_bucket_object.init_sql,
    google_storage_bucket_object.load_data_script
  ]
}

# VPC Connector for Cloud Run
resource "google_vpc_access_connector" "connector" {
  name          = "truora-vpc-connector"
  ip_cidr_range = "10.8.0.0/28"
  network       = google_compute_network.private_network.name
  region        = var.region
}

# Cloud Run service for API
resource "google_cloud_run_service" "stocks_api" {
  name     = "stocks-api"
  location = var.region

  template {
    metadata {
      annotations = {
        "autoscaling.knative.dev/maxScale"      = "10"
        "run.googleapis.com/vpc-access-connector" = google_vpc_access_connector.connector.id
      }
    }
    
    spec {
      containers {
        image = var.api_image_url
        
        ports {
          container_port = 8080
        }
        
        env {
          name  = "API_TOKEN"
          value = var.api_token
        }
        
        env {
          name  = "DATABASE_URL"
          value = "postgresql://root@${google_compute_instance.cockroachdb.network_interface.0.network_ip}:26257/defaultdb?sslmode=disable"
        }
        
        resources {
          limits = {
            memory = "512Mi"
            cpu    = "1000m"
          }
        }
      }
    }
  }

  depends_on = [google_compute_instance.cockroachdb]
}

# Cloud Run Frontend
resource "google_cloud_run_service" "stocks_web" {
  name     = "stocks-web"
  location = var.region

  template {
    metadata {
      annotations = {
        "autoscaling.knative.dev/maxScale" = "10"
      }
    }
    
    spec {
      containers {
        image = var.web_image_url
        
        ports {
          container_port = 80
        }
        
        env {
          name  = "VITE_API_BASE_URL"
          value = google_cloud_run_service.stocks_api.status[0].url
        }
        
        resources {
          limits = {
            memory = "256Mi"
            cpu    = "500m"
          }
        }
      }
    }
  }

  depends_on = [google_cloud_run_service.stocks_api]
}


# Allow public access to Cloud Run
resource "google_cloud_run_service_iam_member" "api_public_access" {
  service  = google_cloud_run_service.stocks_api.name
  location = google_cloud_run_service.stocks_api.location
  role     = "roles/run.invoker"
  member   = "allUsers"
}

resource "google_cloud_run_service_iam_member" "stocks_web_public_access" {
  service  = google_cloud_run_service.stocks_web.name
  location = google_cloud_run_service.stocks_web.location
  role     = "roles/run.invoker"
  member   = "allUsers"
}

# DNS Managed Zone
resource "google_dns_managed_zone" "main_zone" {
  name     = "finpulseinsights-zone"
  dns_name = "finpulseinsights.com."
  
  description = "Zona DNS para finpulseinsights.com"
}

# Domain Mapping Frontend
resource "google_cloud_run_domain_mapping" "web_domain" {
  location = var.region
  name     = "finpulseinsights.com"

  metadata {
    namespace = var.project_id
  }

  spec {
    route_name = google_cloud_run_service.frontend.name
  }

  depends_on = [google_cloud_run_service.frontend]
}

# Domain Mapping API
resource "google_cloud_run_domain_mapping" "api_domain" {
  location = var.region
  name     = "api.finpulseinsights.com"

  metadata {
    namespace = var.project_id
  }

  spec {
    route_name = google_cloud_run_service.stocks_api.name
  }

  depends_on = [google_cloud_run_service.stocks_api]
}


