resource "epilot-entitymapping_entity_mapping" "my_entitymapping" {
  id     = "20fec022-fdfb-42b1-b545-9c73d2673f99"
  org_id = "...my_org_id..."
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
  version = 34.53
}