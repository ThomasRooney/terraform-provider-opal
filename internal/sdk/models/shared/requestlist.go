// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// RequestList - # Request List
// ### Description
// The `RequestList` object is used to represent a list of requests.
//
// ### Usage Example
// Returned from the `GET Requests` endpoint.
type RequestList struct {
	// The cursor to use in the next request to get the next page of results.
	Cursor *string `json:"cursor,omitempty"`
	// The list of requests.
	Requests []Request `json:"requests,omitempty"`
}

func (o *RequestList) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

func (o *RequestList) GetRequests() []Request {
	if o == nil {
		return nil
	}
	return o.Requests
}
