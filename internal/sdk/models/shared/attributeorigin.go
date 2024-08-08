// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// AttributeOrigin - Origin of an attribute.
type AttributeOrigin string

const (
	AttributeOriginSystemRecommendation               AttributeOrigin = "system_recommendation"
	AttributeOriginUserManually                       AttributeOrigin = "user_manually"
	AttributeOriginEntityUpdatingSystemRecommendation AttributeOrigin = "entity_updating_system_recommendation"
)

func (e AttributeOrigin) ToPointer() *AttributeOrigin {
	return &e
}
func (e *AttributeOrigin) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "system_recommendation":
		fallthrough
	case "user_manually":
		fallthrough
	case "entity_updating_system_recommendation":
		*e = AttributeOrigin(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AttributeOrigin: %v", v)
	}
}
