// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateBillingGroupInput struct {
	_ struct{} `type:"structure"`

	// The name you wish to give to the billing group.
	//
	// BillingGroupName is a required field
	BillingGroupName *string `location:"uri" locationName:"billingGroupName" min:"1" type:"string" required:"true"`

	// The properties of the billing group.
	BillingGroupProperties *BillingGroupProperties `locationName:"billingGroupProperties" type:"structure"`

	// Metadata which can be used to manage the billing group.
	Tags []Tag `locationName:"tags" type:"list"`
}

// String returns the string representation
func (s CreateBillingGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateBillingGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateBillingGroupInput"}

	if s.BillingGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("BillingGroupName"))
	}
	if s.BillingGroupName != nil && len(*s.BillingGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("BillingGroupName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateBillingGroupInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.BillingGroupProperties != nil {
		v := s.BillingGroupProperties

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "billingGroupProperties", v, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "tags", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.BillingGroupName != nil {
		v := *s.BillingGroupName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "billingGroupName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type CreateBillingGroupOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the billing group.
	BillingGroupArn *string `locationName:"billingGroupArn" type:"string"`

	// The ID of the billing group.
	BillingGroupId *string `locationName:"billingGroupId" min:"1" type:"string"`

	// The name you gave to the billing group.
	BillingGroupName *string `locationName:"billingGroupName" min:"1" type:"string"`
}

// String returns the string representation
func (s CreateBillingGroupOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateBillingGroupOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.BillingGroupArn != nil {
		v := *s.BillingGroupArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "billingGroupArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.BillingGroupId != nil {
		v := *s.BillingGroupId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "billingGroupId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.BillingGroupName != nil {
		v := *s.BillingGroupName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "billingGroupName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateBillingGroup = "CreateBillingGroup"

// CreateBillingGroupRequest returns a request value for making API operation for
// AWS IoT.
//
// Creates a billing group.
//
//    // Example sending a request using CreateBillingGroupRequest.
//    req := client.CreateBillingGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CreateBillingGroupRequest(input *CreateBillingGroupInput) CreateBillingGroupRequest {
	op := &aws.Operation{
		Name:       opCreateBillingGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/billing-groups/{billingGroupName}",
	}

	if input == nil {
		input = &CreateBillingGroupInput{}
	}

	req := c.newRequest(op, input, &CreateBillingGroupOutput{})
	return CreateBillingGroupRequest{Request: req, Input: input, Copy: c.CreateBillingGroupRequest}
}

// CreateBillingGroupRequest is the request type for the
// CreateBillingGroup API operation.
type CreateBillingGroupRequest struct {
	*aws.Request
	Input *CreateBillingGroupInput
	Copy  func(*CreateBillingGroupInput) CreateBillingGroupRequest
}

// Send marshals and sends the CreateBillingGroup API request.
func (r CreateBillingGroupRequest) Send(ctx context.Context) (*CreateBillingGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateBillingGroupResponse{
		CreateBillingGroupOutput: r.Request.Data.(*CreateBillingGroupOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateBillingGroupResponse is the response type for the
// CreateBillingGroup API operation.
type CreateBillingGroupResponse struct {
	*CreateBillingGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateBillingGroup request.
func (r *CreateBillingGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
