// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package inspector

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeCrossAccountAccessRoleInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DescribeCrossAccountAccessRoleInput) String() string {
	return awsutil.Prettify(s)
}

type DescribeCrossAccountAccessRoleOutput struct {
	_ struct{} `type:"structure"`

	// The date when the cross-account access role was registered.
	//
	// RegisteredAt is a required field
	RegisteredAt *time.Time `locationName:"registeredAt" type:"timestamp" required:"true"`

	// The ARN that specifies the IAM role that Amazon Inspector uses to access
	// your AWS account.
	//
	// RoleArn is a required field
	RoleArn *string `locationName:"roleArn" min:"1" type:"string" required:"true"`

	// A Boolean value that specifies whether the IAM role has the necessary policies
	// attached to enable Amazon Inspector to access your AWS account.
	//
	// Valid is a required field
	Valid *bool `locationName:"valid" type:"boolean" required:"true"`
}

// String returns the string representation
func (s DescribeCrossAccountAccessRoleOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeCrossAccountAccessRole = "DescribeCrossAccountAccessRole"

// DescribeCrossAccountAccessRoleRequest returns a request value for making API operation for
// Amazon Inspector.
//
// Describes the IAM role that enables Amazon Inspector to access your AWS account.
//
//    // Example sending a request using DescribeCrossAccountAccessRoleRequest.
//    req := client.DescribeCrossAccountAccessRoleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/inspector-2016-02-16/DescribeCrossAccountAccessRole
func (c *Client) DescribeCrossAccountAccessRoleRequest(input *DescribeCrossAccountAccessRoleInput) DescribeCrossAccountAccessRoleRequest {
	op := &aws.Operation{
		Name:       opDescribeCrossAccountAccessRole,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeCrossAccountAccessRoleInput{}
	}

	req := c.newRequest(op, input, &DescribeCrossAccountAccessRoleOutput{})
	return DescribeCrossAccountAccessRoleRequest{Request: req, Input: input, Copy: c.DescribeCrossAccountAccessRoleRequest}
}

// DescribeCrossAccountAccessRoleRequest is the request type for the
// DescribeCrossAccountAccessRole API operation.
type DescribeCrossAccountAccessRoleRequest struct {
	*aws.Request
	Input *DescribeCrossAccountAccessRoleInput
	Copy  func(*DescribeCrossAccountAccessRoleInput) DescribeCrossAccountAccessRoleRequest
}

// Send marshals and sends the DescribeCrossAccountAccessRole API request.
func (r DescribeCrossAccountAccessRoleRequest) Send(ctx context.Context) (*DescribeCrossAccountAccessRoleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeCrossAccountAccessRoleResponse{
		DescribeCrossAccountAccessRoleOutput: r.Request.Data.(*DescribeCrossAccountAccessRoleOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeCrossAccountAccessRoleResponse is the response type for the
// DescribeCrossAccountAccessRole API operation.
type DescribeCrossAccountAccessRoleResponse struct {
	*DescribeCrossAccountAccessRoleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeCrossAccountAccessRole request.
func (r *DescribeCrossAccountAccessRoleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
