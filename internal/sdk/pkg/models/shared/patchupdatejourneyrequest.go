// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
)

// PatchUpdateJourneyRequest - Patch request to update a journey (journey id is required) Support for nested properties (e.g. steps[0].uischema.elements[0].products) is supported.
type PatchUpdateJourneyRequest struct {
	JourneyID string `json:"journeyId"`

	AdditionalProperties map[string]interface{} `json:"-"`
}
type _PatchUpdateJourneyRequest PatchUpdateJourneyRequest

func (c *PatchUpdateJourneyRequest) UnmarshalJSON(bs []byte) error {
	data := _PatchUpdateJourneyRequest{}

	if err := json.Unmarshal(bs, &data); err != nil {
		return err
	}
	*c = PatchUpdateJourneyRequest(data)

	additionalFields := make(map[string]interface{})

	if err := json.Unmarshal(bs, &additionalFields); err != nil {
		return err
	}
	delete(additionalFields, "journeyId")

	c.AdditionalProperties = additionalFields

	return nil
}

func (c PatchUpdateJourneyRequest) MarshalJSON() ([]byte, error) {
	out := map[string]interface{}{}
	bs, err := json.Marshal(_PatchUpdateJourneyRequest(c))
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(bs), &out); err != nil {
		return nil, err
	}

	bs, err = json.Marshal(c.AdditionalProperties)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(bs), &out); err != nil {
		return nil, err
	}

	return json.Marshal(out)
}