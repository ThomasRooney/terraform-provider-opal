// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/internal/utils"
	"time"
)

// GroupUser - # Group Access User Object
// ### Description
// The `GroupAccessUser` object is used to represent a user with access to a group.
//
// ### Usage Example
// Fetch from the `LIST GroupUsers` endpoint.
type GroupUser struct {
	// # Access Level Object
	// ### Description
	// The `GroupAccessLevel` object is used to represent the level of access that a user has to a group or a group has to a group. The "default" access
	// level is a `GroupAccessLevel` object whose fields are all empty strings.
	//
	// ### Usage Example
	// View the `GroupAccessLevel` of a group/user or group/group pair to see the level of access granted to the group.
	AccessLevel *GroupAccessLevel `json:"access_level,omitempty"`
	// The user's email.
	Email string `json:"email"`
	// The day and time the user's access will expire.
	ExpirationDate *time.Time `json:"expiration_date"`
	// The user's full name.
	FullName string `json:"full_name"`
	// The ID of the group.
	GroupID string `json:"group_id"`
	// The ID of the user.
	UserID string `json:"user_id"`
}

func (g GroupUser) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GroupUser) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GroupUser) GetAccessLevel() *GroupAccessLevel {
	if o == nil {
		return nil
	}
	return o.AccessLevel
}

func (o *GroupUser) GetEmail() string {
	if o == nil {
		return ""
	}
	return o.Email
}

func (o *GroupUser) GetExpirationDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.ExpirationDate
}

func (o *GroupUser) GetFullName() string {
	if o == nil {
		return ""
	}
	return o.FullName
}

func (o *GroupUser) GetGroupID() string {
	if o == nil {
		return ""
	}
	return o.GroupID
}

func (o *GroupUser) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}
