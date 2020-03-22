// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// ManagementConditionRequestBuilder is request builder for ManagementCondition
type ManagementConditionRequestBuilder struct{ BaseRequestBuilder }

// Request returns ManagementConditionRequest
func (b *ManagementConditionRequestBuilder) Request() *ManagementConditionRequest {
	return &ManagementConditionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ManagementConditionRequest is request for ManagementCondition
type ManagementConditionRequest struct{ BaseRequest }

// Get performs GET request for ManagementCondition
func (r *ManagementConditionRequest) Get(ctx context.Context) (resObj *ManagementCondition, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ManagementCondition
func (r *ManagementConditionRequest) Update(ctx context.Context, reqObj *ManagementCondition) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ManagementCondition
func (r *ManagementConditionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ManagementConditionStatementRequestBuilder is request builder for ManagementConditionStatement
type ManagementConditionStatementRequestBuilder struct{ BaseRequestBuilder }

// Request returns ManagementConditionStatementRequest
func (b *ManagementConditionStatementRequestBuilder) Request() *ManagementConditionStatementRequest {
	return &ManagementConditionStatementRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ManagementConditionStatementRequest is request for ManagementConditionStatement
type ManagementConditionStatementRequest struct{ BaseRequest }

// Get performs GET request for ManagementConditionStatement
func (r *ManagementConditionStatementRequest) Get(ctx context.Context) (resObj *ManagementConditionStatement, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ManagementConditionStatement
func (r *ManagementConditionStatementRequest) Update(ctx context.Context, reqObj *ManagementConditionStatement) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ManagementConditionStatement
func (r *ManagementConditionStatementRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}