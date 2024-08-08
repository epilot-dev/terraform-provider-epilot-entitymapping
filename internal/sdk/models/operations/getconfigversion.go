// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-entitymapping/internal/sdk/models/shared"
	"net/http"
)

type GetConfigVersionRequest struct {
	// Mapping Config Id
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Version to be loaded
	Version float64 `pathParam:"style=simple,explode=false,name=version"`
}

func (o *GetConfigVersionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetConfigVersionRequest) GetVersion() float64 {
	if o == nil {
		return 0.0
	}
	return o.Version
}

type GetConfigVersionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Entity Mapping Config
	MappingConfig *shared.MappingConfig
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetConfigVersionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetConfigVersionResponse) GetMappingConfig() *shared.MappingConfig {
	if o == nil {
		return nil
	}
	return o.MappingConfig
}

func (o *GetConfigVersionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetConfigVersionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
