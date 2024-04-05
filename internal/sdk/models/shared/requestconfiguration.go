// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// RequestConfiguration - # Request Configuration Object
// ### Description
// The `RequestConfiguration` object is used to represent a request configuration.
//
// ### Usage Example
// Returned from the `GET Request Configurations` endpoint.
type RequestConfiguration struct {
	// A bool representing whether or not to allow requests for this resource.
	AllowRequests bool `json:"allow_requests"`
	// A bool representing whether or not to automatically approve requests for this resource.
	AutoApproval bool       `json:"auto_approval"`
	Condition    *Condition `json:"condition,omitempty"`
	// The maximum duration for which the resource can be requested (in minutes).
	MaxDuration *int64 `json:"max_duration_minutes,omitempty"`
	// The priority of the request configuration.
	Priority int64 `json:"priority"`
	// The recommended duration for which the resource should be requested (in minutes). -1 represents an indefinite duration.
	RecommendedDuration *int64 `json:"recommended_duration_minutes,omitempty"`
	// The ID of the associated request template.
	RequestTemplateID *string `json:"request_template_id,omitempty"`
	// A bool representing whether or not to require MFA for requesting access to this resource.
	RequireMfaToRequest bool `json:"require_mfa_to_request"`
	// A bool representing whether or not access requests to the resource require an access ticket.
	RequireSupportTicket bool `json:"require_support_ticket"`
	// The list of reviewer stages for the request configuration.
	ReviewerStages []ReviewerStage `json:"reviewer_stages"`
}

func (o *RequestConfiguration) GetAllowRequests() bool {
	if o == nil {
		return false
	}
	return o.AllowRequests
}

func (o *RequestConfiguration) GetAutoApproval() bool {
	if o == nil {
		return false
	}
	return o.AutoApproval
}

func (o *RequestConfiguration) GetCondition() *Condition {
	if o == nil {
		return nil
	}
	return o.Condition
}

func (o *RequestConfiguration) GetMaxDuration() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxDuration
}

func (o *RequestConfiguration) GetPriority() int64 {
	if o == nil {
		return 0
	}
	return o.Priority
}

func (o *RequestConfiguration) GetRecommendedDuration() *int64 {
	if o == nil {
		return nil
	}
	return o.RecommendedDuration
}

func (o *RequestConfiguration) GetRequestTemplateID() *string {
	if o == nil {
		return nil
	}
	return o.RequestTemplateID
}

func (o *RequestConfiguration) GetRequireMfaToRequest() bool {
	if o == nil {
		return false
	}
	return o.RequireMfaToRequest
}

func (o *RequestConfiguration) GetRequireSupportTicket() bool {
	if o == nil {
		return false
	}
	return o.RequireSupportTicket
}

func (o *RequestConfiguration) GetReviewerStages() []ReviewerStage {
	if o == nil {
		return []ReviewerStage{}
	}
	return o.ReviewerStages
}
