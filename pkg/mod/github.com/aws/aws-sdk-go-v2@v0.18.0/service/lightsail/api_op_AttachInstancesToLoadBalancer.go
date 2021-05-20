// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AttachInstancesToLoadBalancerInput struct {
	_ struct{} `type:"structure"`

	// An array of strings representing the instance name(s) you want to attach
	// to your load balancer.
	//
	// An instance must be running before you can attach it to your load balancer.
	//
	// There are no additional limits on the number of instances you can attach
	// to your load balancer, aside from the limit of Lightsail instances you can
	// create in your account (20).
	//
	// InstanceNames is a required field
	InstanceNames []string `locationName:"instanceNames" type:"list" required:"true"`

	// The name of the load balancer.
	//
	// LoadBalancerName is a required field
	LoadBalancerName *string `locationName:"loadBalancerName" type:"string" required:"true"`
}

// String returns the string representation
func (s AttachInstancesToLoadBalancerInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AttachInstancesToLoadBalancerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AttachInstancesToLoadBalancerInput"}

	if s.InstanceNames == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceNames"))
	}

	if s.LoadBalancerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("LoadBalancerName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AttachInstancesToLoadBalancerOutput struct {
	_ struct{} `type:"structure"`

	// An object representing the API operations.
	Operations []Operation `locationName:"operations" type:"list"`
}

// String returns the string representation
func (s AttachInstancesToLoadBalancerOutput) String() string {
	return awsutil.Prettify(s)
}

const opAttachInstancesToLoadBalancer = "AttachInstancesToLoadBalancer"

// AttachInstancesToLoadBalancerRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Attaches one or more Lightsail instances to a load balancer.
//
// After some time, the instances are attached to the load balancer and the
// health check status is available.
//
// The attach instances to load balancer operation supports tag-based access
// control via resource tags applied to the resource identified by load balancer
// name. For more information, see the Lightsail Dev Guide (https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-controlling-access-using-tags).
//
//    // Example sending a request using AttachInstancesToLoadBalancerRequest.
//    req := client.AttachInstancesToLoadBalancerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/AttachInstancesToLoadBalancer
func (c *Client) AttachInstancesToLoadBalancerRequest(input *AttachInstancesToLoadBalancerInput) AttachInstancesToLoadBalancerRequest {
	op := &aws.Operation{
		Name:       opAttachInstancesToLoadBalancer,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AttachInstancesToLoadBalancerInput{}
	}

	req := c.newRequest(op, input, &AttachInstancesToLoadBalancerOutput{})
	return AttachInstancesToLoadBalancerRequest{Request: req, Input: input, Copy: c.AttachInstancesToLoadBalancerRequest}
}

// AttachInstancesToLoadBalancerRequest is the request type for the
// AttachInstancesToLoadBalancer API operation.
type AttachInstancesToLoadBalancerRequest struct {
	*aws.Request
	Input *AttachInstancesToLoadBalancerInput
	Copy  func(*AttachInstancesToLoadBalancerInput) AttachInstancesToLoadBalancerRequest
}

// Send marshals and sends the AttachInstancesToLoadBalancer API request.
func (r AttachInstancesToLoadBalancerRequest) Send(ctx context.Context) (*AttachInstancesToLoadBalancerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AttachInstancesToLoadBalancerResponse{
		AttachInstancesToLoadBalancerOutput: r.Request.Data.(*AttachInstancesToLoadBalancerOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AttachInstancesToLoadBalancerResponse is the response type for the
// AttachInstancesToLoadBalancer API operation.
type AttachInstancesToLoadBalancerResponse struct {
	*AttachInstancesToLoadBalancerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AttachInstancesToLoadBalancer request.
func (r *AttachInstancesToLoadBalancerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
