// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateNatGatewayInput struct {
	_ struct{} `type:"structure"`

	// The allocation ID of an Elastic IP address to associate with the NAT gateway.
	// If the Elastic IP address is associated with another resource, you must first
	// disassociate it.
	//
	// AllocationId is a required field
	AllocationId *string `type:"string" required:"true"`

	// Unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. For more information, see How to Ensure Idempotency (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html).
	//
	// Constraint: Maximum 64 ASCII characters.
	ClientToken *string `type:"string"`

	// The subnet in which to create the NAT gateway.
	//
	// SubnetId is a required field
	SubnetId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateNatGatewayInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateNatGatewayInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateNatGatewayInput"}

	if s.AllocationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AllocationId"))
	}

	if s.SubnetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SubnetId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateNatGatewayOutput struct {
	_ struct{} `type:"structure"`

	// Unique, case-sensitive identifier to ensure the idempotency of the request.
	// Only returned if a client token was provided in the request.
	ClientToken *string `locationName:"clientToken" type:"string"`

	// Information about the NAT gateway.
	NatGateway *NatGateway `locationName:"natGateway" type:"structure"`
}

// String returns the string representation
func (s CreateNatGatewayOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateNatGateway = "CreateNatGateway"

// CreateNatGatewayRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Creates a NAT gateway in the specified public subnet. This action creates
// a network interface in the specified subnet with a private IP address from
// the IP address range of the subnet. Internet-bound traffic from a private
// subnet can be routed to the NAT gateway, therefore enabling instances in
// the private subnet to connect to the internet. For more information, see
// NAT Gateways (https://docs.aws.amazon.com/vpc/latest/userguide/vpc-nat-gateway.html)
// in the Amazon Virtual Private Cloud User Guide.
//
//    // Example sending a request using CreateNatGatewayRequest.
//    req := client.CreateNatGatewayRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateNatGateway
func (c *Client) CreateNatGatewayRequest(input *CreateNatGatewayInput) CreateNatGatewayRequest {
	op := &aws.Operation{
		Name:       opCreateNatGateway,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateNatGatewayInput{}
	}

	req := c.newRequest(op, input, &CreateNatGatewayOutput{})
	return CreateNatGatewayRequest{Request: req, Input: input, Copy: c.CreateNatGatewayRequest}
}

// CreateNatGatewayRequest is the request type for the
// CreateNatGateway API operation.
type CreateNatGatewayRequest struct {
	*aws.Request
	Input *CreateNatGatewayInput
	Copy  func(*CreateNatGatewayInput) CreateNatGatewayRequest
}

// Send marshals and sends the CreateNatGateway API request.
func (r CreateNatGatewayRequest) Send(ctx context.Context) (*CreateNatGatewayResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateNatGatewayResponse{
		CreateNatGatewayOutput: r.Request.Data.(*CreateNatGatewayOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateNatGatewayResponse is the response type for the
// CreateNatGateway API operation.
type CreateNatGatewayResponse struct {
	*CreateNatGatewayOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateNatGateway request.
func (r *CreateNatGatewayResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
