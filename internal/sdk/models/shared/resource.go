// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// # Resource Object
// ### Description
// The `Resource` object is used to represent a resource.
//
// ### Usage Example
// Update from the `UPDATE Resources` endpoint.
type Resource struct {
	// The ID of the owner of the resource.
	AdminOwnerID *string `json:"admin_owner_id,omitempty"`
	// The ID of the app.
	AppID *string `json:"app_id,omitempty"`
	// A description of the resource.
	Description *string `json:"description,omitempty"`
	// The ID of the resource.
	ID *string `json:"resource_id,omitempty"`
	// The name of the resource.
	Name *string `json:"name,omitempty"`
	// The ID of the parent resource.
	ParentResourceID *string `json:"parent_resource_id,omitempty"`
	// Information that defines the remote resource. This replaces the deprecated remote_id and metadata fields.
	RemoteInfo *ResourceRemoteInfo `json:"remote_info,omitempty"`
	// A list of configurations for requests to this resource.
	RequestConfigurations []RequestConfiguration `json:"request_configurations,omitempty"`
	// A bool representing whether or not to require MFA for reviewers to approve requests for this resource.
	RequireMfaToApprove *bool `json:"require_mfa_to_approve,omitempty"`
	// A bool representing whether or not to require MFA to connect to this resource.
	RequireMfaToConnect *bool `json:"require_mfa_to_connect,omitempty"`
	// The type of the resource.
	ResourceType *ResourceTypeEnum `json:"resource_type,omitempty"`
}

func (o *Resource) GetAdminOwnerID() *string {
	if o == nil {
		return nil
	}
	return o.AdminOwnerID
}

func (o *Resource) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *Resource) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *Resource) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Resource) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Resource) GetParentResourceID() *string {
	if o == nil {
		return nil
	}
	return o.ParentResourceID
}

func (o *Resource) GetRemoteInfo() *ResourceRemoteInfo {
	if o == nil {
		return nil
	}
	return o.RemoteInfo
}

func (o *Resource) GetRequestConfigurations() []RequestConfiguration {
	if o == nil {
		return nil
	}
	return o.RequestConfigurations
}

func (o *Resource) GetRequireMfaToApprove() *bool {
	if o == nil {
		return nil
	}
	return o.RequireMfaToApprove
}

func (o *Resource) GetRequireMfaToConnect() *bool {
	if o == nil {
		return nil
	}
	return o.RequireMfaToConnect
}

func (o *Resource) GetResourceType() *ResourceTypeEnum {
	if o == nil {
		return nil
	}
	return o.ResourceType
}
