// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"epilot-journey/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateJourneyResponse struct {
	ContentType string
	// Success
	JourneyResponse *shared.JourneyResponse
	StatusCode      int
	RawResponse     *http.Response
}
