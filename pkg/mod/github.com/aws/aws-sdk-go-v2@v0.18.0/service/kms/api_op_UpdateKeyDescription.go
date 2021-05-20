// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type UpdateKeyDescriptionInput struct {
	_ struct{} `type:"structure"`

	// New description for the CMK.
	//
	// Description is a required field
	Description *string `type:"string" required:"true"`

	// A unique identifier for the customer master key (CMK).
	//
	// Specify the key ID or the Amazon Resource Name (ARN) of the CMK.
	//
	// For example:
	//
	//    * Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//    * Key ARN: arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// To get the key ID and key ARN for a CMK, use ListKeys or DescribeKey.
	//
	// KeyId is a required field
	KeyId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateKeyDescriptionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateKeyDescriptionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateKeyDescriptionInput"}

	if s.Description == nil {
		invalidParams.Add(aws.NewErrParamRequired("Description"))
	}

	if s.KeyId == nil {
		invalidParams.Add(aws.NewErrParamRequired("KeyId"))
	}
	if s.KeyId != nil && len(*s.KeyId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("KeyId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateKeyDescriptionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateKeyDescriptionOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateKeyDescription = "UpdateKeyDescription"

// UpdateKeyDescriptionRequest returns a request value for making API operation for
// AWS Key Management Service.
//
// Updates the description of a customer master key (CMK). To see the description
// of a CMK, use DescribeKey.
//
// You cannot perform this operation on a CMK in a different AWS account.
//
// The CMK that you use for this operation must be in a compatible key state.
// For details, see How Key State Affects Use of a Customer Master Key (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html)
// in the AWS Key Management Service Developer Guide.
//
//    // Example sending a request using UpdateKeyDescriptionRequest.
//    req := client.UpdateKeyDescriptionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kms-2014-11-01/UpdateKeyDescription
func (c *Client) UpdateKeyDescriptionRequest(input *UpdateKeyDescriptionInput) UpdateKeyDescriptionRequest {
	op := &aws.Operation{
		Name:       opUpdateKeyDescription,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateKeyDescriptionInput{}
	}

	req := c.newRequest(op, input, &UpdateKeyDescriptionOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return UpdateKeyDescriptionRequest{Request: req, Input: input, Copy: c.UpdateKeyDescriptionRequest}
}

// UpdateKeyDescriptionRequest is the request type for the
// UpdateKeyDescription API operation.
type UpdateKeyDescriptionRequest struct {
	*aws.Request
	Input *UpdateKeyDescriptionInput
	Copy  func(*UpdateKeyDescriptionInput) UpdateKeyDescriptionRequest
}

// Send marshals and sends the UpdateKeyDescription API request.
func (r UpdateKeyDescriptionRequest) Send(ctx context.Context) (*UpdateKeyDescriptionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateKeyDescriptionResponse{
		UpdateKeyDescriptionOutput: r.Request.Data.(*UpdateKeyDescriptionOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateKeyDescriptionResponse is the response type for the
// UpdateKeyDescription API operation.
type UpdateKeyDescriptionResponse struct {
	*UpdateKeyDescriptionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateKeyDescription request.
func (r *UpdateKeyDescriptionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
