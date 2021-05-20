// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package acmpca

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type UpdateCertificateAuthorityInput struct {
	_ struct{} `type:"structure"`

	// Amazon Resource Name (ARN) of the private CA that issued the certificate
	// to be revoked. This must be of the form:
	//
	// arn:aws:acm-pca:region:account:certificate-authority/12345678-1234-1234-1234-123456789012
	//
	// CertificateAuthorityArn is a required field
	CertificateAuthorityArn *string `min:"5" type:"string" required:"true"`

	// Revocation information for your private CA.
	RevocationConfiguration *RevocationConfiguration `type:"structure"`

	// Status of your private CA.
	Status CertificateAuthorityStatus `type:"string" enum:"true"`
}

// String returns the string representation
func (s UpdateCertificateAuthorityInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateCertificateAuthorityInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateCertificateAuthorityInput"}

	if s.CertificateAuthorityArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("CertificateAuthorityArn"))
	}
	if s.CertificateAuthorityArn != nil && len(*s.CertificateAuthorityArn) < 5 {
		invalidParams.Add(aws.NewErrParamMinLen("CertificateAuthorityArn", 5))
	}
	if s.RevocationConfiguration != nil {
		if err := s.RevocationConfiguration.Validate(); err != nil {
			invalidParams.AddNested("RevocationConfiguration", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateCertificateAuthorityOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateCertificateAuthorityOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateCertificateAuthority = "UpdateCertificateAuthority"

// UpdateCertificateAuthorityRequest returns a request value for making API operation for
// AWS Certificate Manager Private Certificate Authority.
//
// Updates the status or configuration of a private certificate authority (CA).
// Your private CA must be in the ACTIVE or DISABLED state before you can update
// it. You can disable a private CA that is in the ACTIVE state or make a CA
// that is in the DISABLED state active again.
//
//    // Example sending a request using UpdateCertificateAuthorityRequest.
//    req := client.UpdateCertificateAuthorityRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/acm-pca-2017-08-22/UpdateCertificateAuthority
func (c *Client) UpdateCertificateAuthorityRequest(input *UpdateCertificateAuthorityInput) UpdateCertificateAuthorityRequest {
	op := &aws.Operation{
		Name:       opUpdateCertificateAuthority,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateCertificateAuthorityInput{}
	}

	req := c.newRequest(op, input, &UpdateCertificateAuthorityOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return UpdateCertificateAuthorityRequest{Request: req, Input: input, Copy: c.UpdateCertificateAuthorityRequest}
}

// UpdateCertificateAuthorityRequest is the request type for the
// UpdateCertificateAuthority API operation.
type UpdateCertificateAuthorityRequest struct {
	*aws.Request
	Input *UpdateCertificateAuthorityInput
	Copy  func(*UpdateCertificateAuthorityInput) UpdateCertificateAuthorityRequest
}

// Send marshals and sends the UpdateCertificateAuthority API request.
func (r UpdateCertificateAuthorityRequest) Send(ctx context.Context) (*UpdateCertificateAuthorityResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateCertificateAuthorityResponse{
		UpdateCertificateAuthorityOutput: r.Request.Data.(*UpdateCertificateAuthorityOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateCertificateAuthorityResponse is the response type for the
// UpdateCertificateAuthority API operation.
type UpdateCertificateAuthorityResponse struct {
	*UpdateCertificateAuthorityOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateCertificateAuthority request.
func (r *UpdateCertificateAuthorityResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
