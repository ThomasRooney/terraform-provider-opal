// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/models/shared"
	"net/http"
)

type UpdateOwnerUsersRequest struct {
	UserIDList shared.UserIDList `request:"mediaType=application/json"`
	// The ID of the owner.
	ID string `pathParam:"style=simple,explode=false,name=owner_id"`
}

func (o *UpdateOwnerUsersRequest) GetUserIDList() shared.UserIDList {
	if o == nil {
		return shared.UserIDList{}
	}
	return o.UserIDList
}

func (o *UpdateOwnerUsersRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type UpdateOwnerUsersResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The updated users for the owner.
	UserList *shared.UserList
}

func (o *UpdateOwnerUsersResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateOwnerUsersResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateOwnerUsersResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateOwnerUsersResponse) GetUserList() *shared.UserList {
	if o == nil {
		return nil
	}
	return o.UserList
}
