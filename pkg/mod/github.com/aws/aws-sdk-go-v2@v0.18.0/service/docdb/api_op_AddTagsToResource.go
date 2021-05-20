// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package docdb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

// Represents the input to AddTagsToResource.
type AddTagsToResourceInput struct {
	_ struct{} `type:"structure"`

	// The Amazon DocumentDB resource that the tags are added to. This value is
	// an Amazon Resource Name (ARN).
	//
	// ResourceName is a required field
	ResourceName *string `type:"string" required:"true"`

	// The tags to be assigned to the Amazon DocumentDB resource.
	//
	// Tags is a required field
	Tags []Tag `locationNameList:"Tag" type:"list" required:"true"`
}

// String returns the string representation
func (s AddTagsToResourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AddTagsToResourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AddTagsToResourceInput"}

	if s.ResourceName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceName"))
	}

	if s.Tags == nil {
		invalidParams.Add(aws.NewErrParamRequired("Tags"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AddTagsToResourceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AddTagsToResourceOutput) String() string {
	return awsutil.Prettify(s)
}

const opAddTagsToResource = "AddTagsToResource"

// AddTagsToResourceRequest returns a request value for making API operation for
// Amazon DocumentDB with MongoDB compatibility.
//
// Adds metadata tags to an Amazon DocumentDB resource. You can use these tags
// with cost allocation reporting to track costs that are associated with Amazon
// DocumentDB resources. or in a Condition statement in an AWS Identity and
// Access Management (IAM) policy for Amazon DocumentDB.
//
//    // Example sending a request using AddTagsToResourceRequest.
//    req := client.AddTagsToResourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/docdb-2014-10-31/AddTagsToResource
func (c *Client) AddTagsToResourceRequest(input *AddTagsToResourceInput) AddTagsToResourceRequest {
	op := &aws.Operation{
		Name:       opAddTagsToResource,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AddTagsToResourceInput{}
	}

	req := c.newRequest(op, input, &AddTagsToResourceOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return AddTagsToResourceRequest{Request: req, Input: input, Copy: c.AddTagsToResourceRequest}
}

// AddTagsToResourceRequest is the request type for the
// AddTagsToResource API operation.
type AddTagsToResourceRequest struct {
	*aws.Request
	Input *AddTagsToResourceInput
	Copy  func(*AddTagsToResourceInput) AddTagsToResourceRequest
}

// Send marshals and sends the AddTagsToResource API request.
func (r AddTagsToResourceRequest) Send(ctx context.Context) (*AddTagsToResourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AddTagsToResourceResponse{
		AddTagsToResourceOutput: r.Request.Data.(*AddTagsToResourceOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AddTagsToResourceResponse is the response type for the
// AddTagsToResource API operation.
type AddTagsToResourceResponse struct {
	*AddTagsToResourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AddTagsToResource request.
func (r *AddTagsToResourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
