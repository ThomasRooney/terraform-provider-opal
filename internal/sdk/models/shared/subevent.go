// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/internal/utils"
)

// SubEvent - # Sub event Object
// ### Description
// The `SubEvent` object is used to represent a subevent.
//
// ### Usage Example
// Fetch from the `LIST Events` endpoint.
type SubEvent struct {
	AdditionalProperties interface{} `additionalProperties:"true" json:"-"`
	// The subevent type.
	SubEventType string `json:"sub_event_type"`
}

func (s SubEvent) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SubEvent) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SubEvent) GetAdditionalProperties() interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *SubEvent) GetSubEventType() string {
	if o == nil {
		return ""
	}
	return o.SubEventType
}
