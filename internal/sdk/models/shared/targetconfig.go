// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package shared

import (
	"errors"
	"fmt"
	"github.com/epilot-dev/terraform-provider-epilot-entitymapping/internal/sdk/internal/utils"
)

// LoopConfig - contains config in case of running in loop mode
type LoopConfig struct {
	// a hard limit of how many times the loop is allowed to run.
	Length *float64 `default:"the length of the array" json:"length"`
	// path to the array from the entity payload
	SourcePath *string `json:"source_path,omitempty"`
}

func (l LoopConfig) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *LoopConfig) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *LoopConfig) GetLength() *float64 {
	if o == nil {
		return nil
	}
	return o.Length
}

func (o *LoopConfig) GetSourcePath() *string {
	if o == nil {
		return nil
	}
	return o.SourcePath
}

type MappingAttributesType string

const (
	MappingAttributesTypeMappingAttributeV2 MappingAttributesType = "MappingAttributeV2"
	MappingAttributesTypeMappingAttribute   MappingAttributesType = "MappingAttribute"
)

type MappingAttributes struct {
	MappingAttributeV2 *MappingAttributeV2
	MappingAttribute   *MappingAttribute

	Type MappingAttributesType
}

func CreateMappingAttributesMappingAttributeV2(mappingAttributeV2 MappingAttributeV2) MappingAttributes {
	typ := MappingAttributesTypeMappingAttributeV2

	return MappingAttributes{
		MappingAttributeV2: &mappingAttributeV2,
		Type:               typ,
	}
}

func CreateMappingAttributesMappingAttribute(mappingAttribute MappingAttribute) MappingAttributes {
	typ := MappingAttributesTypeMappingAttribute

	return MappingAttributes{
		MappingAttribute: &mappingAttribute,
		Type:             typ,
	}
}

func (u *MappingAttributes) UnmarshalJSON(data []byte) error {

	var mappingAttributeV2 MappingAttributeV2 = MappingAttributeV2{}
	if err := utils.UnmarshalJSON(data, &mappingAttributeV2, "", true, true); err == nil {
		u.MappingAttributeV2 = &mappingAttributeV2
		u.Type = MappingAttributesTypeMappingAttributeV2
		return nil
	}

	var mappingAttribute MappingAttribute = MappingAttribute{}
	if err := utils.UnmarshalJSON(data, &mappingAttribute, "", true, true); err == nil {
		u.MappingAttribute = &mappingAttribute
		u.Type = MappingAttributesTypeMappingAttribute
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for MappingAttributes", string(data))
}

func (u MappingAttributes) MarshalJSON() ([]byte, error) {
	if u.MappingAttributeV2 != nil {
		return utils.MarshalJSON(u.MappingAttributeV2, "", true)
	}

	if u.MappingAttribute != nil {
		return utils.MarshalJSON(u.MappingAttribute, "", true)
	}

	return nil, errors.New("could not marshal union type MappingAttributes: all fields are null")
}

type TargetConfig struct {
	// Pass it as true, when you don't want failures to interrupt the mapping process.
	AllowFailure  *bool `json:"allow_failure,omitempty"`
	ConditionMode any   `json:"conditionMode,omitempty"`
	Conditions    any   `json:"conditions,omitempty"`
	// Identifier for target configuration. Useful for later usages when trying to identify which target config to map to.
	ID *string `json:"id,omitempty"`
	// Relation attribute on the main entity where the target entity will be linked. Set to false to disable linkback
	//
	LinkbackRelationAttribute *string `default:"mapped_entities" json:"linkback_relation_attribute"`
	// Relation tags (labels) to include in main entity linkback relation attribute
	LinkbackRelationTags []string `json:"linkback_relation_tags,omitempty"`
	// contains config in case of running in loop mode
	LoopConfig *LoopConfig `json:"loop_config,omitempty"`
	// Attribute mappings
	MappingAttributes []MappingAttributes `json:"mapping_attributes"`
	// A name for this configuration
	Name *string `json:"name,omitempty"`
	// Relation mappings
	RelationAttributes []RelationAttribute `json:"relation_attributes,omitempty"`
	// Schema of target entity
	TargetSchema string `json:"target_schema"`
	// Unique key for target entity (see upsertEntity of Entity API)
	TargetUnique []string `json:"target_unique,omitempty"`
}

func (t TargetConfig) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TargetConfig) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TargetConfig) GetAllowFailure() *bool {
	if o == nil {
		return nil
	}
	return o.AllowFailure
}

func (o *TargetConfig) GetConditionMode() any {
	if o == nil {
		return nil
	}
	return o.ConditionMode
}

func (o *TargetConfig) GetConditions() any {
	if o == nil {
		return nil
	}
	return o.Conditions
}

func (o *TargetConfig) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TargetConfig) GetLinkbackRelationAttribute() *string {
	if o == nil {
		return nil
	}
	return o.LinkbackRelationAttribute
}

func (o *TargetConfig) GetLinkbackRelationTags() []string {
	if o == nil {
		return nil
	}
	return o.LinkbackRelationTags
}

func (o *TargetConfig) GetLoopConfig() *LoopConfig {
	if o == nil {
		return nil
	}
	return o.LoopConfig
}

func (o *TargetConfig) GetMappingAttributes() []MappingAttributes {
	if o == nil {
		return []MappingAttributes{}
	}
	return o.MappingAttributes
}

func (o *TargetConfig) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *TargetConfig) GetRelationAttributes() []RelationAttribute {
	if o == nil {
		return nil
	}
	return o.RelationAttributes
}

func (o *TargetConfig) GetTargetSchema() string {
	if o == nil {
		return ""
	}
	return o.TargetSchema
}

func (o *TargetConfig) GetTargetUnique() []string {
	if o == nil {
		return nil
	}
	return o.TargetUnique
}