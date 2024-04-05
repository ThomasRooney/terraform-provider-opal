// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// # Owner Object
// ### Description
// The `Owner` object is used to represent an owner.
type Owner struct {
	// The amount of time (in minutes) before the next reviewer is notified. Use 0 to remove escalation policy.
	AccessRequestEscalationPeriod *int64 `json:"access_request_escalation_period,omitempty"`
	// A description of the owner.
	Description *string `json:"description,omitempty"`
	// The ID of the owner.
	ID *string `json:"owner_id,omitempty"`
	// The name of the owner.
	Name                     *string `json:"name,omitempty"`
	ReviewerMessageChannelID *string `json:"reviewer_message_channel_id,omitempty"`
	SourceGroupID            *string `json:"source_group_id,omitempty"`
}

func (o *Owner) GetAccessRequestEscalationPeriod() *int64 {
	if o == nil {
		return nil
	}
	return o.AccessRequestEscalationPeriod
}

func (o *Owner) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *Owner) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Owner) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Owner) GetReviewerMessageChannelID() *string {
	if o == nil {
		return nil
	}
	return o.ReviewerMessageChannelID
}

func (o *Owner) GetSourceGroupID() *string {
	if o == nil {
		return nil
	}
	return o.SourceGroupID
}
