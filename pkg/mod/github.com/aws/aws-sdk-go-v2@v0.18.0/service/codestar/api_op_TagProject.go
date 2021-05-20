// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codestar

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type TagProjectInput struct {
	_ struct{} `type:"structure"`

	// The ID of the project you want to add a tag to.
	//
	// Id is a required field
	Id *string `locationName:"id" min:"2" type:"string" required:"true"`

	// The tags you want to add to the project.
	//
	// Tags is a required field
	Tags map[string]string `locationName:"tags" type:"map" required:"true"`
}

// String returns the string representation
func (s TagProjectInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TagProjectInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "TagProjectInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}
	if s.Id != nil && len(*s.Id) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("Id", 2))
	}

	if s.Tags == nil {
		invalidParams.Add(aws.NewErrParamRequired("Tags"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type TagProjectOutput struct {
	_ struct{} `type:"structure"`

	// The tags for the project.
	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s TagProjectOutput) String() string {
	return awsutil.Prettify(s)
}

const opTagProject = "TagProject"

// TagProjectRequest returns a request value for making API operation for
// AWS CodeStar.
//
// Adds tags to a project.
//
//    // Example sending a request using TagProjectRequest.
//    req := client.TagProjectRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codestar-2017-04-19/TagProject
func (c *Client) TagProjectRequest(input *TagProjectInput) TagProjectRequest {
	op := &aws.Operation{
		Name:       opTagProject,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &TagProjectInput{}
	}

	req := c.newRequest(op, input, &TagProjectOutput{})
	return TagProjectRequest{Request: req, Input: input, Copy: c.TagProjectRequest}
}

// TagProjectRequest is the request type for the
// TagProject API operation.
type TagProjectRequest struct {
	*aws.Request
	Input *TagProjectInput
	Copy  func(*TagProjectInput) TagProjectRequest
}

// Send marshals and sends the TagProject API request.
func (r TagProjectRequest) Send(ctx context.Context) (*TagProjectResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &TagProjectResponse{
		TagProjectOutput: r.Request.Data.(*TagProjectOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// TagProjectResponse is the response type for the
// TagProject API operation.
type TagProjectResponse struct {
	*TagProjectOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// TagProject request.
func (r *TagProjectResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
