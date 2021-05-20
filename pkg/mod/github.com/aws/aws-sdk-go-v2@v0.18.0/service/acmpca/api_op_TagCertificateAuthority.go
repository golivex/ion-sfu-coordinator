// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package acmpca

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type TagCertificateAuthorityInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) that was returned when you called CreateCertificateAuthority.
	// This must be of the form:
	//
	// arn:aws:acm-pca:region:account:certificate-authority/12345678-1234-1234-1234-123456789012
	//
	// CertificateAuthorityArn is a required field
	CertificateAuthorityArn *string `min:"5" type:"string" required:"true"`

	// List of tags to be associated with the CA.
	//
	// Tags is a required field
	Tags []Tag `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s TagCertificateAuthorityInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TagCertificateAuthorityInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "TagCertificateAuthorityInput"}

	if s.CertificateAuthorityArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("CertificateAuthorityArn"))
	}
	if s.CertificateAuthorityArn != nil && len(*s.CertificateAuthorityArn) < 5 {
		invalidParams.Add(aws.NewErrParamMinLen("CertificateAuthorityArn", 5))
	}

	if s.Tags == nil {
		invalidParams.Add(aws.NewErrParamRequired("Tags"))
	}
	if s.Tags != nil && len(s.Tags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Tags", 1))
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

type TagCertificateAuthorityOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s TagCertificateAuthorityOutput) String() string {
	return awsutil.Prettify(s)
}

const opTagCertificateAuthority = "TagCertificateAuthority"

// TagCertificateAuthorityRequest returns a request value for making API operation for
// AWS Certificate Manager Private Certificate Authority.
//
// Adds one or more tags to your private CA. Tags are labels that you can use
// to identify and organize your AWS resources. Each tag consists of a key and
// an optional value. You specify the private CA on input by its Amazon Resource
// Name (ARN). You specify the tag by using a key-value pair. You can apply
// a tag to just one private CA if you want to identify a specific characteristic
// of that CA, or you can apply the same tag to multiple private CAs if you
// want to filter for a common relationship among those CAs. To remove one or
// more tags, use the UntagCertificateAuthority action. Call the ListTags action
// to see what tags are associated with your CA.
//
//    // Example sending a request using TagCertificateAuthorityRequest.
//    req := client.TagCertificateAuthorityRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/acm-pca-2017-08-22/TagCertificateAuthority
func (c *Client) TagCertificateAuthorityRequest(input *TagCertificateAuthorityInput) TagCertificateAuthorityRequest {
	op := &aws.Operation{
		Name:       opTagCertificateAuthority,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &TagCertificateAuthorityInput{}
	}

	req := c.newRequest(op, input, &TagCertificateAuthorityOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return TagCertificateAuthorityRequest{Request: req, Input: input, Copy: c.TagCertificateAuthorityRequest}
}

// TagCertificateAuthorityRequest is the request type for the
// TagCertificateAuthority API operation.
type TagCertificateAuthorityRequest struct {
	*aws.Request
	Input *TagCertificateAuthorityInput
	Copy  func(*TagCertificateAuthorityInput) TagCertificateAuthorityRequest
}

// Send marshals and sends the TagCertificateAuthority API request.
func (r TagCertificateAuthorityRequest) Send(ctx context.Context) (*TagCertificateAuthorityResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &TagCertificateAuthorityResponse{
		TagCertificateAuthorityOutput: r.Request.Data.(*TagCertificateAuthorityOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// TagCertificateAuthorityResponse is the response type for the
// TagCertificateAuthority API operation.
type TagCertificateAuthorityResponse struct {
	*TagCertificateAuthorityOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// TagCertificateAuthority request.
func (r *TagCertificateAuthorityResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
