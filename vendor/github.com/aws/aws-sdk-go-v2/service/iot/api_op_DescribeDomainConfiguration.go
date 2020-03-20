// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DescribeDomainConfigurationInput struct {
	_ struct{} `type:"structure"`

	// The name of the domain configuration.
	//
	// DomainConfigurationName is a required field
	DomainConfigurationName *string `location:"uri" locationName:"domainConfigurationName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeDomainConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDomainConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeDomainConfigurationInput"}

	if s.DomainConfigurationName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainConfigurationName"))
	}
	if s.DomainConfigurationName != nil && len(*s.DomainConfigurationName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DomainConfigurationName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeDomainConfigurationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.DomainConfigurationName != nil {
		v := *s.DomainConfigurationName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "domainConfigurationName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DescribeDomainConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// An object that specifies the authorization service for a domain.
	AuthorizerConfig *AuthorizerConfig `locationName:"authorizerConfig" type:"structure"`

	// The ARN of the domain configuration.
	DomainConfigurationArn *string `locationName:"domainConfigurationArn" type:"string"`

	// The name of the domain configuration.
	DomainConfigurationName *string `locationName:"domainConfigurationName" min:"1" type:"string"`

	// A Boolean value that specifies the current state of the domain configuration.
	DomainConfigurationStatus DomainConfigurationStatus `locationName:"domainConfigurationStatus" type:"string" enum:"true"`

	// The name of the domain.
	DomainName *string `locationName:"domainName" min:"1" type:"string"`

	// The type of the domain.
	DomainType DomainType `locationName:"domainType" type:"string" enum:"true"`

	// A list containing summary information about the server certificate included
	// in the domain configuration.
	ServerCertificates []ServerCertificateSummary `locationName:"serverCertificates" type:"list"`

	// The type of service delivered by the endpoint.
	ServiceType ServiceType `locationName:"serviceType" type:"string" enum:"true"`
}

// String returns the string representation
func (s DescribeDomainConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeDomainConfigurationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.AuthorizerConfig != nil {
		v := s.AuthorizerConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "authorizerConfig", v, metadata)
	}
	if s.DomainConfigurationArn != nil {
		v := *s.DomainConfigurationArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "domainConfigurationArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DomainConfigurationName != nil {
		v := *s.DomainConfigurationName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "domainConfigurationName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.DomainConfigurationStatus) > 0 {
		v := s.DomainConfigurationStatus

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "domainConfigurationStatus", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.DomainName != nil {
		v := *s.DomainName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "domainName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.DomainType) > 0 {
		v := s.DomainType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "domainType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.ServerCertificates != nil {
		v := s.ServerCertificates

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "serverCertificates", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if len(s.ServiceType) > 0 {
		v := s.ServiceType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "serviceType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

const opDescribeDomainConfiguration = "DescribeDomainConfiguration"

// DescribeDomainConfigurationRequest returns a request value for making API operation for
// AWS IoT.
//
// Gets summary information about a domain configuration.
//
// The domain configuration feature is in public preview and is subject to change.
//
//    // Example sending a request using DescribeDomainConfigurationRequest.
//    req := client.DescribeDomainConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DescribeDomainConfigurationRequest(input *DescribeDomainConfigurationInput) DescribeDomainConfigurationRequest {
	op := &aws.Operation{
		Name:       opDescribeDomainConfiguration,
		HTTPMethod: "GET",
		HTTPPath:   "/domainConfigurations/{domainConfigurationName}",
	}

	if input == nil {
		input = &DescribeDomainConfigurationInput{}
	}

	req := c.newRequest(op, input, &DescribeDomainConfigurationOutput{})
	return DescribeDomainConfigurationRequest{Request: req, Input: input, Copy: c.DescribeDomainConfigurationRequest}
}

// DescribeDomainConfigurationRequest is the request type for the
// DescribeDomainConfiguration API operation.
type DescribeDomainConfigurationRequest struct {
	*aws.Request
	Input *DescribeDomainConfigurationInput
	Copy  func(*DescribeDomainConfigurationInput) DescribeDomainConfigurationRequest
}

// Send marshals and sends the DescribeDomainConfiguration API request.
func (r DescribeDomainConfigurationRequest) Send(ctx context.Context) (*DescribeDomainConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeDomainConfigurationResponse{
		DescribeDomainConfigurationOutput: r.Request.Data.(*DescribeDomainConfigurationOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeDomainConfigurationResponse is the response type for the
// DescribeDomainConfiguration API operation.
type DescribeDomainConfigurationResponse struct {
	*DescribeDomainConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeDomainConfiguration request.
func (r *DescribeDomainConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
