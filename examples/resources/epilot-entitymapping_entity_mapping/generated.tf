# __generated__ by Terraform
# Please review these resources and move them into your main configuration files.

# __generated__ by Terraform
terraform {
  required_providers {
    epilot-entitymapping = {
      source  = "epilot-dev/epilot-entitymapping"
      version = "0.4.3"
    }
  }
}

provider "epilot-entitymapping" {
  # Configuration options
  epilot_auth = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IjRaRnpITVkyVjJyMXRjcW15bEJUaVozSkRSR3pKVW5JeG5Rcm9rQVNGOUEifQ.eyJ0b2tlbl91c2UiOiJhY2Nlc3MiLCJ0b2tlbl9pZCI6ImFwaV9telpFVGNuY2c4aTlzdVBrbUdzcFAiLCJ0b2tlbl9uYW1lIjoiSm91cm5leSBUZXJyYWZvcm0iLCJvcmdfaWQiOiI2NiIsInVzZXJfaWQiOiJhcGlfbXpaRVRjbmNnOGk5c3VQa21Hc3BQIiwidG9rZW5fdHlwZSI6ImFwaSIsImN1c3RvbTppdnlfb3JnX2lkIjoiNjYiLCJjdXN0b206aXZ5X3VzZXJfaWQiOiJhcGlfbXpaRVRjbmNnOGk5c3VQa21Hc3BQIiwiYXNzdW1lX3JvbGVzIjpbIjY2OmFkbWluaXN0cmF0b3IiXSwiaXNzIjoiaHR0cHM6Ly9hY2Nlc3MtdG9rZW4uc2xzLmVwaWxvdC5pby92MS9hY2Nlc3MtdG9rZW5zIiwiaWF0IjoxNzA1MzI5MDk4fQ.KOx25TuzWb3DFmfpOeuVHVeUt_8GGB4UAoBsyE9Vyzh61VvretqBkNX2hZqw_Ac3OYfKnCCEXvkfEEAfshTB6xIWkvAin0guxJHLKDXYJN6Tz_qH663MZgEZcgftZgLOUzlVdMrWaxXrgpq39SK6ML1J3J2UxLdpoAX-NW6WR_qfqXGOqCdrN0x9lXnkXYoziqKVMB06N7VJWavD7zE2YhTlS7Rwcj4mzQ6-1V-QCu4m9knY06JcnJOeBqqTLmGt3F8RGJpkIo4dqFVSKMfxu92bibfSHpld_IjWyhOF6SK8RG7QRilsSBcg4qTUynE60DQmiihnVdECWJlLcbdMhQ"
}


resource "epilot-entitymapping_entity_mapping" "my_entitymapping" {
  id = "1deca750-e078-11ee-8ecf-21b7dd827de6"
  source = {
    config = {
      entity_ref = null
      journey_ref = {
        journey_id = "1d53e9c0-e078-11ee-8ecf-21b7dd827de6"
      }
    }
    type = "journey"
  }
  targets = [
    {
      allow_failure               = null
      condition_mode              = null
      conditions                  = null
      linkback_relation_attribute = "mapped_entities"
      linkback_relation_tags      = null
      mapping_attributes = [
        {
          mapping_attribute = null
          mapping_attribute_v2 = {
            operation = {
              any                   = "{\"_copy\":[\"journey_context.opportunity_id\",\"opportunity._id\"]}"
              operation_object_node = null
            }
            origin = "system_recommendation"
            target = "_id"
          }
        },
        {
          mapping_attribute = null
          mapping_attribute_v2 = {
            operation = {
              any = null
              operation_object_node = {
                additional_properties = null
                append                = null
                copy                  = null
                random                = null
                set                   = "\"Biggoooo\""
                template              = null
                uniq                  = null
              }
            }
            origin = "system_recommendation"
            target = "opportunity_title"
          }
        },
        {
          mapping_attribute = null
          mapping_attribute_v2 = {
            operation = {
              any = null
              operation_object_node = {
                additional_properties = null
                append                = null
                copy                  = "submission.source"
                random                = null
                set                   = null
                template              = null
                uniq                  = null
              }
            }
            origin = "system_recommendation"
            target = "source"
          }
        },
        {
          mapping_attribute = null
          mapping_attribute_v2 = {
            operation = {
              any                   = "{\"_append\":{\"_copy\":[\"order._tags\"]},\"_uniq\":true}"
              operation_object_node = null
            }
            origin = "system_recommendation"
            target = "_tags"
          }
        },
        {
          mapping_attribute = null
          mapping_attribute_v2 = {
            operation = {
              any = null
              operation_object_node = {
                additional_properties = "{\"$relation\":{\"_append\":{\"_copy\":[\"order.products.$relation\"]},\"_uniq\":[\"entity_id\"]}}"
                append                = null
                copy                  = null
                random                = null
                set                   = null
                template              = null
                uniq                  = null
              }
            }
            origin = "system_recommendation"
            target = "products"
          }
        },
        {
          mapping_attribute = null
          mapping_attribute_v2 = {
            operation = {
              any = null
              operation_object_node = {
                additional_properties = "{\"$relation\":{\"_append\":{\"_copy\":[\"order.prices.$relation\"]},\"_uniq\":[\"entity_id\"]}}"
                append                = null
                copy                  = null
                random                = null
                set                   = null
                template              = null
                uniq                  = null
              }
            }
            origin = "system_recommendation"
            target = "prices"
          }
        },
        {
          mapping_attribute = null
          mapping_attribute_v2 = {
            operation = {
              any = null
              operation_object_node = {
                additional_properties = null
                append                = ["{\"entity_id\":{\"_copy\":\"submission.journey_context.relation_id\"}}"]
                copy                  = null
                random                = null
                set                   = null
                template              = null
                uniq = {
                  array_ofstr = ["entity_id"]
                  boolean     = null
                }
              }
            }
            origin = "system_recommendation"
            target = "customer.$relation"
          }
        },
        {
          mapping_attribute = null
          mapping_attribute_v2 = {
            operation = {
              any = null
              operation_object_node = {
                additional_properties = null
                append                = null
                copy                  = null
                random                = null
                set                   = null
                template              = null
                uniq                  = null
              }
            }
            origin = "system_recommendation"
            target = "journey_data"
          }
        },
      ]
      name = "Opportunity from Journey"
      relation_attributes = [
        {
          mode       = "append"
          origin     = null
          related_to = null
          source_filter = {
            attribute    = null
            limit        = null
            relation_tag = null
            schema       = "order"
            self         = false
            tag          = null
          }
          target                     = "items"
          target_tags                = null
          target_tags_include_source = true
        },
        {
          mode       = "append"
          origin     = null
          related_to = null
          source_filter = {
            attribute    = null
            limit        = 1
            relation_tag = "customer"
            schema       = "contact"
            self         = false
            tag          = null
          }
          target                     = "customer"
          target_tags                = null
          target_tags_include_source = false
        },
        {
          mode       = "append"
          origin     = null
          related_to = null
          source_filter = {
            attribute    = null
            limit        = null
            relation_tag = null
            schema       = "file"
            self         = false
            tag          = null
          }
          target                     = "_files"
          target_tags                = null
          target_tags_include_source = true
        },
      ]
      target_schema = "opportunity"
      target_unique = ["_id"]
    },
  ]
}
