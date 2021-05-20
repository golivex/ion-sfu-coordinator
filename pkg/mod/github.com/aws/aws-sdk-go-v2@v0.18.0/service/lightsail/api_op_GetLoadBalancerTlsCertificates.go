// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetLoadBalancerTlsCertificatesInput struct {
	_ struct{} `type:"structure"`

	// The name of the load balancer you associated with your SSL/TLS certificate.
	//
	// LoadBalancerName is a required field
	LoadBalancerName *string `locationName:"loadBalancerName" type:"string" required:"true"`
}

// String returns the string representation
func (s GetLoadBalancerTlsCertificatesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetLoadBalancerTlsCertificatesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetLoadBalancerTlsCertificatesInput"}

	if s.LoadBalancerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("LoadBalancerName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetLoadBalancerTlsCertificatesOutput struct {
	_ struct{} `type:"structure"`

	// An array of LoadBalancerTlsCertificate objects describing your SSL/TLS certificates.
	TlsCertificates []LoadBalancerTlsCertificate `locationName:"tlsCertificates" type:"list"`
}

// String returns the string representation
func (s GetLoadBalancerTlsCertificatesOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetLoadBalancerTlsCertificates = "GetLoadBalancerTlsCertificates"

// GetLoadBalancerTlsCertificatesRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Returns information about the TLS certificates that are associated with the
// specified Lightsail load balancer.
//
// TLS is just an updated, more secure version of Secure Socket Layer (SSL).
//
// You can have a maximum of 2 certificates associated with a Lightsail load
// balancer. One is active and the other is inactive.
//
//    // Example sending a request using GetLoadBalancerTlsCertificatesRequest.
//    req := client.GetLoadBalancerTlsCertificatesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/GetLoadBalancerTlsCertificates
func (c *Client) GetLoadBalancerTlsCertificatesRequest(input *GetLoadBalancerTlsCertificatesInput) GetLoadBalancerTlsCertificatesRequest {
	op := &aws.Operation{
		Name:       opGetLoadBalancerTlsCertificates,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetLoadBalancerTlsCertificatesInput{}
	}

	req := c.newRequest(op, input, &GetLoadBalancerTlsCertificatesOutput{})
	return GetLoadBalancerTlsCertificatesRequest{Request: req, Input: input, Copy: c.GetLoadBalancerTlsCertificatesRequest}
}

// GetLoadBalancerTlsCertificatesRequest is the request type for the
// GetLoadBalancerTlsCertificates API operation.
type GetLoadBalancerTlsCertificatesRequest struct {
	*aws.Request
	Input *GetLoadBalancerTlsCertificatesInput
	Copy  func(*GetLoadBalancerTlsCertificatesInput) GetLoadBalancerTlsCertificatesRequest
}

// Send marshals and sends the GetLoadBalancerTlsCertificates API request.
func (r GetLoadBalancerTlsCertificatesRequest) Send(ctx context.Context) (*GetLoadBalancerTlsCertificatesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetLoadBalancerTlsCertificatesResponse{
		GetLoadBalancerTlsCertificatesOutput: r.Request.Data.(*GetLoadBalancerTlsCertificatesOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetLoadBalancerTlsCertificatesResponse is the response type for the
// GetLoadBalancerTlsCertificates API operation.
type GetLoadBalancerTlsCertificatesResponse struct {
	*GetLoadBalancerTlsCertificatesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetLoadBalancerTlsCertificates request.
func (r *GetLoadBalancerTlsCertificatesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
