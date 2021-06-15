# consul agent -dev
# nomad agent -dev
# nomad job run example.nomad

job "example" {
  datacenters = ["dc1"]

  group "hclfmt" {
    network {
      mode = "bridge"
      port "in" {}
    }

    service {
      name = "hclfmt"
      port = "in"
      connect {
        native = true
      }
    }

    task "hclfmt-web" {
      driver = "raw_exec"

      env {
        SERVICE = "hclfmt"
        BIND    = "0.0.0.0"
        PORT    = "${NOMAD_PORT_in}"
      }

      config {
        command = "hclfmt-web"
      }
    }
  }
}
