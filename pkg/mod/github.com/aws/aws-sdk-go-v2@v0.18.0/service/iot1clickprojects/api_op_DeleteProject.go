// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot1clickprojects

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DeleteProjectInput struct {
	_ struct{} `type:"structure"`

	// The name of the empty project to delete.
	//
	// ProjectName is a required field
	ProjectName *string `location:"uri" locationName:"projectName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteProjectInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteProjectInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteProjectInput"}

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

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteProjectInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ProjectName != nil {
		v := *s.ProjectName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "projectName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteProjectOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteProjectOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteProjectOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteProject = "DeleteProject"

// DeleteProjectRequest returns a request value for making API operation for
// AWS IoT 1-Click Projects Service.
//
// Deletes a project. To delete a project, it must not have any placements associated
// with it.
//
// When you delete a project, all associated data becomes irretrievable.
//
//    // Example sending a request using DeleteProjectRequest.
//    req := client.DeleteProjectRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot1click-projects-2018-05-14/DeleteProject
func (c *Client) DeleteProjectRequest(input *DeleteProjectInput) DeleteProjectRequest {
	op := &aws.Operation{
		Name:       opDeleteProject,
		HTTPMethod: "DELETE",
		HTTPPath:   "/projects/{projectName}",
	}

	if input == nil {
		input = &DeleteProjectInput{}
	}

	req := c.newRequest(op, input, &DeleteProjectOutput{})
	return DeleteProjectRequest{Request: req, Input: input, Copy: c.DeleteProjectRequest}
}

// DeleteProjectRequest is the request type for the
// DeleteProject API operation.
type DeleteProjectRequest struct {
	*aws.Request
	Input *DeleteProjectInput
	Copy  func(*DeleteProjectInput) DeleteProjectRequest
}

// Send marshals and sends the DeleteProject API request.
func (r DeleteProjectRequest) Send(ctx context.Context) (*DeleteProjectResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteProjectResponse{
		DeleteProjectOutput: r.Request.Data.(*DeleteProjectOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteProjectResponse is the response type for the
// DeleteProject API operation.
type DeleteProjectResponse struct {
	*DeleteProjectOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteProject request.
func (r *DeleteProjectResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
