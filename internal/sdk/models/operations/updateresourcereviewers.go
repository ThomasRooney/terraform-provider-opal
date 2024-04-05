// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/models/shared"
	"net/http"
)

type UpdateResourceReviewersRequest struct {
	ReviewerIDList shared.ReviewerIDList `request:"mediaType=application/json"`
	// The ID of the resource.
	ResourceID string `pathParam:"style=simple,explode=false,name=resource_id"`
}

func (o *UpdateResourceReviewersRequest) GetReviewerIDList() shared.ReviewerIDList {
	if o == nil {
		return shared.ReviewerIDList{}
	}
	return o.ReviewerIDList
}

func (o *UpdateResourceReviewersRequest) GetResourceID() string {
	if o == nil {
		return ""
	}
	return o.ResourceID
}

type UpdateResourceReviewersResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The updated IDs of owners that are reviewers for this resource
	Strings []string
}

func (o *UpdateResourceReviewersResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateResourceReviewersResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateResourceReviewersResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateResourceReviewersResponse) GetStrings() []string {
	if o == nil {
		return nil
	}
	return o.Strings
}
