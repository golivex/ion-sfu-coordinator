// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package migrationhub

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateProgressUpdateStreamInput struct {
	_ struct{} `type:"structure"`

	// Optional boolean flag to indicate whether any effect should take place. Used
	// to test if the caller has permission to make the call.
	DryRun *bool `type:"boolean"`

	// The name of the ProgressUpdateStream. Do not store personal data in this
	// field.
	//
	// ProgressUpdateStreamName is a required field
	ProgressUpdateStreamName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateProgressUpdateStreamInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateProgressUpdateStreamInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateProgressUpdateStreamInput"}

	if s.ProgressUpdateStreamName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProgressUpdateStreamName"))
	}
	if s.ProgressUpdateStreamName != nil && len(*s.ProgressUpdateStreamName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProgressUpdateStreamName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateProgressUpdateStreamOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateProgressUpdateStreamOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateProgressUpdateStream = "CreateProgressUpdateStream"

// CreateProgressUpdateStreamRequest returns a request value for making API operation for
// AWS Migration Hub.
//
// Creates a progress update stream which is an AWS resource used for access
// control as well as a namespace for migration task names that is implicitly
// linked to your AWS account. It must uniquely identify the migration tool
// as it is used for all updates made by the tool; however, it does not need
// to be unique for each AWS account because it is scoped to the AWS account.
//
//    // Example sending a request using CreateProgressUpdateStreamRequest.
//    req := client.CreateProgressUpdateStreamRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/AWSMigrationHub-2017-05-31/CreateProgressUpdateStream
func (c *Client) CreateProgressUpdateStreamRequest(input *CreateProgressUpdateStreamInput) CreateProgressUpdateStreamRequest {
	op := &aws.Operation{
		Name:       opCreateProgressUpdateStream,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateProgressUpdateStreamInput{}
	}

	req := c.newRequest(op, input, &CreateProgressUpdateStreamOutput{})
	return CreateProgressUpdateStreamRequest{Request: req, Input: input, Copy: c.CreateProgressUpdateStreamRequest}
}

// CreateProgressUpdateStreamRequest is the request type for the
// CreateProgressUpdateStream API operation.
type CreateProgressUpdateStreamRequest struct {
	*aws.Request
	Input *CreateProgressUpdateStreamInput
	Copy  func(*CreateProgressUpdateStreamInput) CreateProgressUpdateStreamRequest
}

// Send marshals and sends the CreateProgressUpdateStream API request.
func (r CreateProgressUpdateStreamRequest) Send(ctx context.Context) (*CreateProgressUpdateStreamResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateProgressUpdateStreamResponse{
		CreateProgressUpdateStreamOutput: r.Request.Data.(*CreateProgressUpdateStreamOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateProgressUpdateStreamResponse is the response type for the
// CreateProgressUpdateStream API operation.
type CreateProgressUpdateStreamResponse struct {
	*CreateProgressUpdateStreamOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateProgressUpdateStream request.
func (r *CreateProgressUpdateStreamResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
