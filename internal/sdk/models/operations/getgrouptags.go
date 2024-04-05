// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/models/shared"
	"net/http"
)

type GetGroupTagsRequest struct {
	// The ID of the group whose tags to return.
	GroupID string `pathParam:"style=simple,explode=false,name=group_id"`
}

func (o *GetGroupTagsRequest) GetGroupID() string {
	if o == nil {
		return ""
	}
	return o.GroupID
}

type GetGroupTagsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The tags applied to the group.
	TagsList *shared.TagsList
}

func (o *GetGroupTagsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetGroupTagsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetGroupTagsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetGroupTagsResponse) GetTagsList() *shared.TagsList {
	if o == nil {
		return nil
	}
	return o.TagsList
}
