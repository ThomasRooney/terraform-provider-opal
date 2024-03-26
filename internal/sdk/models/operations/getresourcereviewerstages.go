// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/models/shared"
	"net/http"
)

type GetResourceReviewerStagesRequest struct {
	// The ID of the resource.
	ResourceID string `pathParam:"style=simple,explode=false,name=resource_id"`
}

func (o *GetResourceReviewerStagesRequest) GetResourceID() string {
	if o == nil {
		return ""
	}
	return o.ResourceID
}

type GetResourceReviewerStagesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The reviewer stages for this resource.
	Classes []shared.ReviewerStage
}

func (o *GetResourceReviewerStagesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetResourceReviewerStagesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetResourceReviewerStagesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetResourceReviewerStagesResponse) GetClasses() []shared.ReviewerStage {
	if o == nil {
		return nil
	}
	return o.Classes
}
