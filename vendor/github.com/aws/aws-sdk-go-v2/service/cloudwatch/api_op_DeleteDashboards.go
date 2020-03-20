// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatch

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteDashboardsInput struct {
	_ struct{} `type:"structure"`

	// The dashboards to be deleted. This parameter is required.
	//
	// DashboardNames is a required field
	DashboardNames []string `type:"list" required:"true"`
}

// String returns the string representation
func (s DeleteDashboardsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDashboardsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteDashboardsInput"}

	if s.DashboardNames == nil {
		invalidParams.Add(aws.NewErrParamRequired("DashboardNames"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteDashboardsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteDashboardsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteDashboards = "DeleteDashboards"

// DeleteDashboardsRequest returns a request value for making API operation for
// Amazon CloudWatch.
//
// Deletes all dashboards that you specify. You may specify up to 100 dashboards
// to delete. If there is an error during this call, no dashboards are deleted.
//
//    // Example sending a request using DeleteDashboardsRequest.
//    req := client.DeleteDashboardsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/monitoring-2010-08-01/DeleteDashboards
func (c *Client) DeleteDashboardsRequest(input *DeleteDashboardsInput) DeleteDashboardsRequest {
	op := &aws.Operation{
		Name:       opDeleteDashboards,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteDashboardsInput{}
	}

	req := c.newRequest(op, input, &DeleteDashboardsOutput{})
	return DeleteDashboardsRequest{Request: req, Input: input, Copy: c.DeleteDashboardsRequest}
}

// DeleteDashboardsRequest is the request type for the
// DeleteDashboards API operation.
type DeleteDashboardsRequest struct {
	*aws.Request
	Input *DeleteDashboardsInput
	Copy  func(*DeleteDashboardsInput) DeleteDashboardsRequest
}

// Send marshals and sends the DeleteDashboards API request.
func (r DeleteDashboardsRequest) Send(ctx context.Context) (*DeleteDashboardsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteDashboardsResponse{
		DeleteDashboardsOutput: r.Request.Data.(*DeleteDashboardsOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteDashboardsResponse is the response type for the
// DeleteDashboards API operation.
type DeleteDashboardsResponse struct {
	*DeleteDashboardsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteDashboards request.
func (r *DeleteDashboardsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
