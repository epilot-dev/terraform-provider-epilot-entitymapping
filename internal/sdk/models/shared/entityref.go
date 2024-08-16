// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type EntityRef struct {
	// id of the source entity to be mapped
	EntityID string `json:"entity_id"`
	// schema of the source entity
	EntitySchema *string `json:"entity_schema,omitempty"`
}

func (o *EntityRef) GetEntityID() string {
	if o == nil {
		return ""
	}
	return o.EntityID
}

func (o *EntityRef) GetEntitySchema() *string {
	if o == nil {
		return nil
	}
	return o.EntitySchema
}
