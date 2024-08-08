// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type ExecuteMappingResp struct {
	Failures       []MappingFailure `json:"failures,omitempty"`
	MappedEntities []any            `json:"mapped_entities"`
	Warnings       []MappingWarning `json:"warnings,omitempty"`
}

func (o *ExecuteMappingResp) GetFailures() []MappingFailure {
	if o == nil {
		return nil
	}
	return o.Failures
}

func (o *ExecuteMappingResp) GetMappedEntities() []any {
	if o == nil {
		return []any{}
	}
	return o.MappedEntities
}

func (o *ExecuteMappingResp) GetWarnings() []MappingWarning {
	if o == nil {
		return nil
	}
	return o.Warnings
}
