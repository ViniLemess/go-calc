
source "docker" "go-calc" {
  commit  = true
  image   = "golang:1.18.3"
  changes = [
    "EXPOSE 8080",
    "ENTRYPOINT /home/go-calc"
  ]
}

build {
  sources = ["source.docker.go-calc"]

  provisioner "file" {
    source = "go-calc/go-calc"
    destination = "/home/go-calc"
  }

  post-processor "docker-tag" {
    repository = "go-calc"
    tag = ["latest"]
  }
}
