// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type RelationItem struct {
	Attribute string `json:"attribute"`
	EntityID  string `json:"entity_id"`
}

func (o *RelationItem) GetAttribute() string {
	if o == nil {
		return ""
	}
	return o.Attribute
}

func (o *RelationItem) GetEntityID() string {
	if o == nil {
		return ""
	}
	return o.EntityID
}