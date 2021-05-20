// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workspaces

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateTagsInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the WorkSpaces resource. The supported resource types are
	// WorkSpaces, registered directories, images, custom bundles, and IP access
	// control groups.
	//
	// ResourceId is a required field
	ResourceId *string `min:"1" type:"string" required:"true"`

	// The tags. Each WorkSpaces resource can have a maximum of 50 tags.
	//
	// Tags is a required field
	Tags []Tag `type:"list" required:"true"`
}

// String returns the string representation
func (s CreateTagsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateTagsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateTagsInput"}

	if s.ResourceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceId"))
	}
	if s.ResourceId != nil && len(*s.ResourceId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceId", 1))
	}

	if s.Tags == nil {
		invalidParams.Add(aws.NewErrParamRequired("Tags"))
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateTagsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateTagsOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateTags = "CreateTags"

// CreateTagsRequest returns a request value for making API operation for
// Amazon WorkSpaces.
//
// Creates the specified tags for the specified WorkSpaces resource.
//
//    // Example sending a request using CreateTagsRequest.
//    req := client.CreateTagsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workspaces-2015-04-08/CreateTags
func (c *Client) CreateTagsRequest(input *CreateTagsInput) CreateTagsRequest {
	op := &aws.Operation{
		Name:       opCreateTags,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateTagsInput{}
	}

	req := c.newRequest(op, input, &CreateTagsOutput{})
	return CreateTagsRequest{Request: req, Input: input, Copy: c.CreateTagsRequest}
}

// CreateTagsRequest is the request type for the
// CreateTags API operation.
type CreateTagsRequest struct {
	*aws.Request
	Input *CreateTagsInput
	Copy  func(*CreateTagsInput) CreateTagsRequest
}

// Send marshals and sends the CreateTags API request.
func (r CreateTagsRequest) Send(ctx context.Context) (*CreateTagsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateTagsResponse{
		CreateTagsOutput: r.Request.Data.(*CreateTagsOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateTagsResponse is the response type for the
// CreateTags API operation.
type CreateTagsResponse struct {
	*CreateTagsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateTags request.
func (r *CreateTagsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
