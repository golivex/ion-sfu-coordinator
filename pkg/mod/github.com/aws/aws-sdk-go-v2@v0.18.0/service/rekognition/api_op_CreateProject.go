// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rekognition

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateProjectInput struct {
	_ struct{} `type:"structure"`

	// The name of the project to create.
	//
	// ProjectName is a required field
	ProjectName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateProjectInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateProjectInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateProjectInput"}

	if s.ProjectName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProjectName"))
	}
	if s.ProjectName != nil && len(*s.ProjectName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProjectName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateProjectOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the new project. You can use the ARN to
	// configure IAM access to the project.
	ProjectArn *string `min:"20" type:"string"`
}

// String returns the string representation
func (s CreateProjectOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateProject = "CreateProject"

// CreateProjectRequest returns a request value for making API operation for
// Amazon Rekognition.
//
// Creates a new Amazon Rekognition Custom Labels project. A project is a logical
// grouping of resources (images, Labels, models) and operations (training,
// evaluation and detection).
//
// This operation requires permissions to perform the rekognition:CreateProject
// action.
//
//    // Example sending a request using CreateProjectRequest.
//    req := client.CreateProjectRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CreateProjectRequest(input *CreateProjectInput) CreateProjectRequest {
	op := &aws.Operation{
		Name:       opCreateProject,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateProjectInput{}
	}

	req := c.newRequest(op, input, &CreateProjectOutput{})
	return CreateProjectRequest{Request: req, Input: input, Copy: c.CreateProjectRequest}
}

// CreateProjectRequest is the request type for the
// CreateProject API operation.
type CreateProjectRequest struct {
	*aws.Request
	Input *CreateProjectInput
	Copy  func(*CreateProjectInput) CreateProjectRequest
}

// Send marshals and sends the CreateProject API request.
func (r CreateProjectRequest) Send(ctx context.Context) (*CreateProjectResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateProjectResponse{
		CreateProjectOutput: r.Request.Data.(*CreateProjectOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateProjectResponse is the response type for the
// CreateProject API operation.
type CreateProjectResponse struct {
	*CreateProjectOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateProject request.
func (r *CreateProjectResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
