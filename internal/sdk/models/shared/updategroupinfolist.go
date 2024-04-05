// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type UpdateGroupInfoList struct {
	// A list of groups with information to update.
	Groups []UpdateGroupInfo `json:"groups"`
}

func (o *UpdateGroupInfoList) GetGroups() []UpdateGroupInfo {
	if o == nil {
		return []UpdateGroupInfo{}
	}
	return o.Groups
}
