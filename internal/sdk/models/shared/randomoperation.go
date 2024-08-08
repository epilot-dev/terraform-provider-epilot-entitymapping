// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/epilot-dev/terraform-provider-epilot-entitymapping/internal/sdk/internal/utils"
)

type RandomOperationType string

const (
	RandomOperationTypeNumber RandomOperationType = "number"
)

func (e RandomOperationType) ToPointer() *RandomOperationType {
	return &e
}
func (e *RandomOperationType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "number":
		*e = RandomOperationType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RandomOperationType: %v", v)
	}
}

type Two struct {
	Max  *float64            `default:"1" json:"max"`
	Min  *float64            `default:"0" json:"min"`
	Type RandomOperationType `json:"type"`
}

func (t Two) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *Two) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *Two) GetMax() *float64 {
	if o == nil {
		return nil
	}
	return o.Max
}

func (o *Two) GetMin() *float64 {
	if o == nil {
		return nil
	}
	return o.Min
}

func (o *Two) GetType() RandomOperationType {
	if o == nil {
		return RandomOperationType("")
	}
	return o.Type
}

type RandomOperationSchemasType string

const (
	RandomOperationSchemasTypeUUID   RandomOperationSchemasType = "uuid"
	RandomOperationSchemasTypeNanoid RandomOperationSchemasType = "nanoid"
)

func (e RandomOperationSchemasType) ToPointer() *RandomOperationSchemasType {
	return &e
}
func (e *RandomOperationSchemasType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "uuid":
		fallthrough
	case "nanoid":
		*e = RandomOperationSchemasType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RandomOperationSchemasType: %v", v)
	}
}

type One struct {
	Type RandomOperationSchemasType `json:"type"`
}

func (o *One) GetType() RandomOperationSchemasType {
	if o == nil {
		return RandomOperationSchemasType("")
	}
	return o.Type
}

type RandomOperationUnionType string

const (
	RandomOperationUnionTypeOne RandomOperationUnionType = "1"
	RandomOperationUnionTypeTwo RandomOperationUnionType = "2"
)

type RandomOperation struct {
	One *One
	Two *Two

	Type RandomOperationUnionType
}

func CreateRandomOperationOne(one One) RandomOperation {
	typ := RandomOperationUnionTypeOne

	return RandomOperation{
		One:  &one,
		Type: typ,
	}
}

func CreateRandomOperationTwo(two Two) RandomOperation {
	typ := RandomOperationUnionTypeTwo

	return RandomOperation{
		Two:  &two,
		Type: typ,
	}
}

func (u *RandomOperation) UnmarshalJSON(data []byte) error {

	var one One = One{}
	if err := utils.UnmarshalJSON(data, &one, "", true, true); err == nil {
		u.One = &one
		u.Type = RandomOperationUnionTypeOne
		return nil
	}

	var two Two = Two{}
	if err := utils.UnmarshalJSON(data, &two, "", true, true); err == nil {
		u.Two = &two
		u.Type = RandomOperationUnionTypeTwo
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for RandomOperation", string(data))
}

func (u RandomOperation) MarshalJSON() ([]byte, error) {
	if u.One != nil {
		return utils.MarshalJSON(u.One, "", true)
	}

	if u.Two != nil {
		return utils.MarshalJSON(u.Two, "", true)
	}

	return nil, errors.New("could not marshal union type RandomOperation: all fields are null")
}
