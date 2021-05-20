// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package gamelift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input for a request action.
type DescribeScalingPoliciesInput struct {
	_ struct{} `type:"structure"`

	// Unique identifier for a fleet to retrieve scaling policies for.
	//
	// FleetId is a required field
	FleetId *string `type:"string" required:"true"`

	// Maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages.
	Limit *int64 `min:"1" type:"integer"`

	// Token that indicates the start of the next sequential page of results. Use
	// the token that is returned with a previous call to this action. To start
	// at the beginning of the result set, do not specify a value.
	NextToken *string `min:"1" type:"string"`

	// Scaling policy status to filter results on. A scaling policy is only in force
	// when in an ACTIVE status.
	//
	//    * ACTIVE -- The scaling policy is currently in force.
	//
	//    * UPDATEREQUESTED -- A request to update the scaling policy has been received.
	//
	//    * UPDATING -- A change is being made to the scaling policy.
	//
	//    * DELETEREQUESTED -- A request to delete the scaling policy has been received.
	//
	//    * DELETING -- The scaling policy is being deleted.
	//
	//    * DELETED -- The scaling policy has been deleted.
	//
	//    * ERROR -- An error occurred in creating the policy. It should be removed
	//    and recreated.
	StatusFilter ScalingStatusType `type:"string" enum:"true"`
}

// String returns the string representation
func (s DescribeScalingPoliciesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeScalingPoliciesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeScalingPoliciesInput"}

	if s.FleetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("FleetId"))
	}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the returned data in response to a request action.
type DescribeScalingPoliciesOutput struct {
	_ struct{} `type:"structure"`

	// Token that indicates where to resume retrieving results on the next call
	// to this action. If no token is returned, these results represent the end
	// of the list.
	NextToken *string `min:"1" type:"string"`

	// Collection of objects containing the scaling policies matching the request.
	ScalingPolicies []ScalingPolicy `type:"list"`
}

// String returns the string representation
func (s DescribeScalingPoliciesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeScalingPolicies = "DescribeScalingPolicies"

// DescribeScalingPoliciesRequest returns a request value for making API operation for
// Amazon GameLift.
//
// Retrieves all scaling policies applied to a fleet.
//
// To get a fleet's scaling policies, specify the fleet ID. You can filter this
// request by policy status, such as to retrieve only active scaling policies.
// Use the pagination parameters to retrieve results as a set of sequential
// pages. If successful, set of ScalingPolicy objects is returned for the fleet.
//
// A fleet may have all of its scaling policies suspended (StopFleetActions).
// This action does not affect the status of the scaling policies, which remains
// ACTIVE. To see whether a fleet's scaling policies are in force or suspended,
// call DescribeFleetAttributes and check the stopped actions.
//
//    * DescribeFleetCapacity
//
//    * UpdateFleetCapacity
//
//    * DescribeEC2InstanceLimits
//
//    * Manage scaling policies: PutScalingPolicy (auto-scaling) DescribeScalingPolicies
//    (auto-scaling) DeleteScalingPolicy (auto-scaling)
//
//    * Manage fleet actions: StartFleetActions StopFleetActions
//
//    // Example sending a request using DescribeScalingPoliciesRequest.
//    req := client.DescribeScalingPoliciesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/DescribeScalingPolicies
func (c *Client) DescribeScalingPoliciesRequest(input *DescribeScalingPoliciesInput) DescribeScalingPoliciesRequest {
	op := &aws.Operation{
		Name:       opDescribeScalingPolicies,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeScalingPoliciesInput{}
	}

	req := c.newRequest(op, input, &DescribeScalingPoliciesOutput{})
	return DescribeScalingPoliciesRequest{Request: req, Input: input, Copy: c.DescribeScalingPoliciesRequest}
}

// DescribeScalingPoliciesRequest is the request type for the
// DescribeScalingPolicies API operation.
type DescribeScalingPoliciesRequest struct {
	*aws.Request
	Input *DescribeScalingPoliciesInput
	Copy  func(*DescribeScalingPoliciesInput) DescribeScalingPoliciesRequest
}

// Send marshals and sends the DescribeScalingPolicies API request.
func (r DescribeScalingPoliciesRequest) Send(ctx context.Context) (*DescribeScalingPoliciesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeScalingPoliciesResponse{
		DescribeScalingPoliciesOutput: r.Request.Data.(*DescribeScalingPoliciesOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeScalingPoliciesResponse is the response type for the
// DescribeScalingPolicies API operation.
type DescribeScalingPoliciesResponse struct {
	*DescribeScalingPoliciesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeScalingPolicies request.
func (r *DescribeScalingPoliciesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
