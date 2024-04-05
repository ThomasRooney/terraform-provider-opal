// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// # CreateOwnerInfo Object
// ### Description
// The `CreateOwnerInfo` object is used to store creation info for an owner.
//
// ### Usage Example
// Use in the `POST Owners` endpoint.
type CreateOwnerInfo struct {
	// The amount of time (in minutes) before the next reviewer is notified. Use 0 to remove escalation policy.
	AccessRequestEscalationPeriod *int64 `json:"access_request_escalation_period,omitempty"`
	// A description of the owner.
	Description *string `json:"description,omitempty"`
	// The name of the owner.
	Name string `json:"name"`
	// The message channel id for the reviewer channel.
	ReviewerMessageChannelID *string `json:"reviewer_message_channel_id,omitempty"`
	// Sync this owner's user list with a source group.
	SourceGroupID *string `json:"source_group_id,omitempty"`
	// Users to add to the created owner. If setting a source_group_id this list must be empty.
	UserIds []string `json:"user_ids"`
}

func (o *CreateOwnerInfo) GetAccessRequestEscalationPeriod() *int64 {
	if o == nil {
		return nil
	}
	return o.AccessRequestEscalationPeriod
}

func (o *CreateOwnerInfo) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CreateOwnerInfo) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateOwnerInfo) GetReviewerMessageChannelID() *string {
	if o == nil {
		return nil
	}
	return o.ReviewerMessageChannelID
}

func (o *CreateOwnerInfo) GetSourceGroupID() *string {
	if o == nil {
		return nil
	}
	return o.SourceGroupID
}

func (o *CreateOwnerInfo) GetUserIds() []string {
	if o == nil {
		return []string{}
	}
	return o.UserIds
}
