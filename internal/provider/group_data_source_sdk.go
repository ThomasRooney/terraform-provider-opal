// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/models/operations"
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/models/shared"
)

func (r *GroupDataSourceModel) RefreshFromSharedGroup(resp *shared.Group) {
	if resp != nil {
		r.ID = types.StringPointerValue(resp.ID)
	}
}

func (r *GroupDataSourceModel) RefreshFromOperationsGetGroupMessageChannelsResponseBody(resp *operations.GetGroupMessageChannelsResponseBody) {
	if resp != nil {
	}
}

func (r *GroupDataSourceModel) RefreshFromOperationsGetGroupVisibilityResponseBody(resp *operations.GetGroupVisibilityResponseBody) {
	if resp != nil {
		r.Visibility = types.StringValue(string(resp.Visibility))
		r.VisibilityGroupIds = nil
		for _, v := range resp.VisibilityGroupIds {
			r.VisibilityGroupIds = append(r.VisibilityGroupIds, types.StringValue(v))
		}
	}
}
