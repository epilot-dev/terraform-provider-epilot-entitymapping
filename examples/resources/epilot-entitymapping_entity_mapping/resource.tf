terraform {
  required_providers {
    epilot-entitymapping = {
      source  = "epilot-dev/epilot-entitymapping"
      version = "0.4.1"
    }
  }
}


variable "epilot_auth" {
  type = string
}

provider "epilot-entitymapping" {
  # Configuration options
  epilot_auth = var.epilot_auth
}


resource "epilot-entitymapping_entity_mapping" "new_mapping" {
    # Configuration options
}