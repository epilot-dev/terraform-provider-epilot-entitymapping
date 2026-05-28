terraform {
  required_providers {
    epilot-entitymapping = {
      source  = "epilot-dev/epilot-entitymapping"
      version = "0.10.0"
    }
  }
}

provider "epilot-entitymapping" {
  epilot_auth = "<YOUR_EPILOT_AUTH>" # Required
  server_url  = "..."                # Optional
}