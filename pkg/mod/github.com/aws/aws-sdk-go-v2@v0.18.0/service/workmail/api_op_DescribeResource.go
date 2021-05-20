// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workmail

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeResourceInput struct {
	_ struct{} `type:"structure"`

	// The identifier associated with the organization for which the resource is
	// described.
	//
	// OrganizationId is a required field
	OrganizationId *string `type:"string" required:"true"`

	// The identifier of the resource to be described.
	//
	// ResourceId is a required field
	ResourceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeResourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeResourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeResourceInput"}

	if s.OrganizationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("OrganizationId"))
	}

	if s.ResourceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeResourceOutput struct {
	_ struct{} `type:"structure"`

	// The booking options for the described resource.
	BookingOptions *BookingOptions `type:"structure"`

	// The date and time when a resource was disabled from WorkMail, in UNIX epoch
	// time format.
	DisabledDate *time.Time `type:"timestamp"`

	// The email of the described resource.
	Email *string `min:"1" type:"string"`

	// The date and time when a resource was enabled for WorkMail, in UNIX epoch
	// time format.
	EnabledDate *time.Time `type:"timestamp"`

	// The name of the described resource.
	Name *string `min:"1" type:"string"`

	// The identifier of the described resource.
	ResourceId *string `type:"string"`

	// The state of the resource: enabled (registered to Amazon WorkMail) or disabled
	// (deregistered or never registered to WorkMail).
	State EntityState `type:"string" enum:"true"`

	// The type of the described resource.
	Type ResourceType `type:"string" enum:"true"`
}

// String returns the string representation
func (s DescribeResourceOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeResource = "DescribeResource"

// DescribeResourceRequest returns a request value for making API operation for
// Amazon WorkMail.
//
// Returns the data available for the resource.
//
//    // Example sending a request using DescribeResourceRequest.
//    req := client.DescribeResourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workmail-2017-10-01/DescribeResource
func (c *Client) DescribeResourceRequest(input *DescribeResourceInput) DescribeResourceRequest {
	op := &aws.Operation{
		Name:       opDescribeResource,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeResourceInput{}
	}

	req := c.newRequest(op, input, &DescribeResourceOutput{})
	return DescribeResourceRequest{Request: req, Input: input, Copy: c.DescribeResourceRequest}
}

// DescribeResourceRequest is the request type for the
// DescribeResource API operation.
type DescribeResourceRequest struct {
	*aws.Request
	Input *DescribeResourceInput
	Copy  func(*DescribeResourceInput) DescribeResourceRequest
}

// Send marshals and sends the DescribeResource API request.
func (r DescribeResourceRequest) Send(ctx context.Context) (*DescribeResourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeResourceResponse{
		DescribeResourceOutput: r.Request.Data.(*DescribeResourceOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeResourceResponse is the response type for the
// DescribeResource API operation.
type DescribeResourceResponse struct {
	*DescribeResourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeResource request.
func (r *DescribeResourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
