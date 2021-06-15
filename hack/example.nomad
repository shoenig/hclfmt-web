# consul agent -dev
# nomad agent -dev
# nomad job run example.nomad

job "example" {
  datacenters = ["dc1"]

  group "hclfmt" {
    network {
      mode = "bridge"
    }

    service {
      name = "hclfmt"
      port = "9100"

      connect {
	native = true
      }
    }

    task "hclfmt-web" {
      driver = "raw_exec"

      config {
	command = "hclfmt-web"
      }
    }
  }
}
