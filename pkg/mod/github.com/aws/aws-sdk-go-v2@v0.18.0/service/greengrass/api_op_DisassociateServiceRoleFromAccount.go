// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DisassociateServiceRoleFromAccountInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DisassociateServiceRoleFromAccountInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DisassociateServiceRoleFromAccountInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	return nil
}

type DisassociateServiceRoleFromAccountOutput struct {
	_ struct{} `type:"structure"`

	// The time when the service role was disassociated from the account.
	DisassociatedAt *string `type:"string"`
}

// String returns the string representation
func (s DisassociateServiceRoleFromAccountOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DisassociateServiceRoleFromAccountOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DisassociatedAt != nil {
		v := *s.DisassociatedAt

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DisassociatedAt", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opDisassociateServiceRoleFromAccount = "DisassociateServiceRoleFromAccount"

// DisassociateServiceRoleFromAccountRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Disassociates the service role from your account. Without a service role,
// deployments will not work.
//
//    // Example sending a request using DisassociateServiceRoleFromAccountRequest.
//    req := client.DisassociateServiceRoleFromAccountRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/DisassociateServiceRoleFromAccount
func (c *Client) DisassociateServiceRoleFromAccountRequest(input *DisassociateServiceRoleFromAccountInput) DisassociateServiceRoleFromAccountRequest {
	op := &aws.Operation{
		Name:       opDisassociateServiceRoleFromAccount,
		HTTPMethod: "DELETE",
		HTTPPath:   "/greengrass/servicerole",
	}

	if input == nil {
		input = &DisassociateServiceRoleFromAccountInput{}
	}

	req := c.newRequest(op, input, &DisassociateServiceRoleFromAccountOutput{})
	return DisassociateServiceRoleFromAccountRequest{Request: req, Input: input, Copy: c.DisassociateServiceRoleFromAccountRequest}
}

// DisassociateServiceRoleFromAccountRequest is the request type for the
// DisassociateServiceRoleFromAccount API operation.
type DisassociateServiceRoleFromAccountRequest struct {
	*aws.Request
	Input *DisassociateServiceRoleFromAccountInput
	Copy  func(*DisassociateServiceRoleFromAccountInput) DisassociateServiceRoleFromAccountRequest
}

// Send marshals and sends the DisassociateServiceRoleFromAccount API request.
func (r DisassociateServiceRoleFromAccountRequest) Send(ctx context.Context) (*DisassociateServiceRoleFromAccountResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisassociateServiceRoleFromAccountResponse{
		DisassociateServiceRoleFromAccountOutput: r.Request.Data.(*DisassociateServiceRoleFromAccountOutput),
		response:                                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisassociateServiceRoleFromAccountResponse is the response type for the
// DisassociateServiceRoleFromAccount API operation.
type DisassociateServiceRoleFromAccountResponse struct {
	*DisassociateServiceRoleFromAccountOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisassociateServiceRoleFromAccount request.
func (r *DisassociateServiceRoleFromAccountResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
