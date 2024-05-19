provider "digitalocean" {}

resource "digitalocean_kubernetes_cluster" "tfg-prod" {
  name   = "tfg-pro"
  region = "ams3"
  # Grab the latest version slug from `doctl kubernetes options versions`
  version = "1.29.1-do.0"

  node_pool {
    name       = "tfgpro-pool"
    size       = "s-2vcpu-2gb"
    node_count = 2
  }
}
