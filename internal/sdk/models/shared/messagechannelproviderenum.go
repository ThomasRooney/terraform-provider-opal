// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// MessageChannelProviderEnum - The third party provider of the message channel.
type MessageChannelProviderEnum string

const (
	MessageChannelProviderEnumSlack MessageChannelProviderEnum = "SLACK"
)

func (e MessageChannelProviderEnum) ToPointer() *MessageChannelProviderEnum {
	return &e
}
func (e *MessageChannelProviderEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SLACK":
		*e = MessageChannelProviderEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MessageChannelProviderEnum: %v", v)
	}
}
