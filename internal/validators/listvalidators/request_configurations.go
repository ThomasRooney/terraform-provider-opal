// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package listvalidators

import (
	"context"
	"slices"
	"sort"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/opalsecurity/terraform-provider-opal/internal/provider/types"
)

var _ validator.List = ListRequestConfigurationsValidator{}

type ListRequestConfigurationsValidator struct{}

// Description describes the validation in plain text formatting.
func (v ListRequestConfigurationsValidator) Description(_ context.Context) string {
	return "TODO: add validator description"
}

// MarkdownDescription describes the validation in Markdown formatting.
func (v ListRequestConfigurationsValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

// Validate performs the validation.
func (v ListRequestConfigurationsValidator) ValidateList(ctx context.Context, req validator.ListRequest, resp *validator.ListResponse) {
	requestConfigurations := []types.RequestConfiguration{}

	req.ConfigValue.ElementsAs(ctx, &requestConfigurations, true)

	sort.Slice(requestConfigurations, func(i, j int) bool {
		return requestConfigurations[i].Priority.ValueInt64() < requestConfigurations[j].Priority.ValueInt64()
	})

	for idx, requestConfiguration := range requestConfigurations {
		if requestConfiguration.Priority.ValueInt64() != int64(idx) {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				"request configurations must have sequential priority values starting from 0",
				req.Path.String()+": "+v.Description(ctx),
			)
		}
		
		if requestConfiguration.Priority.ValueInt64() != 0 && requestConfiguration.Condition == nil {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				"non-default request configurations must have a condition",
				req.Path.String()+": "+v.Description(ctx),
			)
		} else if requestConfiguration.Priority.ValueInt64() == 0 && requestConfiguration.Condition != nil {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				"default request configurations must not have a condition",
				req.Path.String()+": "+v.Description(ctx),
			)
		}

		hasReviewerStages, isAutoApprove, isNotRequestable := len(requestConfiguration.ReviewerStages) > 0, requestConfiguration.AutoApproval.ValueBool(), !requestConfiguration.AllowRequests.ValueBool()
		if !(hasReviewerStages || isAutoApprove || isNotRequestable) {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				"invalid request configuration. Please specify a reviewer_stage, set auto_approve to true, or set is_requestable to false",
				req.Path.String()+": "+v.Description(ctx),
			)
		}

		if hasReviewerStages && isAutoApprove {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				"request configuration cannot have reviewer_stages when auto_approve is set to true",
				req.Path.String()+": "+v.Description(ctx),
			)
		}
	}
}

func RequestConfigurations() validator.List {
	return ListRequestConfigurationsValidator{}
}
