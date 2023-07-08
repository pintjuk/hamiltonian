terraform {
  backend "gcs" {
    bucket = "pintjuk-routemaster-tf-state"
    prefix  = "terraform/state"
  }
}

provider "google" {
  project     = "pintjuk-routemaster"
  region      = "europe-north1"
}

variable "service_version" {
  type = string
  default = "v0.0.0"
}
resource "google_cloud_run_v2_service" "default" {
  name     = "routemaster"
  location = "europe-north1"
  ingress = "INGRESS_TRAFFIC_ALL"

  template {
    containers {
      image = "docker.io/pintjuk/routemaster:${var.service_version}"
    }

  }
}

resource "google_cloud_run_service_iam_member" "noauth" {
  service     = google_cloud_run_v2_service.default.name
  location    = google_cloud_run_v2_service.default.location
  role        = "roles/run.invoker"
  member      = "allUsers"
}

output "url" {
  value     = "${google_cloud_run_v2_service.default.uri}"
  sensitive = false
}