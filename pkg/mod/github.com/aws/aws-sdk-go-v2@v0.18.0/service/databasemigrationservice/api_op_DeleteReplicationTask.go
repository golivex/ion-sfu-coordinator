// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package databasemigrationservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteReplicationTaskInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the replication task to be deleted.
	//
	// ReplicationTaskArn is a required field
	ReplicationTaskArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteReplicationTaskInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteReplicationTaskInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteReplicationTaskInput"}

	if s.ReplicationTaskArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ReplicationTaskArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteReplicationTaskOutput struct {
	_ struct{} `type:"structure"`

	// The deleted replication task.
	ReplicationTask *ReplicationTask `type:"structure"`
}

// String returns the string representation
func (s DeleteReplicationTaskOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteReplicationTask = "DeleteReplicationTask"

// DeleteReplicationTaskRequest returns a request value for making API operation for
// AWS Database Migration Service.
//
// Deletes the specified replication task.
//
//    // Example sending a request using DeleteReplicationTaskRequest.
//    req := client.DeleteReplicationTaskRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dms-2016-01-01/DeleteReplicationTask
func (c *Client) DeleteReplicationTaskRequest(input *DeleteReplicationTaskInput) DeleteReplicationTaskRequest {
	op := &aws.Operation{
		Name:       opDeleteReplicationTask,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteReplicationTaskInput{}
	}

	req := c.newRequest(op, input, &DeleteReplicationTaskOutput{})
	return DeleteReplicationTaskRequest{Request: req, Input: input, Copy: c.DeleteReplicationTaskRequest}
}

// DeleteReplicationTaskRequest is the request type for the
// DeleteReplicationTask API operation.
type DeleteReplicationTaskRequest struct {
	*aws.Request
	Input *DeleteReplicationTaskInput
	Copy  func(*DeleteReplicationTaskInput) DeleteReplicationTaskRequest
}

// Send marshals and sends the DeleteReplicationTask API request.
func (r DeleteReplicationTaskRequest) Send(ctx context.Context) (*DeleteReplicationTaskResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteReplicationTaskResponse{
		DeleteReplicationTaskOutput: r.Request.Data.(*DeleteReplicationTaskOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteReplicationTaskResponse is the response type for the
// DeleteReplicationTask API operation.
type DeleteReplicationTaskResponse struct {
	*DeleteReplicationTaskOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteReplicationTask request.
func (r *DeleteReplicationTaskResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
