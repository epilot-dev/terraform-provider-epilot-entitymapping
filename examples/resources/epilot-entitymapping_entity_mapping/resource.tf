resource "epilot-entitymapping_entity_mapping" "my_entitymapping" {
  id = "70542580-2b38-4bfc-af8d-bb90102f9f47"
  source = {
    config = {
      entity_ref = {
        entity_id     = "...my_entity_id..."
        entity_schema = "submission"
      }
      journey_ref = {
        journey_id = "...my_journey_id..."
      }
    }
    type = "journey"
  }
  targets = [
    {
      allow_failure               = true
      condition_mode              = "{ \"see\": \"documentation\" }"
      conditions                  = "{ \"see\": \"documentation\" }"
      id                          = "...my_id..."
      linkback_relation_attribute = "...my_linkback_relation_attribute..."
      linkback_relation_tags = [
        "..."
      ]
      loop_config = {
        length      = 8.78
        source_path = "...my_source_path..."
      }
      mapping_attributes = [
        {
          any = "{ \"see\": \"documentation\" }"
          mapping_attribute = {
            append_value_mapper = {
              mode   = "copy_if_exists"
              source = "...my_source..."
              target = "...my_target..."
              target_unique = [
                "..."
              ]
              value_json = "...my_value_json..."
            }
            copy_value_mapper = {
              mode   = "copy_if_exists"
              source = "...my_source..."
              target = "...my_target..."
            }
            set_value_mapper = {
              mode   = "set_value"
              target = "...my_target..."
              value  = "{ \"see\": \"documentation\" }"
            }
          }
        }
      ]
      name = "...my_name..."
      relation_attributes = [
        {
          mode   = "prepend"
          origin = "entity_updating_system_recommendation"
          source_filter = {
            attribute    = "...my_attribute..."
            limit        = 5
            relation_tag = "...my_relation_tag..."
            schema       = "...my_schema..."
            self         = false
            tag          = "...my_tag..."
          }
          target = "...my_target..."
          target_tags = [
            "..."
          ]
          target_tags_include_source = false
        }
      ]
      target_schema = "...my_target_schema..."
      target_unique = [
        "..."
      ]
    }
  ]
}