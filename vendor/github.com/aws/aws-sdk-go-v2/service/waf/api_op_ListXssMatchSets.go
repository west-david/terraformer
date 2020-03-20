// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// A request to list the XssMatchSet objects created by the current AWS account.
type ListXssMatchSetsInput struct {
	_ struct{} `type:"structure"`

	// Specifies the number of XssMatchSet objects that you want AWS WAF to return
	// for this request. If you have more XssMatchSet objects than the number you
	// specify for Limit, the response includes a NextMarker value that you can
	// use to get another batch of Rules.
	Limit *int64 `type:"integer"`

	// If you specify a value for Limit and you have more XssMatchSet objects than
	// the value of Limit, AWS WAF returns a NextMarker value in the response that
	// allows you to list another group of XssMatchSets. For the second and subsequent
	// ListXssMatchSets requests, specify the value of NextMarker from the previous
	// response to get information about another batch of XssMatchSets.
	NextMarker *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListXssMatchSetsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListXssMatchSetsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListXssMatchSetsInput"}
	if s.NextMarker != nil && len(*s.NextMarker) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextMarker", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The response to a ListXssMatchSets request.
type ListXssMatchSetsOutput struct {
	_ struct{} `type:"structure"`

	// If you have more XssMatchSet objects than the number that you specified for
	// Limit in the request, the response includes a NextMarker value. To list more
	// XssMatchSet objects, submit another ListXssMatchSets request, and specify
	// the NextMarker value from the response in the NextMarker value in the next
	// request.
	NextMarker *string `min:"1" type:"string"`

	// An array of XssMatchSetSummary objects.
	XssMatchSets []XssMatchSetSummary `type:"list"`
}

// String returns the string representation
func (s ListXssMatchSetsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListXssMatchSets = "ListXssMatchSets"

// ListXssMatchSetsRequest returns a request value for making API operation for
// AWS WAF.
//
// Returns an array of XssMatchSet objects.
//
//    // Example sending a request using ListXssMatchSetsRequest.
//    req := client.ListXssMatchSetsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-2015-08-24/ListXssMatchSets
func (c *Client) ListXssMatchSetsRequest(input *ListXssMatchSetsInput) ListXssMatchSetsRequest {
	op := &aws.Operation{
		Name:       opListXssMatchSets,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListXssMatchSetsInput{}
	}

	req := c.newRequest(op, input, &ListXssMatchSetsOutput{})
	return ListXssMatchSetsRequest{Request: req, Input: input, Copy: c.ListXssMatchSetsRequest}
}

// ListXssMatchSetsRequest is the request type for the
// ListXssMatchSets API operation.
type ListXssMatchSetsRequest struct {
	*aws.Request
	Input *ListXssMatchSetsInput
	Copy  func(*ListXssMatchSetsInput) ListXssMatchSetsRequest
}

// Send marshals and sends the ListXssMatchSets API request.
func (r ListXssMatchSetsRequest) Send(ctx context.Context) (*ListXssMatchSetsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListXssMatchSetsResponse{
		ListXssMatchSetsOutput: r.Request.Data.(*ListXssMatchSetsOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListXssMatchSetsResponse is the response type for the
// ListXssMatchSets API operation.
type ListXssMatchSetsResponse struct {
	*ListXssMatchSetsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListXssMatchSets request.
func (r *ListXssMatchSetsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
