// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package shared

type MappingConfig struct {
	ID      string         `json:"id"`
	OrgID   string         `json:"org_id"`
	Source  SourceConfig   `json:"source"`
	Targets []TargetConfig `json:"targets"`
	Version int64          `json:"version"`
}

func (o *MappingConfig) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *MappingConfig) GetOrgID() string {
	if o == nil {
		return ""
	}
	return o.OrgID
}

func (o *MappingConfig) GetSource() SourceConfig {
	if o == nil {
		return SourceConfig{}
	}
	return o.Source
}

func (o *MappingConfig) GetTargets() []TargetConfig {
	if o == nil {
		return []TargetConfig{}
	}
	return o.Targets
}

func (o *MappingConfig) GetVersion() int64 {
	if o == nil {
		return 0
	}
	return o.Version
}