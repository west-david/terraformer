// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatchevents

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListRulesInput struct {
	_ struct{} `type:"structure"`

	// Limits the results to show only the rules associated with the specified event
	// bus.
	EventBusName *string `min:"1" type:"string"`

	// The maximum number of results to return.
	Limit *int64 `min:"1" type:"integer"`

	// The prefix matching the rule name.
	NamePrefix *string `min:"1" type:"string"`

	// The token returned by a previous call to retrieve the next set of results.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListRulesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListRulesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListRulesInput"}
	if s.EventBusName != nil && len(*s.EventBusName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EventBusName", 1))
	}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}
	if s.NamePrefix != nil && len(*s.NamePrefix) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NamePrefix", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListRulesOutput struct {
	_ struct{} `type:"structure"`

	// Indicates whether there are additional results to retrieve. If there are
	// no more results, the value is null.
	NextToken *string `min:"1" type:"string"`

	// The rules that match the specified criteria.
	Rules []Rule `type:"list"`
}

// String returns the string representation
func (s ListRulesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListRules = "ListRules"

// ListRulesRequest returns a request value for making API operation for
// Amazon CloudWatch Events.
//
// Lists your EventBridge rules. You can either list all the rules or provide
// a prefix to match to the rule names.
//
// ListRules doesn't list the targets of a rule. To see the targets associated
// with a rule, use ListTargetsByRule.
//
//    // Example sending a request using ListRulesRequest.
//    req := client.ListRulesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/events-2015-10-07/ListRules
func (c *Client) ListRulesRequest(input *ListRulesInput) ListRulesRequest {
	op := &aws.Operation{
		Name:       opListRules,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListRulesInput{}
	}

	req := c.newRequest(op, input, &ListRulesOutput{})
	return ListRulesRequest{Request: req, Input: input, Copy: c.ListRulesRequest}
}

// ListRulesRequest is the request type for the
// ListRules API operation.
type ListRulesRequest struct {
	*aws.Request
	Input *ListRulesInput
	Copy  func(*ListRulesInput) ListRulesRequest
}

// Send marshals and sends the ListRules API request.
func (r ListRulesRequest) Send(ctx context.Context) (*ListRulesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListRulesResponse{
		ListRulesOutput: r.Request.Data.(*ListRulesOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListRulesResponse is the response type for the
// ListRules API operation.
type ListRulesResponse struct {
	*ListRulesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListRules request.
func (r *ListRulesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
