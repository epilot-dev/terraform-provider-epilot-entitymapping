// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type MappingWarning struct {
	Context     *string `json:"context,omitempty"`
	Explanation string  `json:"explanation"`
	ID          *string `json:"id,omitempty"`
}

func (o *MappingWarning) GetContext() *string {
	if o == nil {
		return nil
	}
	return o.Context
}

func (o *MappingWarning) GetExplanation() string {
	if o == nil {
		return ""
	}
	return o.Explanation
}

func (o *MappingWarning) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}
