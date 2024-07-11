// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// # UpdateResourceInfo Object
// ### Description
// The `UpdateResourceInfo` object is used as an input to the UpdateResource API.
type UpdateResourceInfo struct {
	// The ID of the owner of the resource.
	AdminOwnerID *string `json:"admin_owner_id,omitempty"`
	// A description of the resource.
	Description *string `json:"description,omitempty"`
	// The ID of the resource.
	ID *string `json:"resource_id,omitempty"`
	// The name of the resource.
	Name *string `json:"name,omitempty"`
	// A list of configurations for requests to this resource. If not provided, the default request configuration will be used.
	RequestConfigurations []RequestConfiguration `json:"request_configurations"`
	// A bool representing whether or not to require MFA for reviewers to approve requests for this resource.
	RequireMfaToApprove *bool `json:"require_mfa_to_approve,omitempty"`
	// A bool representing whether or not to require MFA to connect to this resource.
	RequireMfaToConnect *bool `json:"require_mfa_to_connect,omitempty"`
	// Configuration for ticket propagation, when enabled, a ticket will be created for access changes related to the users in this resource.
	TicketPropagation *TicketPropagationConfiguration `json:"ticket_propagation,omitempty"`
}

func (o *UpdateResourceInfo) GetAdminOwnerID() *string {
	if o == nil {
		return nil
	}
	return o.AdminOwnerID
}

func (o *UpdateResourceInfo) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *UpdateResourceInfo) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UpdateResourceInfo) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *UpdateResourceInfo) GetRequestConfigurations() []RequestConfiguration {
	if o == nil {
		return []RequestConfiguration{}
	}
	return o.RequestConfigurations
}

func (o *UpdateResourceInfo) GetRequireMfaToApprove() *bool {
	if o == nil {
		return nil
	}
	return o.RequireMfaToApprove
}

func (o *UpdateResourceInfo) GetRequireMfaToConnect() *bool {
	if o == nil {
		return nil
	}
	return o.RequireMfaToConnect
}

func (o *UpdateResourceInfo) GetTicketPropagation() *TicketPropagationConfiguration {
	if o == nil {
		return nil
	}
	return o.TicketPropagation
}
