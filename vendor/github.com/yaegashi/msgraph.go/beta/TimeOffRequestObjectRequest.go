// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// TimeOffRequestObjectRequestBuilder is request builder for TimeOffRequestObject
type TimeOffRequestObjectRequestBuilder struct{ BaseRequestBuilder }

// Request returns TimeOffRequestObjectRequest
func (b *TimeOffRequestObjectRequestBuilder) Request() *TimeOffRequestObjectRequest {
	return &TimeOffRequestObjectRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TimeOffRequestObjectRequest is request for TimeOffRequestObject
type TimeOffRequestObjectRequest struct{ BaseRequest }

// Get performs GET request for TimeOffRequestObject
func (r *TimeOffRequestObjectRequest) Get(ctx context.Context) (resObj *TimeOffRequestObject, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TimeOffRequestObject
func (r *TimeOffRequestObjectRequest) Update(ctx context.Context, reqObj *TimeOffRequestObject) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TimeOffRequestObject
func (r *TimeOffRequestObjectRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}