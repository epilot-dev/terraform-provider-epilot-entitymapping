// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SearchMappingReq struct {
	Source *SourceConfig `json:"source,omitempty"`
}

func (o *SearchMappingReq) GetSource() *SourceConfig {
	if o == nil {
		return nil
	}
	return o.Source
}
