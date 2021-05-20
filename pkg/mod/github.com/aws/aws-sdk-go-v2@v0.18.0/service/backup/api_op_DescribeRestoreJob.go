// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package backup

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DescribeRestoreJobInput struct {
	_ struct{} `type:"structure"`

	// Uniquely identifies the job that restores a recovery point.
	//
	// RestoreJobId is a required field
	RestoreJobId *string `location:"uri" locationName:"restoreJobId" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeRestoreJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeRestoreJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeRestoreJobInput"}

	if s.RestoreJobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RestoreJobId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeRestoreJobInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.RestoreJobId != nil {
		v := *s.RestoreJobId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "restoreJobId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DescribeRestoreJobOutput struct {
	_ struct{} `type:"structure"`

	// The size, in bytes, of the restored resource.
	BackupSizeInBytes *int64 `type:"long"`

	// The date and time that a job to restore a recovery point is completed, in
	// Unix format and Coordinated Universal Time (UTC). The value of CompletionDate
	// is accurate to milliseconds. For example, the value 1516925490.087 represents
	// Friday, January 26, 2018 12:11:30.087 AM.
	CompletionDate *time.Time `type:"timestamp"`

	// An Amazon Resource Name (ARN) that uniquely identifies a resource whose recovery
	// point is being restored. The format of the ARN depends on the resource type
	// of the backed-up resource.
	CreatedResourceArn *string `type:"string"`

	// The date and time that a restore job is created, in Unix format and Coordinated
	// Universal Time (UTC). The value of CreationDate is accurate to milliseconds.
	// For example, the value 1516925490.087 represents Friday, January 26, 2018
	// 12:11:30.087 AM.
	CreationDate *time.Time `type:"timestamp"`

	// The amount of time in minutes that a job restoring a recovery point is expected
	// to take.
	ExpectedCompletionTimeMinutes *int64 `type:"long"`

	// Specifies the IAM role ARN used to create the target recovery point; for
	// example, arn:aws:iam::123456789012:role/S3Access.
	IamRoleArn *string `type:"string"`

	// Contains an estimated percentage that is complete of a job at the time the
	// job status was queried.
	PercentDone *string `type:"string"`

	// An ARN that uniquely identifies a recovery point; for example, arn:aws:backup:us-east-1:123456789012:recovery-point:1EB3B5E7-9EB0-435A-A80B-108B488B0D45.
	RecoveryPointArn *string `type:"string"`

	// Uniquely identifies the job that restores a recovery point.
	RestoreJobId *string `type:"string"`

	// Status code specifying the state of the job that is initiated by AWS Backup
	// to restore a recovery point.
	Status RestoreJobStatus `type:"string" enum:"true"`

	// A detailed message explaining the status of a job to restore a recovery point.
	StatusMessage *string `type:"string"`
}

// String returns the string representation
func (s DescribeRestoreJobOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeRestoreJobOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.BackupSizeInBytes != nil {
		v := *s.BackupSizeInBytes

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "BackupSizeInBytes", protocol.Int64Value(v), metadata)
	}
	if s.CompletionDate != nil {
		v := *s.CompletionDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CompletionDate",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.CreatedResourceArn != nil {
		v := *s.CreatedResourceArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CreatedResourceArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CreationDate != nil {
		v := *s.CreationDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CreationDate",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.ExpectedCompletionTimeMinutes != nil {
		v := *s.ExpectedCompletionTimeMinutes

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ExpectedCompletionTimeMinutes", protocol.Int64Value(v), metadata)
	}
	if s.IamRoleArn != nil {
		v := *s.IamRoleArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "IamRoleArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PercentDone != nil {
		v := *s.PercentDone

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "PercentDone", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RecoveryPointArn != nil {
		v := *s.RecoveryPointArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RecoveryPointArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RestoreJobId != nil {
		v := *s.RestoreJobId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RestoreJobId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.StatusMessage != nil {
		v := *s.StatusMessage

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "StatusMessage", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opDescribeRestoreJob = "DescribeRestoreJob"

// DescribeRestoreJobRequest returns a request value for making API operation for
// AWS Backup.
//
// Returns metadata associated with a restore job that is specified by a job
// ID.
//
//    // Example sending a request using DescribeRestoreJobRequest.
//    req := client.DescribeRestoreJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/backup-2018-11-15/DescribeRestoreJob
func (c *Client) DescribeRestoreJobRequest(input *DescribeRestoreJobInput) DescribeRestoreJobRequest {
	op := &aws.Operation{
		Name:       opDescribeRestoreJob,
		HTTPMethod: "GET",
		HTTPPath:   "/restore-jobs/{restoreJobId}",
	}

	if input == nil {
		input = &DescribeRestoreJobInput{}
	}

	req := c.newRequest(op, input, &DescribeRestoreJobOutput{})
	return DescribeRestoreJobRequest{Request: req, Input: input, Copy: c.DescribeRestoreJobRequest}
}

// DescribeRestoreJobRequest is the request type for the
// DescribeRestoreJob API operation.
type DescribeRestoreJobRequest struct {
	*aws.Request
	Input *DescribeRestoreJobInput
	Copy  func(*DescribeRestoreJobInput) DescribeRestoreJobRequest
}

// Send marshals and sends the DescribeRestoreJob API request.
func (r DescribeRestoreJobRequest) Send(ctx context.Context) (*DescribeRestoreJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeRestoreJobResponse{
		DescribeRestoreJobOutput: r.Request.Data.(*DescribeRestoreJobOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeRestoreJobResponse is the response type for the
// DescribeRestoreJob API operation.
type DescribeRestoreJobResponse struct {
	*DescribeRestoreJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeRestoreJob request.
func (r *DescribeRestoreJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
