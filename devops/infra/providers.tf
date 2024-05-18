terraform {
    required_version = ">= 0.13"

    required_providers {
        dockerhub = {
            source = "BarnabyShearer/dockerhub"
            version = ">= 0.0.15"
        }
        digitalocean = {
            source = "digitalocean/digitalocean"
            version = "~> 2.0"
        }
    }
}
