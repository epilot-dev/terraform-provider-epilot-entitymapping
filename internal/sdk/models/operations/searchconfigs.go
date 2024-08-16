// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-entitymapping/internal/sdk/models/shared"
	"net/http"
)

type SearchConfigsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Entity Mapping Config
	MappingConfig *shared.MappingConfig
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *SearchConfigsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *SearchConfigsResponse) GetMappingConfig() *shared.MappingConfig {
	if o == nil {
		return nil
	}
	return o.MappingConfig
}

func (o *SearchConfigsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *SearchConfigsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
