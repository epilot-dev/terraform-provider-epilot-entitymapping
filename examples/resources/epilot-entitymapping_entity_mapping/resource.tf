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
  targets = "{ \"see\": \"documentation\" }"
}