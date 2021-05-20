// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type DeleteDomainInput struct {
	_ struct{} `type:"structure"`

	// The domain ID.
	//
	// DomainId is a required field
	DomainId *string `type:"string" required:"true"`

	// The retention policy for this domain, which specifies which resources will
	// be retained after the Domain is deleted. By default, all resources are retained
	// (not automatically deleted).
	RetentionPolicy *RetentionPolicy `type:"structure"`
}

// String returns the string representation
func (s DeleteDomainInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDomainInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteDomainInput"}

	if s.DomainId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteDomainOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteDomainOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteDomain = "DeleteDomain"

// DeleteDomainRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Used to delete a domain. If you on-boarded with IAM mode, you will need to
// delete your domain to on-board again using SSO. Use with caution. All of
// the members of the domain will lose access to their EFS volume, including
// data, notebooks, and other artifacts.
//
//    // Example sending a request using DeleteDomainRequest.
//    req := client.DeleteDomainRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/DeleteDomain
func (c *Client) DeleteDomainRequest(input *DeleteDomainInput) DeleteDomainRequest {
	op := &aws.Operation{
		Name:       opDeleteDomain,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteDomainInput{}
	}

	req := c.newRequest(op, input, &DeleteDomainOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteDomainRequest{Request: req, Input: input, Copy: c.DeleteDomainRequest}
}

// DeleteDomainRequest is the request type for the
// DeleteDomain API operation.
type DeleteDomainRequest struct {
	*aws.Request
	Input *DeleteDomainInput
	Copy  func(*DeleteDomainInput) DeleteDomainRequest
}

// Send marshals and sends the DeleteDomain API request.
func (r DeleteDomainRequest) Send(ctx context.Context) (*DeleteDomainResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteDomainResponse{
		DeleteDomainOutput: r.Request.Data.(*DeleteDomainOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteDomainResponse is the response type for the
// DeleteDomain API operation.
type DeleteDomainResponse struct {
	*DeleteDomainOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteDomain request.
func (r *DeleteDomainResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
