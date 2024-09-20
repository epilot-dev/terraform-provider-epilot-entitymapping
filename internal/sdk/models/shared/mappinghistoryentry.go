// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type MappingHistoryEntry struct {
	ID                     string         `json:"id"`
	MappedEntitiesSnapshot []any          `json:"mapped_entities_snapshot"`
	SourceEntitySnapshot   any            `json:"source_entity_snapshot"`
	TargetConfigsSnapshot  []TargetConfig `json:"target_configs_snapshot"`
	Timestamp              string         `json:"timestamp"`
}

func (o *MappingHistoryEntry) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *MappingHistoryEntry) GetMappedEntitiesSnapshot() []any {
	if o == nil {
		return []any{}
	}
	return o.MappedEntitiesSnapshot
}

func (o *MappingHistoryEntry) GetSourceEntitySnapshot() any {
	if o == nil {
		return nil
	}
	return o.SourceEntitySnapshot
}

func (o *MappingHistoryEntry) GetTargetConfigsSnapshot() []TargetConfig {
	if o == nil {
		return []TargetConfig{}
	}
	return o.TargetConfigsSnapshot
}

func (o *MappingHistoryEntry) GetTimestamp() string {
	if o == nil {
		return ""
	}
	return o.Timestamp
}
