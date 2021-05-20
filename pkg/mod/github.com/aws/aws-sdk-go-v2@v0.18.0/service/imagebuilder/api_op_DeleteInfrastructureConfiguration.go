// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package imagebuilder

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DeleteInfrastructureConfigurationInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the infrastructure configuration to delete.
	//
	// InfrastructureConfigurationArn is a required field
	InfrastructureConfigurationArn *string `location:"querystring" locationName:"infrastructureConfigurationArn" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteInfrastructureConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteInfrastructureConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteInfrastructureConfigurationInput"}

	if s.InfrastructureConfigurationArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("InfrastructureConfigurationArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteInfrastructureConfigurationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.InfrastructureConfigurationArn != nil {
		v := *s.InfrastructureConfigurationArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "infrastructureConfigurationArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteInfrastructureConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the infrastructure configuration that was
	// deleted.
	InfrastructureConfigurationArn *string `locationName:"infrastructureConfigurationArn" type:"string"`

	// The request ID that uniquely identifies this request.
	RequestId *string `locationName:"requestId" min:"1" type:"string"`
}

// String returns the string representation
func (s DeleteInfrastructureConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteInfrastructureConfigurationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.InfrastructureConfigurationArn != nil {
		v := *s.InfrastructureConfigurationArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "infrastructureConfigurationArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "requestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opDeleteInfrastructureConfiguration = "DeleteInfrastructureConfiguration"

// DeleteInfrastructureConfigurationRequest returns a request value for making API operation for
// EC2 Image Builder.
//
// Deletes an infrastructure configuration.
//
//    // Example sending a request using DeleteInfrastructureConfigurationRequest.
//    req := client.DeleteInfrastructureConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/imagebuilder-2019-12-02/DeleteInfrastructureConfiguration
func (c *Client) DeleteInfrastructureConfigurationRequest(input *DeleteInfrastructureConfigurationInput) DeleteInfrastructureConfigurationRequest {
	op := &aws.Operation{
		Name:       opDeleteInfrastructureConfiguration,
		HTTPMethod: "DELETE",
		HTTPPath:   "/DeleteInfrastructureConfiguration",
	}

	if input == nil {
		input = &DeleteInfrastructureConfigurationInput{}
	}

	req := c.newRequest(op, input, &DeleteInfrastructureConfigurationOutput{})
	return DeleteInfrastructureConfigurationRequest{Request: req, Input: input, Copy: c.DeleteInfrastructureConfigurationRequest}
}

// DeleteInfrastructureConfigurationRequest is the request type for the
// DeleteInfrastructureConfiguration API operation.
type DeleteInfrastructureConfigurationRequest struct {
	*aws.Request
	Input *DeleteInfrastructureConfigurationInput
	Copy  func(*DeleteInfrastructureConfigurationInput) DeleteInfrastructureConfigurationRequest
}

// Send marshals and sends the DeleteInfrastructureConfiguration API request.
func (r DeleteInfrastructureConfigurationRequest) Send(ctx context.Context) (*DeleteInfrastructureConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteInfrastructureConfigurationResponse{
		DeleteInfrastructureConfigurationOutput: r.Request.Data.(*DeleteInfrastructureConfigurationOutput),
		response:                                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteInfrastructureConfigurationResponse is the response type for the
// DeleteInfrastructureConfiguration API operation.
type DeleteInfrastructureConfigurationResponse struct {
	*DeleteInfrastructureConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteInfrastructureConfiguration request.
func (r *DeleteInfrastructureConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
