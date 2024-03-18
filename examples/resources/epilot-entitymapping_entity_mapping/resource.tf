terraform {
  required_providers {
    epilot-entitymapping = {
      source  = "epilot-dev/epilot-entitymapping"
      version = "0.6.3"
    }
  }
}

provider "epilot-entitymapping" {
  epilot_auth = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IjRaRnpITVkyVjJyMXRjcW15bEJUaVozSkRSR3pKVW5JeG5Rcm9rQVNGOUEifQ.eyJ0b2tlbl91c2UiOiJhY2Nlc3MiLCJ0b2tlbl9pZCI6ImFwaV9telpFVGNuY2c4aTlzdVBrbUdzcFAiLCJ0b2tlbl9uYW1lIjoiSm91cm5leSBUZXJyYWZvcm0iLCJvcmdfaWQiOiI2NiIsInVzZXJfaWQiOiJhcGlfbXpaRVRjbmNnOGk5c3VQa21Hc3BQIiwidG9rZW5fdHlwZSI6ImFwaSIsImN1c3RvbTppdnlfb3JnX2lkIjoiNjYiLCJjdXN0b206aXZ5X3VzZXJfaWQiOiJhcGlfbXpaRVRjbmNnOGk5c3VQa21Hc3BQIiwiYXNzdW1lX3JvbGVzIjpbIjY2OmFkbWluaXN0cmF0b3IiXSwiaXNzIjoiaHR0cHM6Ly9hY2Nlc3MtdG9rZW4uc2xzLmVwaWxvdC5pby92MS9hY2Nlc3MtdG9rZW5zIiwiaWF0IjoxNzA1MzI5MDk4fQ.KOx25TuzWb3DFmfpOeuVHVeUt_8GGB4UAoBsyE9Vyzh61VvretqBkNX2hZqw_Ac3OYfKnCCEXvkfEEAfshTB6xIWkvAin0guxJHLKDXYJN6Tz_qH663MZgEZcgftZgLOUzlVdMrWaxXrgpq39SK6ML1J3J2UxLdpoAX-NW6WR_qfqXGOqCdrN0x9lXnkXYoziqKVMB06N7VJWavD7zE2YhTlS7Rwcj4mzQ6-1V-QCu4m9knY06JcnJOeBqqTLmGt3F8RGJpkIo4dqFVSKMfxu92bibfSHpld_IjWyhOF6SK8RG7QRilsSBcg4qTUynE60DQmiihnVdECWJlLcbdMhQ"
}
resource "epilot-entitymapping_entity_mapping" "my_entitymapping" {
  id = "20fec022-fdfb-42b1-b545-9c73d2673f99"
  source = {
    config = {
      entity_ref = {
        entity_id     = "...my_entity_id..."
        entity_schema = "submission"
      }
    }
    type = "journey"
  }
  targets = [
    {
      allow_failure               = false
      condition_mode              = "{ \"see\": \"documentation\" }"
      conditions                  = "{ \"see\": \"documentation\" }"
      id                          = "cd4d3a54-7d70-4901-810e-0737604f94fa"
      linkback_relation_attribute = "...my_linkback_relation_attribute..."
      linkback_relation_tags = [
        "...",
      ]
      mapping_attributes = [
        {
          mapping_attribute = {
            append_value_mapper = {
              mode   = "set_value"
              source = "...my_source..."
              target = "...my_target..."
              target_unique = [
                "...",
              ]
              value_json = "...my_value_json..."
            }
          }
        },
      ]
      name = "Terence Wehner"
      relation_attributes = [
        {
          mode   = "append"
          origin = "system_recommendation"
          related_to = {
            "Future" = "{ \"see\": \"documentation\" }"
            "Pickup" = "{ \"see\": \"documentation\" }"
          }
          source_filter = {
            attribute    = "...my_attribute..."
            limit        = 4
            relation_tag = "...my_relation_tag..."
            schema       = "...my_schema..."
            self         = true
            tag          = "...my_tag..."
          }
          target = "...my_target..."
          target_tags = [
            "...",
          ]
          target_tags_include_source = false
        },
      ]
      target_schema = "...my_target_schema..."
      target_unique = [
        "...",
      ]
    },
  ]
}