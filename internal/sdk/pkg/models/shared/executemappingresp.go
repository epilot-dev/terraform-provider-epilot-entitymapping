// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ExecuteMappingResp struct {
	Failures       []MappingFailure `json:"failures,omitempty"`
	MappedEntities []interface{}    `json:"mapped_entities"`
	Warnings       []MappingWarning `json:"warnings,omitempty"`
}

func (o *ExecuteMappingResp) GetFailures() []MappingFailure {
	if o == nil {
		return nil
	}
	return o.Failures
}

func (o *ExecuteMappingResp) GetMappedEntities() []interface{} {
	if o == nil {
		return []interface{}{}
	}
	return o.MappedEntities
}

func (o *ExecuteMappingResp) GetWarnings() []MappingWarning {
	if o == nil {
		return nil
	}
	return o.Warnings
}
