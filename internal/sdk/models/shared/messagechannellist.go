// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type MessageChannelList struct {
	Channels []MessageChannel `json:"channels"`
}

func (o *MessageChannelList) GetChannels() []MessageChannel {
	if o == nil {
		return []MessageChannel{}
	}
	return o.Channels
}
