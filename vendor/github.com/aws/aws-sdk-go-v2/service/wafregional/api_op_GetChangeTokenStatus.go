// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafregional

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetChangeTokenStatusInput struct {
	_ struct{} `type:"structure"`

	// The change token for which you want to get the status. This change token
	// was previously returned in the GetChangeToken response.
	//
	// ChangeToken is a required field
	ChangeToken *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetChangeTokenStatusInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetChangeTokenStatusInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetChangeTokenStatusInput"}

	if s.ChangeToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("ChangeToken"))
	}
	if s.ChangeToken != nil && len(*s.ChangeToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ChangeToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetChangeTokenStatusOutput struct {
	_ struct{} `type:"structure"`

	// The status of the change token.
	ChangeTokenStatus ChangeTokenStatus `type:"string" enum:"true"`
}

// String returns the string representation
func (s GetChangeTokenStatusOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetChangeTokenStatus = "GetChangeTokenStatus"

// GetChangeTokenStatusRequest returns a request value for making API operation for
// AWS WAF Regional.
//
// Returns the status of a ChangeToken that you got by calling GetChangeToken.
// ChangeTokenStatus is one of the following values:
//
//    * PROVISIONED: You requested the change token by calling GetChangeToken,
//    but you haven't used it yet in a call to create, update, or delete an
//    AWS WAF object.
//
//    * PENDING: AWS WAF is propagating the create, update, or delete request
//    to all AWS WAF servers.
//
//    * INSYNC: Propagation is complete.
//
//    // Example sending a request using GetChangeTokenStatusRequest.
//    req := client.GetChangeTokenStatusRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28/GetChangeTokenStatus
func (c *Client) GetChangeTokenStatusRequest(input *GetChangeTokenStatusInput) GetChangeTokenStatusRequest {
	op := &aws.Operation{
		Name:       opGetChangeTokenStatus,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetChangeTokenStatusInput{}
	}

	req := c.newRequest(op, input, &GetChangeTokenStatusOutput{})
	return GetChangeTokenStatusRequest{Request: req, Input: input, Copy: c.GetChangeTokenStatusRequest}
}

// GetChangeTokenStatusRequest is the request type for the
// GetChangeTokenStatus API operation.
type GetChangeTokenStatusRequest struct {
	*aws.Request
	Input *GetChangeTokenStatusInput
	Copy  func(*GetChangeTokenStatusInput) GetChangeTokenStatusRequest
}

// Send marshals and sends the GetChangeTokenStatus API request.
func (r GetChangeTokenStatusRequest) Send(ctx context.Context) (*GetChangeTokenStatusResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetChangeTokenStatusResponse{
		GetChangeTokenStatusOutput: r.Request.Data.(*GetChangeTokenStatusOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetChangeTokenStatusResponse is the response type for the
// GetChangeTokenStatus API operation.
type GetChangeTokenStatusResponse struct {
	*GetChangeTokenStatusOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetChangeTokenStatus request.
func (r *GetChangeTokenStatusResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
