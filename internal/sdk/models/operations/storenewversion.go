// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-entitymapping/internal/sdk/models/shared"
	"net/http"
)

type StoreNewVersionRequest struct {
	// Mapping Config to store
	MappingConfig *shared.MappingConfig `request:"mediaType=application/json"`
	// Mapping Config Id
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *StoreNewVersionRequest) GetMappingConfig() *shared.MappingConfig {
	if o == nil {
		return nil
	}
	return o.MappingConfig
}

func (o *StoreNewVersionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type StoreNewVersionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// The updated entity mapping config
	MappingConfig *shared.MappingConfig
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *StoreNewVersionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *StoreNewVersionResponse) GetMappingConfig() *shared.MappingConfig {
	if o == nil {
		return nil
	}
	return o.MappingConfig
}

func (o *StoreNewVersionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *StoreNewVersionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
