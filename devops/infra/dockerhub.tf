# Create an image registry
resource "dockerhub_repository" "registry" {
  name        = "scds"
  namespace   = "dlopezlo"
  description = "My corporate container image registry"
}
