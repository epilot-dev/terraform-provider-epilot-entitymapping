// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/epilot-dev/terraform-provider-epilot-entitymapping/internal/sdk/pkg/utils"
)

type ConfigType string

const (
	ConfigTypeJourneyRef ConfigType = "JourneyRef"
	ConfigTypeEntityRef  ConfigType = "EntityRef"
)

type Config struct {
	JourneyRef *JourneyRef
	EntityRef  *EntityRef

	Type ConfigType
}

func CreateConfigJourneyRef(journeyRef JourneyRef) Config {
	typ := ConfigTypeJourneyRef

	return Config{
		JourneyRef: &journeyRef,
		Type:       typ,
	}
}

func CreateConfigEntityRef(entityRef EntityRef) Config {
	typ := ConfigTypeEntityRef

	return Config{
		EntityRef: &entityRef,
		Type:      typ,
	}
}

func (u *Config) UnmarshalJSON(data []byte) error {

	journeyRef := new(JourneyRef)
	if err := utils.UnmarshalJSON(data, &journeyRef, "", true, true); err == nil {
		u.JourneyRef = journeyRef
		u.Type = ConfigTypeJourneyRef
		return nil
	}

	entityRef := new(EntityRef)
	if err := utils.UnmarshalJSON(data, &entityRef, "", true, true); err == nil {
		u.EntityRef = entityRef
		u.Type = ConfigTypeEntityRef
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u Config) MarshalJSON() ([]byte, error) {
	if u.JourneyRef != nil {
		return utils.MarshalJSON(u.JourneyRef, "", true)
	}

	if u.EntityRef != nil {
		return utils.MarshalJSON(u.EntityRef, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type SourceConfigType string

const (
	SourceConfigTypeJourney SourceConfigType = "journey"
	SourceConfigTypeEntity  SourceConfigType = "entity"
)

func (e SourceConfigType) ToPointer() *SourceConfigType {
	return &e
}

func (e *SourceConfigType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "journey":
		fallthrough
	case "entity":
		*e = SourceConfigType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceConfigType: %v", v)
	}
}

type SourceConfig struct {
	Config *Config           `json:"config,omitempty"`
	Type   *SourceConfigType `json:"type,omitempty"`
}

func (o *SourceConfig) GetConfig() *Config {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *SourceConfig) GetType() *SourceConfigType {
	if o == nil {
		return nil
	}
	return o.Type
}
