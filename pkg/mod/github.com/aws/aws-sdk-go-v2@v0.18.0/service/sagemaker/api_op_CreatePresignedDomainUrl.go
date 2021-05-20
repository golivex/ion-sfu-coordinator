// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreatePresignedDomainUrlInput struct {
	_ struct{} `type:"structure"`

	// The domain ID.
	//
	// DomainId is a required field
	DomainId *string `type:"string" required:"true"`

	// The session expiration duration in seconds.
	SessionExpirationDurationInSeconds *int64 `min:"1800" type:"integer"`

	// The name of the UserProfile to sign-in as.
	//
	// UserProfileName is a required field
	UserProfileName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreatePresignedDomainUrlInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreatePresignedDomainUrlInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreatePresignedDomainUrlInput"}

	if s.DomainId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainId"))
	}
	if s.SessionExpirationDurationInSeconds != nil && *s.SessionExpirationDurationInSeconds < 1800 {
		invalidParams.Add(aws.NewErrParamMinValue("SessionExpirationDurationInSeconds", 1800))
	}

	if s.UserProfileName == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserProfileName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreatePresignedDomainUrlOutput struct {
	_ struct{} `type:"structure"`

	// The presigned URL.
	AuthorizedUrl *string `type:"string"`
}

// String returns the string representation
func (s CreatePresignedDomainUrlOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreatePresignedDomainUrl = "CreatePresignedDomainUrl"

// CreatePresignedDomainUrlRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Creates a URL for a specified UserProfile in a Domain. When accessed in a
// web browser, the user will be automatically signed in to Amazon SageMaker
// Amazon SageMaker Studio (Studio), and granted access to all of the Apps and
// files associated with that Amazon Elastic File System (EFS). This operation
// can only be called when AuthMode equals IAM.
//
//    // Example sending a request using CreatePresignedDomainUrlRequest.
//    req := client.CreatePresignedDomainUrlRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/CreatePresignedDomainUrl
func (c *Client) CreatePresignedDomainUrlRequest(input *CreatePresignedDomainUrlInput) CreatePresignedDomainUrlRequest {
	op := &aws.Operation{
		Name:       opCreatePresignedDomainUrl,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreatePresignedDomainUrlInput{}
	}

	req := c.newRequest(op, input, &CreatePresignedDomainUrlOutput{})
	return CreatePresignedDomainUrlRequest{Request: req, Input: input, Copy: c.CreatePresignedDomainUrlRequest}
}

// CreatePresignedDomainUrlRequest is the request type for the
// CreatePresignedDomainUrl API operation.
type CreatePresignedDomainUrlRequest struct {
	*aws.Request
	Input *CreatePresignedDomainUrlInput
	Copy  func(*CreatePresignedDomainUrlInput) CreatePresignedDomainUrlRequest
}

// Send marshals and sends the CreatePresignedDomainUrl API request.
func (r CreatePresignedDomainUrlRequest) Send(ctx context.Context) (*CreatePresignedDomainUrlResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreatePresignedDomainUrlResponse{
		CreatePresignedDomainUrlOutput: r.Request.Data.(*CreatePresignedDomainUrlOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreatePresignedDomainUrlResponse is the response type for the
// CreatePresignedDomainUrl API operation.
type CreatePresignedDomainUrlResponse struct {
	*CreatePresignedDomainUrlOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreatePresignedDomainUrl request.
func (r *CreatePresignedDomainUrlResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
