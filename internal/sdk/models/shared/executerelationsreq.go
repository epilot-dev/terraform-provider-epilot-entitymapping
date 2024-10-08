// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/epilot-dev/terraform-provider-epilot-entitymapping/internal/sdk/internal/utils"
)

// Linkback - For cases where you want to store a relation between main entity (eg. submission) and current source entity.
type Linkback struct {
	// Relation attribute on the main entity (submission) where the target entity will be linked. Set to false to disable linkback
	//
	Attribute *string `default:"mapped_entities" json:"attribute"`
	// Relation tags (labels) to include in main entity linkback relation attribute
	RelationTags []string `json:"relation_tags"`
}

func (l Linkback) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *Linkback) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Linkback) GetAttribute() *string {
	if o == nil {
		return nil
	}
	return o.Attribute
}

func (o *Linkback) GetRelationTags() []string {
	if o == nil {
		return []string{}
	}
	return o.RelationTags
}

type Target struct {
	// For cases where you want to store a relation between main entity (eg. submission) and current source entity.
	Linkback      *Linkback `json:"linkback,omitempty"`
	MainEntityRef EntityRef `json:"main_entity_ref"`
	// Relation mappings
	RelationAttributes []RelationAttribute `json:"relation_attributes"`
}

func (o *Target) GetLinkback() *Linkback {
	if o == nil {
		return nil
	}
	return o.Linkback
}

func (o *Target) GetMainEntityRef() EntityRef {
	if o == nil {
		return EntityRef{}
	}
	return o.MainEntityRef
}

func (o *Target) GetRelationAttributes() []RelationAttribute {
	if o == nil {
		return []RelationAttribute{}
	}
	return o.RelationAttributes
}

// ExecuteRelationsReq - Build relations between a source entity and one or more target entities, dynamically identified
type ExecuteRelationsReq struct {
	AdditionalRelations []RelationItem `json:"additional_relations,omitempty"`
	SourceRef           EntityRef      `json:"source_ref"`
	Target              *Target        `json:"target,omitempty"`
}

func (o *ExecuteRelationsReq) GetAdditionalRelations() []RelationItem {
	if o == nil {
		return nil
	}
	return o.AdditionalRelations
}

func (o *ExecuteRelationsReq) GetSourceRef() EntityRef {
	if o == nil {
		return EntityRef{}
	}
	return o.SourceRef
}

func (o *ExecuteRelationsReq) GetTarget() *Target {
	if o == nil {
		return nil
	}
	return o.Target
}
