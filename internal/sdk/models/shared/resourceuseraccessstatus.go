// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/internal/utils"
	"time"
)

// ResourceUserAccessStatus - # AccessStatus Object
// ### Description
// The `AccessStatus` object is used to represent the user's access to the resource.
//
// ### Usage Example
// View the `AccessStatus` for a resource/user pair to determine if the user has access to the resource.
type ResourceUserAccessStatus struct {
	// # Access Level Object
	// ### Description
	// The `ResourceAccessLevel` object is used to represent the level of access that a user has to a resource or a resource has to a group. The "default" access
	// level is a `ResourceAccessLevel` object whose fields are all empty strings.
	//
	// ### Usage Example
	// View the `ResourceAccessLevel` of a resource/user or resource/group pair to see the level of access granted to the resource.
	AccessLevel *ResourceAccessLevel `json:"access_level,omitempty"`
	// The day and time the user's access will expire.
	ExpirationDate *time.Time `json:"expiration_date"`
	// The ID of the resource.
	ResourceID string `json:"resource_id"`
	// The status of the user's access to the resource.
	Status ResourceUserAccessStatusEnum `json:"status"`
	// The ID of the user.
	UserID string `json:"user_id"`
}

func (r ResourceUserAccessStatus) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *ResourceUserAccessStatus) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ResourceUserAccessStatus) GetAccessLevel() *ResourceAccessLevel {
	if o == nil {
		return nil
	}
	return o.AccessLevel
}

func (o *ResourceUserAccessStatus) GetExpirationDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.ExpirationDate
}

func (o *ResourceUserAccessStatus) GetResourceID() string {
	if o == nil {
		return ""
	}
	return o.ResourceID
}

func (o *ResourceUserAccessStatus) GetStatus() ResourceUserAccessStatusEnum {
	if o == nil {
		return ResourceUserAccessStatusEnum("")
	}
	return o.Status
}

func (o *ResourceUserAccessStatus) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}
