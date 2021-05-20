// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package backup

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetBackupSelectionInput struct {
	_ struct{} `type:"structure"`

	// Uniquely identifies a backup plan.
	//
	// BackupPlanId is a required field
	BackupPlanId *string `location:"uri" locationName:"backupPlanId" type:"string" required:"true"`

	// Uniquely identifies the body of a request to assign a set of resources to
	// a backup plan.
	//
	// SelectionId is a required field
	SelectionId *string `location:"uri" locationName:"selectionId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetBackupSelectionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetBackupSelectionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetBackupSelectionInput"}

	if s.BackupPlanId == nil {
		invalidParams.Add(aws.NewErrParamRequired("BackupPlanId"))
	}

	if s.SelectionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SelectionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBackupSelectionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.BackupPlanId != nil {
		v := *s.BackupPlanId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "backupPlanId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SelectionId != nil {
		v := *s.SelectionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "selectionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetBackupSelectionOutput struct {
	_ struct{} `type:"structure"`

	// Uniquely identifies a backup plan.
	BackupPlanId *string `type:"string"`

	// Specifies the body of a request to assign a set of resources to a backup
	// plan.
	//
	// It includes an array of resources, an optional array of patterns to exclude
	// resources, an optional role to provide access to the AWS service that the
	// resource belongs to, and an optional array of tags used to identify a set
	// of resources.
	BackupSelection *BackupSelection `type:"structure"`

	// The date and time a backup selection is created, in Unix format and Coordinated
	// Universal Time (UTC). The value of CreationDate is accurate to milliseconds.
	// For example, the value 1516925490.087 represents Friday, January 26, 2018
	// 12:11:30.087 AM.
	CreationDate *time.Time `type:"timestamp"`

	// A unique string that identifies the request and allows failed requests to
	// be retried without the risk of executing the operation twice.
	CreatorRequestId *string `type:"string"`

	// Uniquely identifies the body of a request to assign a set of resources to
	// a backup plan.
	SelectionId *string `type:"string"`
}

// String returns the string representation
func (s GetBackupSelectionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBackupSelectionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.BackupPlanId != nil {
		v := *s.BackupPlanId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "BackupPlanId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.BackupSelection != nil {
		v := s.BackupSelection

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "BackupSelection", v, metadata)
	}
	if s.CreationDate != nil {
		v := *s.CreationDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CreationDate",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.CreatorRequestId != nil {
		v := *s.CreatorRequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CreatorRequestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SelectionId != nil {
		v := *s.SelectionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SelectionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetBackupSelection = "GetBackupSelection"

// GetBackupSelectionRequest returns a request value for making API operation for
// AWS Backup.
//
// Returns selection metadata and a document in JSON format that specifies a
// list of resources that are associated with a backup plan.
//
//    // Example sending a request using GetBackupSelectionRequest.
//    req := client.GetBackupSelectionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/backup-2018-11-15/GetBackupSelection
func (c *Client) GetBackupSelectionRequest(input *GetBackupSelectionInput) GetBackupSelectionRequest {
	op := &aws.Operation{
		Name:       opGetBackupSelection,
		HTTPMethod: "GET",
		HTTPPath:   "/backup/plans/{backupPlanId}/selections/{selectionId}",
	}

	if input == nil {
		input = &GetBackupSelectionInput{}
	}

	req := c.newRequest(op, input, &GetBackupSelectionOutput{})
	return GetBackupSelectionRequest{Request: req, Input: input, Copy: c.GetBackupSelectionRequest}
}

// GetBackupSelectionRequest is the request type for the
// GetBackupSelection API operation.
type GetBackupSelectionRequest struct {
	*aws.Request
	Input *GetBackupSelectionInput
	Copy  func(*GetBackupSelectionInput) GetBackupSelectionRequest
}

// Send marshals and sends the GetBackupSelection API request.
func (r GetBackupSelectionRequest) Send(ctx context.Context) (*GetBackupSelectionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetBackupSelectionResponse{
		GetBackupSelectionOutput: r.Request.Data.(*GetBackupSelectionOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetBackupSelectionResponse is the response type for the
// GetBackupSelection API operation.
type GetBackupSelectionResponse struct {
	*GetBackupSelectionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetBackupSelection request.
func (r *GetBackupSelectionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
