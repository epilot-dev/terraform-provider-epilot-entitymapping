resource "epilot-entitymapping_entity_mapping" "my_entitymapping" {
  id = "70542580-2b38-4bfc-af8d-bb90102f9f47"
  source = {
    config = {
      entity_ref = {
        entity_id     = "...my_entity_id..."
        entity_schema = "submission"
      }
    }
    type = "entity"
  }
  targets = [
    {
      allow_failure               = true
      condition_mode              = "{ \"see\": \"documentation\" }"
      conditions                  = "{ \"see\": \"documentation\" }"
      id                          = "4d3a547d-7090-41c1-8e07-37604f94fabd"
      linkback_relation_attribute = "...my_linkback_relation_attribute..."
      linkback_relation_tags = [
        "...",
      ]
      loop_config = {
        length      = 89.69
        source_path = "...my_source_path..."
      }
      mapping_attributes = "{ \"see\": \"documentation\" }"
      name               = "Harvey Frami"
      relation_attributes = [
        {
          mode   = "prepend"
          origin = "system_recommendation"
          related_to = {
            "Lead" = "{ \"see\": \"documentation\" }"
            "West" = "{ \"see\": \"documentation\" }"
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