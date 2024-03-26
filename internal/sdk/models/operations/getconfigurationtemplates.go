// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/models/shared"
	"net/http"
)

type GetConfigurationTemplatesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// One page worth of configuration templates for your organization.
	PaginatedConfigurationTemplateList *shared.PaginatedConfigurationTemplateList
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetConfigurationTemplatesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetConfigurationTemplatesResponse) GetPaginatedConfigurationTemplateList() *shared.PaginatedConfigurationTemplateList {
	if o == nil {
		return nil
	}
	return o.PaginatedConfigurationTemplateList
}

func (o *GetConfigurationTemplatesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetConfigurationTemplatesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
