// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package guardduty

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DeleteInvitationsInput struct {
	_ struct{} `type:"structure"`

	// A list of account IDs of the AWS accounts that sent invitations to the current
	// member account that you want to delete invitations from.
	//
	// AccountIds is a required field
	AccountIds []string `locationName:"accountIds" min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s DeleteInvitationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteInvitationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteInvitationsInput"}

	if s.AccountIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountIds"))
	}
	if s.AccountIds != nil && len(s.AccountIds) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AccountIds", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteInvitationsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AccountIds != nil {
		v := s.AccountIds

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "accountIds", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

type DeleteInvitationsOutput struct {
	_ struct{} `type:"structure"`

	// A list of objects containing the unprocessed account and a result string
	// explaining why it was unprocessed.
	//
	// UnprocessedAccounts is a required field
	UnprocessedAccounts []UnprocessedAccount `locationName:"unprocessedAccounts" type:"list" required:"true"`
}

// String returns the string representation
func (s DeleteInvitationsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteInvitationsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.UnprocessedAccounts != nil {
		v := s.UnprocessedAccounts

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "unprocessedAccounts", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opDeleteInvitations = "DeleteInvitations"

// DeleteInvitationsRequest returns a request value for making API operation for
// Amazon GuardDuty.
//
// Deletes invitations sent to the current member account by AWS accounts specified
// by their account IDs.
//
//    // Example sending a request using DeleteInvitationsRequest.
//    req := client.DeleteInvitationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/DeleteInvitations
func (c *Client) DeleteInvitationsRequest(input *DeleteInvitationsInput) DeleteInvitationsRequest {
	op := &aws.Operation{
		Name:       opDeleteInvitations,
		HTTPMethod: "POST",
		HTTPPath:   "/invitation/delete",
	}

	if input == nil {
		input = &DeleteInvitationsInput{}
	}

	req := c.newRequest(op, input, &DeleteInvitationsOutput{})
	return DeleteInvitationsRequest{Request: req, Input: input, Copy: c.DeleteInvitationsRequest}
}

// DeleteInvitationsRequest is the request type for the
// DeleteInvitations API operation.
type DeleteInvitationsRequest struct {
	*aws.Request
	Input *DeleteInvitationsInput
	Copy  func(*DeleteInvitationsInput) DeleteInvitationsRequest
}

// Send marshals and sends the DeleteInvitations API request.
func (r DeleteInvitationsRequest) Send(ctx context.Context) (*DeleteInvitationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteInvitationsResponse{
		DeleteInvitationsOutput: r.Request.Data.(*DeleteInvitationsOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteInvitationsResponse is the response type for the
// DeleteInvitations API operation.
type DeleteInvitationsResponse struct {
	*DeleteInvitationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteInvitations request.
func (r *DeleteInvitationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
