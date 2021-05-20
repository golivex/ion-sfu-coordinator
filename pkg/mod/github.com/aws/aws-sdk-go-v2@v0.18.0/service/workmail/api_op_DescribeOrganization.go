// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workmail

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeOrganizationInput struct {
	_ struct{} `type:"structure"`

	// The identifier for the organization to be described.
	//
	// OrganizationId is a required field
	OrganizationId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeOrganizationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeOrganizationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeOrganizationInput"}

	if s.OrganizationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("OrganizationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeOrganizationOutput struct {
	_ struct{} `type:"structure"`

	// The alias for an organization.
	Alias *string `min:"1" type:"string"`

	// The date at which the organization became usable in the WorkMail context,
	// in UNIX epoch time format.
	CompletedDate *time.Time `type:"timestamp"`

	// The default mail domain associated with the organization.
	DefaultMailDomain *string `type:"string"`

	// The identifier for the directory associated with an Amazon WorkMail organization.
	DirectoryId *string `type:"string"`

	// The type of directory associated with the WorkMail organization.
	DirectoryType *string `type:"string"`

	// (Optional) The error message indicating if unexpected behavior was encountered
	// with regards to the organization.
	ErrorMessage *string `type:"string"`

	// The identifier of an organization.
	OrganizationId *string `type:"string"`

	// The state of an organization.
	State *string `type:"string"`
}

// String returns the string representation
func (s DescribeOrganizationOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeOrganization = "DescribeOrganization"

// DescribeOrganizationRequest returns a request value for making API operation for
// Amazon WorkMail.
//
// Provides more information regarding a given organization based on its identifier.
//
//    // Example sending a request using DescribeOrganizationRequest.
//    req := client.DescribeOrganizationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workmail-2017-10-01/DescribeOrganization
func (c *Client) DescribeOrganizationRequest(input *DescribeOrganizationInput) DescribeOrganizationRequest {
	op := &aws.Operation{
		Name:       opDescribeOrganization,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeOrganizationInput{}
	}

	req := c.newRequest(op, input, &DescribeOrganizationOutput{})
	return DescribeOrganizationRequest{Request: req, Input: input, Copy: c.DescribeOrganizationRequest}
}

// DescribeOrganizationRequest is the request type for the
// DescribeOrganization API operation.
type DescribeOrganizationRequest struct {
	*aws.Request
	Input *DescribeOrganizationInput
	Copy  func(*DescribeOrganizationInput) DescribeOrganizationRequest
}

// Send marshals and sends the DescribeOrganization API request.
func (r DescribeOrganizationRequest) Send(ctx context.Context) (*DescribeOrganizationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeOrganizationResponse{
		DescribeOrganizationOutput: r.Request.Data.(*DescribeOrganizationOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeOrganizationResponse is the response type for the
// DescribeOrganization API operation.
type DescribeOrganizationResponse struct {
	*DescribeOrganizationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeOrganization request.
func (r *DescribeOrganizationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
