variables {
  dockerhub_username = ""
  dockerhub_password = ""
}

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

  post-processors {
    post-processor "docker-tag" {
      repository = "${var.dockerhub_username}/go-calc"
      tag = ["latest"]
    }
    post-processor "docker-push" {
      login = true
      login_username = "${var.dockerhub_username}"
      login_password = "${var.dockerhub_password}"
    }
  }
}
