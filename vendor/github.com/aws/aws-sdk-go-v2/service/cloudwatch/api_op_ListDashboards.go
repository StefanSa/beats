// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatch

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/monitoring-2010-08-01/ListDashboardsInput
type ListDashboardsInput struct {
	_ struct{} `type:"structure"`

	// If you specify this parameter, only the dashboards with names starting with
	// the specified string are listed. The maximum length is 255, and valid characters
	// are A-Z, a-z, 0-9, ".", "-", and "_".
	DashboardNamePrefix *string `type:"string"`

	// The token returned by a previous call to indicate that there is more data
	// available.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListDashboardsInput) String() string {
	return awsutil.Prettify(s)
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/monitoring-2010-08-01/ListDashboardsOutput
type ListDashboardsOutput struct {
	_ struct{} `type:"structure"`

	// The list of matching dashboards.
	DashboardEntries []DashboardEntry `type:"list"`

	// The token that marks the start of the next batch of returned results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListDashboardsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListDashboards = "ListDashboards"

// ListDashboardsRequest returns a request value for making API operation for
// Amazon CloudWatch.
//
// Returns a list of the dashboards for your account. If you include DashboardNamePrefix,
// only those dashboards with names starting with the prefix are listed. Otherwise,
// all dashboards in your account are listed.
//
// ListDashboards returns up to 1000 results on one page. If there are more
// than 1000 dashboards, you can call ListDashboards again and include the value
// you received for NextToken in the first call, to receive the next 1000 results.
//
//    // Example sending a request using ListDashboardsRequest.
//    req := client.ListDashboardsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/monitoring-2010-08-01/ListDashboards
func (c *Client) ListDashboardsRequest(input *ListDashboardsInput) ListDashboardsRequest {
	op := &aws.Operation{
		Name:       opListDashboards,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListDashboardsInput{}
	}

	req := c.newRequest(op, input, &ListDashboardsOutput{})
	return ListDashboardsRequest{Request: req, Input: input, Copy: c.ListDashboardsRequest}
}

// ListDashboardsRequest is the request type for the
// ListDashboards API operation.
type ListDashboardsRequest struct {
	*aws.Request
	Input *ListDashboardsInput
	Copy  func(*ListDashboardsInput) ListDashboardsRequest
}

// Send marshals and sends the ListDashboards API request.
func (r ListDashboardsRequest) Send(ctx context.Context) (*ListDashboardsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDashboardsResponse{
		ListDashboardsOutput: r.Request.Data.(*ListDashboardsOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListDashboardsRequestPaginator returns a paginator for ListDashboards.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListDashboardsRequest(input)
//   p := cloudwatch.NewListDashboardsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListDashboardsPaginator(req ListDashboardsRequest) ListDashboardsPaginator {
	return ListDashboardsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListDashboardsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListDashboardsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListDashboardsPaginator struct {
	aws.Pager
}

func (p *ListDashboardsPaginator) CurrentPage() *ListDashboardsOutput {
	return p.Pager.CurrentPage().(*ListDashboardsOutput)
}

// ListDashboardsResponse is the response type for the
// ListDashboards API operation.
type ListDashboardsResponse struct {
	*ListDashboardsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDashboards request.
func (r *ListDashboardsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}