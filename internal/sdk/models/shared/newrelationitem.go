// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type NewRelationItem struct {
	RelationAttr   string   `json:"relation_attr"`
	SourceEntityID string   `json:"source_entity_id"`
	Tags           []string `json:"tags,omitempty"`
	TargetEntityID string   `json:"target_entity_id"`
}

func (o *NewRelationItem) GetRelationAttr() string {
	if o == nil {
		return ""
	}
	return o.RelationAttr
}

func (o *NewRelationItem) GetSourceEntityID() string {
	if o == nil {
		return ""
	}
	return o.SourceEntityID
}

func (o *NewRelationItem) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *NewRelationItem) GetTargetEntityID() string {
	if o == nil {
		return ""
	}
	return o.TargetEntityID
}
