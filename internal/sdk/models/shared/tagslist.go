// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type TagsList struct {
	Tags []Tag `json:"tags"`
}

func (o *TagsList) GetTags() []Tag {
	if o == nil {
		return []Tag{}
	}
	return o.Tags
}
