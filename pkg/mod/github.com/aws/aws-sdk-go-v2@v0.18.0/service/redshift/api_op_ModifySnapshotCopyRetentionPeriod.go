// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ModifySnapshotCopyRetentionPeriodInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier of the cluster for which you want to change the retention
	// period for either automated or manual snapshots that are copied to a destination
	// AWS Region.
	//
	// Constraints: Must be the valid name of an existing cluster that has cross-region
	// snapshot copy enabled.
	//
	// ClusterIdentifier is a required field
	ClusterIdentifier *string `type:"string" required:"true"`

	// Indicates whether to apply the snapshot retention period to newly copied
	// manual snapshots instead of automated snapshots.
	Manual *bool `type:"boolean"`

	// The number of days to retain automated snapshots in the destination AWS Region
	// after they are copied from the source AWS Region.
	//
	// By default, this only changes the retention period of copied automated snapshots.
	//
	// If you decrease the retention period for automated snapshots that are copied
	// to a destination AWS Region, Amazon Redshift deletes any existing automated
	// snapshots that were copied to the destination AWS Region and that fall outside
	// of the new retention period.
	//
	// Constraints: Must be at least 1 and no more than 35 for automated snapshots.
	//
	// If you specify the manual option, only newly copied manual snapshots will
	// have the new retention period.
	//
	// If you specify the value of -1 newly copied manual snapshots are retained
	// indefinitely.
	//
	// Constraints: The number of days must be either -1 or an integer between 1
	// and 3,653 for manual snapshots.
	//
	// RetentionPeriod is a required field
	RetentionPeriod *int64 `type:"integer" required:"true"`
}

// String returns the string representation
func (s ModifySnapshotCopyRetentionPeriodInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifySnapshotCopyRetentionPeriodInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifySnapshotCopyRetentionPeriodInput"}

	if s.ClusterIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClusterIdentifier"))
	}

	if s.RetentionPeriod == nil {
		invalidParams.Add(aws.NewErrParamRequired("RetentionPeriod"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ModifySnapshotCopyRetentionPeriodOutput struct {
	_ struct{} `type:"structure"`

	// Describes a cluster.
	Cluster *Cluster `type:"structure"`
}

// String returns the string representation
func (s ModifySnapshotCopyRetentionPeriodOutput) String() string {
	return awsutil.Prettify(s)
}

const opModifySnapshotCopyRetentionPeriod = "ModifySnapshotCopyRetentionPeriod"

// ModifySnapshotCopyRetentionPeriodRequest returns a request value for making API operation for
// Amazon Redshift.
//
// Modifies the number of days to retain snapshots in the destination AWS Region
// after they are copied from the source AWS Region. By default, this operation
// only changes the retention period of copied automated snapshots. The retention
// periods for both new and existing copied automated snapshots are updated
// with the new retention period. You can set the manual option to change only
// the retention periods of copied manual snapshots. If you set this option,
// only newly copied manual snapshots have the new retention period.
//
//    // Example sending a request using ModifySnapshotCopyRetentionPeriodRequest.
//    req := client.ModifySnapshotCopyRetentionPeriodRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/ModifySnapshotCopyRetentionPeriod
func (c *Client) ModifySnapshotCopyRetentionPeriodRequest(input *ModifySnapshotCopyRetentionPeriodInput) ModifySnapshotCopyRetentionPeriodRequest {
	op := &aws.Operation{
		Name:       opModifySnapshotCopyRetentionPeriod,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifySnapshotCopyRetentionPeriodInput{}
	}

	req := c.newRequest(op, input, &ModifySnapshotCopyRetentionPeriodOutput{})
	return ModifySnapshotCopyRetentionPeriodRequest{Request: req, Input: input, Copy: c.ModifySnapshotCopyRetentionPeriodRequest}
}

// ModifySnapshotCopyRetentionPeriodRequest is the request type for the
// ModifySnapshotCopyRetentionPeriod API operation.
type ModifySnapshotCopyRetentionPeriodRequest struct {
	*aws.Request
	Input *ModifySnapshotCopyRetentionPeriodInput
	Copy  func(*ModifySnapshotCopyRetentionPeriodInput) ModifySnapshotCopyRetentionPeriodRequest
}

// Send marshals and sends the ModifySnapshotCopyRetentionPeriod API request.
func (r ModifySnapshotCopyRetentionPeriodRequest) Send(ctx context.Context) (*ModifySnapshotCopyRetentionPeriodResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifySnapshotCopyRetentionPeriodResponse{
		ModifySnapshotCopyRetentionPeriodOutput: r.Request.Data.(*ModifySnapshotCopyRetentionPeriodOutput),
		response:                                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifySnapshotCopyRetentionPeriodResponse is the response type for the
// ModifySnapshotCopyRetentionPeriod API operation.
type ModifySnapshotCopyRetentionPeriodResponse struct {
	*ModifySnapshotCopyRetentionPeriodOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifySnapshotCopyRetentionPeriod request.
func (r *ModifySnapshotCopyRetentionPeriodResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
