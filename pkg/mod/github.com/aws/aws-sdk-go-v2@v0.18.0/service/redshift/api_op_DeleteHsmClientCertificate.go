// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

type DeleteHsmClientCertificateInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the HSM client certificate to be deleted.
	//
	// HsmClientCertificateIdentifier is a required field
	HsmClientCertificateIdentifier *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteHsmClientCertificateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteHsmClientCertificateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteHsmClientCertificateInput"}

	if s.HsmClientCertificateIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("HsmClientCertificateIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteHsmClientCertificateOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteHsmClientCertificateOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteHsmClientCertificate = "DeleteHsmClientCertificate"

// DeleteHsmClientCertificateRequest returns a request value for making API operation for
// Amazon Redshift.
//
// Deletes the specified HSM client certificate.
//
//    // Example sending a request using DeleteHsmClientCertificateRequest.
//    req := client.DeleteHsmClientCertificateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/DeleteHsmClientCertificate
func (c *Client) DeleteHsmClientCertificateRequest(input *DeleteHsmClientCertificateInput) DeleteHsmClientCertificateRequest {
	op := &aws.Operation{
		Name:       opDeleteHsmClientCertificate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteHsmClientCertificateInput{}
	}

	req := c.newRequest(op, input, &DeleteHsmClientCertificateOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteHsmClientCertificateRequest{Request: req, Input: input, Copy: c.DeleteHsmClientCertificateRequest}
}

// DeleteHsmClientCertificateRequest is the request type for the
// DeleteHsmClientCertificate API operation.
type DeleteHsmClientCertificateRequest struct {
	*aws.Request
	Input *DeleteHsmClientCertificateInput
	Copy  func(*DeleteHsmClientCertificateInput) DeleteHsmClientCertificateRequest
}

// Send marshals and sends the DeleteHsmClientCertificate API request.
func (r DeleteHsmClientCertificateRequest) Send(ctx context.Context) (*DeleteHsmClientCertificateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteHsmClientCertificateResponse{
		DeleteHsmClientCertificateOutput: r.Request.Data.(*DeleteHsmClientCertificateOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteHsmClientCertificateResponse is the response type for the
// DeleteHsmClientCertificate API operation.
type DeleteHsmClientCertificateResponse struct {
	*DeleteHsmClientCertificateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteHsmClientCertificate request.
func (r *DeleteHsmClientCertificateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
