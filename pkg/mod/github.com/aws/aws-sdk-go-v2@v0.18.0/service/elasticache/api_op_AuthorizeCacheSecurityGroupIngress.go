// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of an AuthorizeCacheSecurityGroupIngress operation.
type AuthorizeCacheSecurityGroupIngressInput struct {
	_ struct{} `type:"structure"`

	// The cache security group that allows network ingress.
	//
	// CacheSecurityGroupName is a required field
	CacheSecurityGroupName *string `type:"string" required:"true"`

	// The Amazon EC2 security group to be authorized for ingress to the cache security
	// group.
	//
	// EC2SecurityGroupName is a required field
	EC2SecurityGroupName *string `type:"string" required:"true"`

	// The AWS account number of the Amazon EC2 security group owner. Note that
	// this is not the same thing as an AWS access key ID - you must provide a valid
	// AWS account number for this parameter.
	//
	// EC2SecurityGroupOwnerId is a required field
	EC2SecurityGroupOwnerId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s AuthorizeCacheSecurityGroupIngressInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AuthorizeCacheSecurityGroupIngressInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AuthorizeCacheSecurityGroupIngressInput"}

	if s.CacheSecurityGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("CacheSecurityGroupName"))
	}

	if s.EC2SecurityGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("EC2SecurityGroupName"))
	}

	if s.EC2SecurityGroupOwnerId == nil {
		invalidParams.Add(aws.NewErrParamRequired("EC2SecurityGroupOwnerId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AuthorizeCacheSecurityGroupIngressOutput struct {
	_ struct{} `type:"structure"`

	// Represents the output of one of the following operations:
	//
	//    * AuthorizeCacheSecurityGroupIngress
	//
	//    * CreateCacheSecurityGroup
	//
	//    * RevokeCacheSecurityGroupIngress
	CacheSecurityGroup *CacheSecurityGroup `type:"structure"`
}

// String returns the string representation
func (s AuthorizeCacheSecurityGroupIngressOutput) String() string {
	return awsutil.Prettify(s)
}

const opAuthorizeCacheSecurityGroupIngress = "AuthorizeCacheSecurityGroupIngress"

// AuthorizeCacheSecurityGroupIngressRequest returns a request value for making API operation for
// Amazon ElastiCache.
//
// Allows network ingress to a cache security group. Applications using ElastiCache
// must be running on Amazon EC2, and Amazon EC2 security groups are used as
// the authorization mechanism.
//
// You cannot authorize ingress from an Amazon EC2 security group in one region
// to an ElastiCache cluster in another region.
//
//    // Example sending a request using AuthorizeCacheSecurityGroupIngressRequest.
//    req := client.AuthorizeCacheSecurityGroupIngressRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/AuthorizeCacheSecurityGroupIngress
func (c *Client) AuthorizeCacheSecurityGroupIngressRequest(input *AuthorizeCacheSecurityGroupIngressInput) AuthorizeCacheSecurityGroupIngressRequest {
	op := &aws.Operation{
		Name:       opAuthorizeCacheSecurityGroupIngress,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AuthorizeCacheSecurityGroupIngressInput{}
	}

	req := c.newRequest(op, input, &AuthorizeCacheSecurityGroupIngressOutput{})
	return AuthorizeCacheSecurityGroupIngressRequest{Request: req, Input: input, Copy: c.AuthorizeCacheSecurityGroupIngressRequest}
}

// AuthorizeCacheSecurityGroupIngressRequest is the request type for the
// AuthorizeCacheSecurityGroupIngress API operation.
type AuthorizeCacheSecurityGroupIngressRequest struct {
	*aws.Request
	Input *AuthorizeCacheSecurityGroupIngressInput
	Copy  func(*AuthorizeCacheSecurityGroupIngressInput) AuthorizeCacheSecurityGroupIngressRequest
}

// Send marshals and sends the AuthorizeCacheSecurityGroupIngress API request.
func (r AuthorizeCacheSecurityGroupIngressRequest) Send(ctx context.Context) (*AuthorizeCacheSecurityGroupIngressResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AuthorizeCacheSecurityGroupIngressResponse{
		AuthorizeCacheSecurityGroupIngressOutput: r.Request.Data.(*AuthorizeCacheSecurityGroupIngressOutput),
		response:                                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AuthorizeCacheSecurityGroupIngressResponse is the response type for the
// AuthorizeCacheSecurityGroupIngress API operation.
type AuthorizeCacheSecurityGroupIngressResponse struct {
	*AuthorizeCacheSecurityGroupIngressOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AuthorizeCacheSecurityGroupIngress request.
func (r *AuthorizeCacheSecurityGroupIngressResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
