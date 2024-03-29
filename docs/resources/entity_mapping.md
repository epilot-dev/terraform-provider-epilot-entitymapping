---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "epilot-entitymapping_entity_mapping Resource - terraform-provider-epilot-entitymapping"
subcategory: ""
description: |-
  EntityMapping Resource
---

# epilot-entitymapping_entity_mapping (Resource)

EntityMapping Resource

## Example Usage

```terraform
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
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Mapping Config Id. Requires replacement if changed.
- `source` (Attributes) Requires replacement if changed. (see [below for nested schema](#nestedatt--source))
- `targets` (Attributes List) Requires replacement if changed. (see [below for nested schema](#nestedatt--targets))

<a id="nestedatt--source"></a>
### Nested Schema for `source`

Optional:

- `config` (Attributes) Requires replacement if changed. (see [below for nested schema](#nestedatt--source--config))
- `type` (String) Requires replacement if changed. ; must be one of ["journey", "entity"]

<a id="nestedatt--source--config"></a>
### Nested Schema for `source.config`

Optional:

- `entity_ref` (Attributes) Requires replacement if changed. (see [below for nested schema](#nestedatt--source--config--entity_ref))
- `journey_ref` (Attributes) Requires replacement if changed. (see [below for nested schema](#nestedatt--source--config--journey_ref))

<a id="nestedatt--source--config--entity_ref"></a>
### Nested Schema for `source.config.entity_ref`

Optional:

- `entity_id` (String) id of the source entity to be mapped. Requires replacement if changed. ; Not Null
- `entity_schema` (String) schema of the source entity. Requires replacement if changed.


<a id="nestedatt--source--config--journey_ref"></a>
### Nested Schema for `source.config.journey_ref`

Optional:

- `journey_id` (String) Requires replacement if changed.




<a id="nestedatt--targets"></a>
### Nested Schema for `targets`

Optional:

- `allow_failure` (Boolean) Pass it as true, when you don't want failures to interrupt the mapping process. Requires replacement if changed.
- `condition_mode` (String) Parsed as JSON.
- `conditions` (String) Parsed as JSON.
- `id` (String) Identifier for target configuration. Useful for later usages when trying to identify which target config to map to. Requires replacement if changed.
- `linkback_relation_attribute` (String) Relation attribute on the main entity where the target entity will be linked. Set to false to disable linkback

Requires replacement if changed. ; Default: "mapped_entities"
- `linkback_relation_tags` (List of String) Relation tags (labels) to include in main entity linkback relation attribute. Requires replacement if changed.
- `mapping_attributes` (Attributes List) Attribute mappings. Requires replacement if changed. ; Not Null (see [below for nested schema](#nestedatt--targets--mapping_attributes))
- `name` (String) A name for this configuration. Requires replacement if changed.
- `relation_attributes` (Attributes List) Relation mappings. Requires replacement if changed. (see [below for nested schema](#nestedatt--targets--relation_attributes))
- `target_schema` (String) Schema of target entity. Requires replacement if changed. ; Not Null
- `target_unique` (List of String) Unique key for target entity (see upsertEntity of Entity API). Requires replacement if changed.

<a id="nestedatt--targets--mapping_attributes"></a>
### Nested Schema for `targets.mapping_attributes`

Optional:

- `mapping_attribute` (Attributes) Requires replacement if changed. (see [below for nested schema](#nestedatt--targets--mapping_attributes--mapping_attribute))
- `mapping_attribute_v2` (Attributes) Requires replacement if changed. (see [below for nested schema](#nestedatt--targets--mapping_attributes--mapping_attribute_v2))

<a id="nestedatt--targets--mapping_attributes--mapping_attribute"></a>
### Nested Schema for `targets.mapping_attributes.mapping_attribute`

Optional:

- `append_value_mapper` (Attributes) Requires replacement if changed. (see [below for nested schema](#nestedatt--targets--mapping_attributes--mapping_attribute--append_value_mapper))
- `copy_value_mapper` (Attributes) Requires replacement if changed. (see [below for nested schema](#nestedatt--targets--mapping_attributes--mapping_attribute--copy_value_mapper))
- `set_value_mapper` (Attributes) Requires replacement if changed. (see [below for nested schema](#nestedatt--targets--mapping_attributes--mapping_attribute--set_value_mapper))

<a id="nestedatt--targets--mapping_attributes--mapping_attribute--append_value_mapper"></a>
### Nested Schema for `targets.mapping_attributes.mapping_attribute.set_value_mapper`

Optional:

- `mode` (String) - copy_if_exists - it replaces the target attribute with the source value - append_if_exists - it currently replaces target attribute with array like values. Useful when you have multiple values to be added into one attribute. - set_value - it sets a value to a predefined value. Must be used together with value property.

Requires replacement if changed. ; Not Null; must be one of ["copy_if_exists", "append_if_exists", "set_value"]
- `source` (String) JSON source path for the value to be extracted from the main entity. Eg: steps[1].['Product Info'].price

Requires replacement if changed.
- `target` (String) JSON like target path for the attribute. Eg. last_name. Requires replacement if changed. ; Not Null
- `target_unique` (List of String) Array of keys which should be used when checking for uniqueness. Eg: [country, city, postal_code]

Requires replacement if changed.
- `value_json` (String) To be provided only when mapping json objects into a target attribute. Eg array of addresses.

Requires replacement if changed. ; Not Null


<a id="nestedatt--targets--mapping_attributes--mapping_attribute--copy_value_mapper"></a>
### Nested Schema for `targets.mapping_attributes.mapping_attribute.set_value_mapper`

Optional:

- `mode` (String) - copy_if_exists - it replaces the target attribute with the source value - append_if_exists - it currently replaces target attribute with array like values. Useful when you have multiple values to be added into one attribute. - set_value - it sets a value to a predefined value. Must be used together with value property.

Requires replacement if changed. ; Not Null; must be one of ["copy_if_exists", "append_if_exists", "set_value"]
- `source` (String) JSON source path for the value to be extracted from the main entity. Eg: steps[1].['Product Info'].price

Requires replacement if changed. ; Not Null
- `target` (String) JSON like target path for the attribute. Eg. last_name. Requires replacement if changed. ; Not Null


<a id="nestedatt--targets--mapping_attributes--mapping_attribute--set_value_mapper"></a>
### Nested Schema for `targets.mapping_attributes.mapping_attribute.set_value_mapper`

Optional:

- `mode` (String) - copy_if_exists - it replaces the target attribute with the source value - append_if_exists - it currently replaces target attribute with array like values. Useful when you have multiple values to be added into one attribute. - set_value - it sets a value to a predefined value. Must be used together with value property.

Requires replacement if changed. ; Not Null; must be one of ["copy_if_exists", "append_if_exists", "set_value"]
- `target` (String) JSON like target path for the attribute. Eg. last_name. Requires replacement if changed. ; Not Null
- `value` (String) Any value to be set: string, number, string[], number[], JSON object, etc. It will override existing values, if any.

Parsed as JSON.



<a id="nestedatt--targets--mapping_attributes--mapping_attribute_v2"></a>
### Nested Schema for `targets.mapping_attributes.mapping_attribute_v2`

Optional:

- `operation` (Attributes) Mapping operation nodes are either primitive values or operation node objects. Requires replacement if changed. ; Not Null (see [below for nested schema](#nestedatt--targets--mapping_attributes--mapping_attribute_v2--operation))
- `origin` (String) Origin of an attribute. Requires replacement if changed. ; must be one of ["system_recommendation", "user_manually"]
- `target` (String) Target JSON path for the attribute to set. Requires replacement if changed. ; Not Null

<a id="nestedatt--targets--mapping_attributes--mapping_attribute_v2--operation"></a>
### Nested Schema for `targets.mapping_attributes.mapping_attribute_v2.target`

Optional:

- `any` (String) Parsed as JSON.
- `operation_object_node` (Attributes) Requires replacement if changed. (see [below for nested schema](#nestedatt--targets--mapping_attributes--mapping_attribute_v2--target--operation_object_node))

<a id="nestedatt--targets--mapping_attributes--mapping_attribute_v2--target--operation_object_node"></a>
### Nested Schema for `targets.mapping_attributes.mapping_attribute_v2.target.operation_object_node`

Optional:

- `additional_properties` (String) Parsed as JSON.
- `append` (List of String) Append to array. Requires replacement if changed.
- `copy` (String) Copy JSONPath value from source entity context. Requires replacement if changed.
- `random` (Attributes) Requires replacement if changed. (see [below for nested schema](#nestedatt--targets--mapping_attributes--mapping_attribute_v2--target--operation_object_node--random))
- `set` (String) Parsed as JSON.
- `template` (String) Define handlebars template to output a string. Requires replacement if changed.
- `uniq` (Attributes) Unique array. Requires replacement if changed. (see [below for nested schema](#nestedatt--targets--mapping_attributes--mapping_attribute_v2--target--operation_object_node--uniq))

<a id="nestedatt--targets--mapping_attributes--mapping_attribute_v2--target--operation_object_node--random"></a>
### Nested Schema for `targets.mapping_attributes.mapping_attribute_v2.target.operation_object_node.uniq`

Optional:

- `one` (Attributes) Requires replacement if changed. (see [below for nested schema](#nestedatt--targets--mapping_attributes--mapping_attribute_v2--target--operation_object_node--uniq--one))
- `two` (Attributes) Requires replacement if changed. (see [below for nested schema](#nestedatt--targets--mapping_attributes--mapping_attribute_v2--target--operation_object_node--uniq--two))

<a id="nestedatt--targets--mapping_attributes--mapping_attribute_v2--target--operation_object_node--uniq--one"></a>
### Nested Schema for `targets.mapping_attributes.mapping_attribute_v2.target.operation_object_node.uniq.two`

Optional:

- `type` (String) Requires replacement if changed. ; Not Null; must be one of ["uuid", "nanoid"]


<a id="nestedatt--targets--mapping_attributes--mapping_attribute_v2--target--operation_object_node--uniq--two"></a>
### Nested Schema for `targets.mapping_attributes.mapping_attribute_v2.target.operation_object_node.uniq.two`

Optional:

- `max` (Number) Requires replacement if changed. ; Default: 1
- `min` (Number) Requires replacement if changed. ; Default: 0
- `type` (String) Requires replacement if changed. ; Not Null; must be one of ["number"]



<a id="nestedatt--targets--mapping_attributes--mapping_attribute_v2--target--operation_object_node--uniq"></a>
### Nested Schema for `targets.mapping_attributes.mapping_attribute_v2.target.operation_object_node.uniq`

Optional:

- `array_ofstr` (List of String) Requires replacement if changed.
- `boolean` (Boolean) Requires replacement if changed.






<a id="nestedatt--targets--relation_attributes"></a>
### Nested Schema for `targets.relation_attributes`

Optional:

- `mode` (String) Requires replacement if changed. ; Not Null; must be one of ["append", "prepend", "set"]
- `origin` (String) Origin of an attribute. Requires replacement if changed. ; must be one of ["system_recommendation", "user_manually"]
- `related_to` (Map of String) Requires replacement if changed.
- `source_filter` (Attributes) A filter to identify which source entities to pick as relations from main entity. Requires replacement if changed. (see [below for nested schema](#nestedatt--targets--relation_attributes--source_filter))
- `target` (String) Target attribute to store the relation in. Requires replacement if changed. ; Not Null
- `target_tags` (List of String) Relation tags (labels) to set for the stored relations. Requires replacement if changed.
- `target_tags_include_source` (Boolean) Include all relation tags (labels) present on the main entity relation. Requires replacement if changed. ; Default: false

<a id="nestedatt--targets--relation_attributes--source_filter"></a>
### Nested Schema for `targets.relation_attributes.source_filter`

Optional:

- `attribute` (String) Filter by a specific relation attribute on the main entity. Requires replacement if changed.
- `limit` (Number) Limit relations to maximum number (default, all matched relations). Requires replacement if changed.
- `relation_tag` (String) Filter by relation tag (label) on the main entity. Requires replacement if changed.
- `schema` (String) Filter by specific schema. Requires replacement if changed.
- `self` (Boolean) Picks main entity as relation (overrides other filters). Requires replacement if changed. ; Default: false
- `tag` (String) Filter by a specific tag on the related entity. Requires replacement if changed.


