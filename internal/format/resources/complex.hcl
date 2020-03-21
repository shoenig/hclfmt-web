# This is a valid but poorly formed HCL file.
# Always use hclfmt.

      job   "example"   {
    datacenters   =   [
      "dc1"
    ]
          type = "service"



        update {
          max_parallel = 1

        min_healthy_time = "10s"


        healthy_deadline = "3m"
        progress_deadline = "10m"

          auto_revert = false
        canary = 0

      }
      }


