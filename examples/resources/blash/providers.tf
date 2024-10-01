# Provider definitions
terraform {
  required_providers {
    epilot-journey = {
      source  = "epilot-dev/epilot-journey"
      version = "1.0.8"
    }
    epilot-product = {
      source  = "epilot-dev/epilot-product"
      version = "0.10.3"
    }
    epilot-automation = {
      source  = "epilot-dev/epilot-automation"
      version = "2.0.1"
    }
    epilot-entitymapping = {
      source  = "epilot-dev/epilot-entitymapping"
      version = "1.7.4"
    }
    epilot-designbuilder = {
      source  = "epilot-dev/epilot-designbuilder"
      version = "1.2.0"
    }
    epilot-file = {
      source  = "epilot-dev/epilot-file"
      version = "3.2.1"
    }
    epilot-emailtemplate = {
      source  = "epilot-dev/epilot-emailtemplate"
      version = "0.1.4"
    }
    epilot-schema = {
      source  = "epilot-dev/epilot-schema"
      version = "4.2.2"
    }
    epilot-workflow = {
      source  = "epilot-dev/epilot-workflow"
      version = "1.1.1"
    }
    epilot-taxonomy = {
      source  = "epilot-dev/epilot-taxonomy"
      version = "1.0.1"
    }
    epilot-webhook = {
      source  = "epilot-dev/epilot-webhook"
      version = "0.1.0"
    }
  }
}

# Variables
variable "epilot_auth" {
  type = string
}

variable "manifest_id" {
  type = string
  default = ""
}

variable "journey_api_url" {
  type    = string
  default = "https://journey-config.sls.epilot.io"
}
variable "product_api_url" {
  type    = string
  default = "https://product.sls.epilot.io"
}
variable "automation_api_url" {
  type    = string
  default = "https://automation.sls.epilot.io"
}
variable "file_api_url" {
  type    = string
  default = "https://file.sls.epilot.io"
}
variable "emailtemplate_api_url" {
  type    = string
  default = "https://email-template.sls.epilot.io"
}
variable "designbuilder_api_url" {
  type    = string
  default = "https://design-builder-api.sls.epilot.io"
}
variable "entitymapping_api_url" {
  type    = string
  default = "https://entity-mapping.sls.epilot.io"
}
variable "schema_api_url" {
  type    = string
  default = "https://entity.sls.epilot.io"
}
variable "workflow_api_url" {
  type    = string
  default = "https://workflows-definition.sls.epilot.io"
}
variable "webhook_api_url" {
  type    = string
  default = "https://webhooks.sls.epilot.io"
}

# Providers configuration
provider "epilot-journey" {
  epilot_auth = var.epilot_auth
  server_url  = var.journey_api_url
}
provider "epilot-product" {
  epilot_auth = var.epilot_auth
  server_url  = var.product_api_url
}
provider "epilot-automation" {
  epilot_auth = var.epilot_auth
  server_url  = var.automation_api_url
}
provider "epilot-file" {
  epilot_auth = var.epilot_auth
  server_url  = var.file_api_url
}
provider "epilot-emailtemplate" {
  epilot_auth = var.epilot_auth
  server_url  = var.emailtemplate_api_url
}
provider "epilot-designbuilder" {
  custom_authorizer = var.epilot_auth
  server_url        = var.designbuilder_api_url
}
provider "epilot-entitymapping" {
  epilot_auth = var.epilot_auth
  server_url  = var.entitymapping_api_url
}
provider "epilot-schema" {
  epilot_auth = var.epilot_auth
  server_url  = var.schema_api_url
}
provider "epilot-workflow" {
  bearer_auth = var.epilot_auth
  server_url  = var.workflow_api_url
}
provider "epilot-taxonomy" {
  epilot_auth = var.epilot_auth
  server_url  = var.schema_api_url
}
provider "epilot-webhook" {
  epilot_auth = var.epilot_auth
  server_url= var.webhook_api_url
}
