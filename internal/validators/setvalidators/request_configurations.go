// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package setvalidators

import (
	"context"
	"sort"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/opalsecurity/terraform-provider-opal/internal/provider/types"
	"github.com/hashicorp/terraform-plugin-framework-validators/helpers/validatordiag"
)

var _ validator.Set = SetRequestConfigurationsValidator{}

type SetRequestConfigurationsValidator struct{}

// Description describes the validation in plain text formatting.
func (v SetRequestConfigurationsValidator) Description(_ context.Context) string {
	return "validate request configurations to be well formed"
}

// MarkdownDescription describes the validation in Markdown formatting.
func (v SetRequestConfigurationsValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

// Validate performs the validation.
func (v SetRequestConfigurationsValidator) ValidateSet(ctx context.Context, req validator.SetRequest, resp *validator.SetResponse) {
	requestConfigurations := []types.RequestConfiguration{}

	req.ConfigValue.ElementsAs(ctx, &requestConfigurations, true)


	sort.Slice(requestConfigurations, func(i, j int) bool {
		return requestConfigurations[i].Priority.ValueInt64() < requestConfigurations[j].Priority.ValueInt64()
	})

	for idx, requestConfiguration := range requestConfigurations {
		if requestConfiguration.Priority.ValueInt64() != int64(idx) {
			resp.Diagnostics.Append(validatordiag.InvalidAttributeTypeDiagnostic(
				req.Path,
				"request configurations must have sequential priority values starting from 0",
				req.Path.String()+": "+v.Description(ctx),
			))
		}
		
		if requestConfiguration.Priority.ValueInt64() != 0 && requestConfiguration.Condition == nil {
			resp.Diagnostics.Append(validatordiag.InvalidAttributeTypeDiagnostic(
				req.Path,
				"request configurations must have a condition when priority is not 0",
				req.Path.String()+": "+v.Description(ctx),
			))
		} else if requestConfiguration.Priority.ValueInt64() == 0 && requestConfiguration.Condition != nil {
			resp.Diagnostics.Append(validatordiag.InvalidAttributeTypeDiagnostic(
				req.Path,
				"default request configurations must not have a condition",
				req.Path.String()+": "+v.Description(ctx),
			))
		}

		hasReviewerStages, isAutoApprove, isNotRequestable := len(requestConfiguration.ReviewerStages) > 0, requestConfiguration.AutoApproval.ValueBool(), !requestConfiguration.AllowRequests.ValueBool()
		if !(hasReviewerStages || isAutoApprove || isNotRequestable) {
			resp.Diagnostics.Append(validatordiag.InvalidAttributeTypeDiagnostic(
				req.Path,
				"invalid request configuration. Please specify a reviewer_stage, set auto_approval to true, or set allow_requests to false",
				req.Path.String()+": "+v.Description(ctx),
			))
		}

		if hasReviewerStages && isAutoApprove {
			resp.Diagnostics.Append(validatordiag.InvalidAttributeTypeDiagnostic(
				req.Path,
				"request configuration cannot have reviewer_stages when auto_approve is set to true",
				req.Path.String()+": "+v.Description(ctx),
			))
		}
	}

}

func RequestConfigurations() validator.Set {
	return SetRequestConfigurationsValidator{}
}
