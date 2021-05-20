// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package redshift

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ModifyClusterMaintenanceInput struct {
	_ struct{} `type:"structure"`

	// A unique identifier for the cluster.
	//
	// ClusterIdentifier is a required field
	ClusterIdentifier *string `type:"string" required:"true"`

	// A boolean indicating whether to enable the deferred maintenance window.
	DeferMaintenance *bool `type:"boolean"`

	// An integer indicating the duration of the maintenance window in days. If
	// you specify a duration, you can't specify an end time. The duration must
	// be 45 days or less.
	DeferMaintenanceDuration *int64 `type:"integer"`

	// A timestamp indicating end time for the deferred maintenance window. If you
	// specify an end time, you can't specify a duration.
	DeferMaintenanceEndTime *time.Time `type:"timestamp"`

	// A unique identifier for the deferred maintenance window.
	DeferMaintenanceIdentifier *string `type:"string"`

	// A timestamp indicating the start time for the deferred maintenance window.
	DeferMaintenanceStartTime *time.Time `type:"timestamp"`
}

// String returns the string representation
func (s ModifyClusterMaintenanceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyClusterMaintenanceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyClusterMaintenanceInput"}

	if s.ClusterIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClusterIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ModifyClusterMaintenanceOutput struct {
	_ struct{} `type:"structure"`

	// Describes a cluster.
	Cluster *Cluster `type:"structure"`
}

// String returns the string representation
func (s ModifyClusterMaintenanceOutput) String() string {
	return awsutil.Prettify(s)
}

const opModifyClusterMaintenance = "ModifyClusterMaintenance"

// ModifyClusterMaintenanceRequest returns a request value for making API operation for
// Amazon Redshift.
//
// Modifies the maintenance settings of a cluster. For example, you can defer
// a maintenance window. You can also update or cancel a deferment.
//
//    // Example sending a request using ModifyClusterMaintenanceRequest.
//    req := client.ModifyClusterMaintenanceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/ModifyClusterMaintenance
func (c *Client) ModifyClusterMaintenanceRequest(input *ModifyClusterMaintenanceInput) ModifyClusterMaintenanceRequest {
	op := &aws.Operation{
		Name:       opModifyClusterMaintenance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyClusterMaintenanceInput{}
	}

	req := c.newRequest(op, input, &ModifyClusterMaintenanceOutput{})
	return ModifyClusterMaintenanceRequest{Request: req, Input: input, Copy: c.ModifyClusterMaintenanceRequest}
}

// ModifyClusterMaintenanceRequest is the request type for the
// ModifyClusterMaintenance API operation.
type ModifyClusterMaintenanceRequest struct {
	*aws.Request
	Input *ModifyClusterMaintenanceInput
	Copy  func(*ModifyClusterMaintenanceInput) ModifyClusterMaintenanceRequest
}

// Send marshals and sends the ModifyClusterMaintenance API request.
func (r ModifyClusterMaintenanceRequest) Send(ctx context.Context) (*ModifyClusterMaintenanceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyClusterMaintenanceResponse{
		ModifyClusterMaintenanceOutput: r.Request.Data.(*ModifyClusterMaintenanceOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyClusterMaintenanceResponse is the response type for the
// ModifyClusterMaintenance API operation.
type ModifyClusterMaintenanceResponse struct {
	*ModifyClusterMaintenanceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyClusterMaintenance request.
func (r *ModifyClusterMaintenanceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
