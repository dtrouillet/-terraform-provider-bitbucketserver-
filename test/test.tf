terraform {
  required_providers {
    bitbucketserver = {
      source = "dtrouillet/bitbucketserver"
      version = "0.0.2"
  }
  }
}

provider "bitbucketserver" {
  password = "admin"
  server   = "http://localhost:7990"
  username = "admin"
}

resource "bitbucketserver_project" "test" {
  key  = "TEST"
  name = "test"
}

resource "bitbucketserver_repository" "test1" {
  name    = "test1"
  project = bitbucketserver_project.test.key
}

resource "bitbucketserver_repository" "test2" {
  name    = "test2"
  project = bitbucketserver_project.test.key
}

data "bitbucketserver_project" "test" {
  depends_on = [bitbucketserver_project.test]
  key = "TEST"
}

output "repos" {
  value = data.bitbucketserver_project.test
}
