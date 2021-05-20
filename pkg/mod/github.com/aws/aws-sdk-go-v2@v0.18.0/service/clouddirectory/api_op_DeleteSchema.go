// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package clouddirectory

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DeleteSchemaInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the development schema. For more information,
	// see arns.
	//
	// SchemaArn is a required field
	SchemaArn *string `location:"header" locationName:"x-amz-data-partition" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteSchemaInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteSchemaInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteSchemaInput"}

	if s.SchemaArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("SchemaArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteSchemaInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.SchemaArn != nil {
		v := *s.SchemaArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-data-partition", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteSchemaOutput struct {
	_ struct{} `type:"structure"`

	// The input ARN that is returned as part of the response. For more information,
	// see arns.
	SchemaArn *string `type:"string"`
}

// String returns the string representation
func (s DeleteSchemaOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteSchemaOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.SchemaArn != nil {
		v := *s.SchemaArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SchemaArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opDeleteSchema = "DeleteSchema"

// DeleteSchemaRequest returns a request value for making API operation for
// Amazon CloudDirectory.
//
// Deletes a given schema. Schemas in a development and published state can
// only be deleted.
//
//    // Example sending a request using DeleteSchemaRequest.
//    req := client.DeleteSchemaRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/DeleteSchema
func (c *Client) DeleteSchemaRequest(input *DeleteSchemaInput) DeleteSchemaRequest {
	op := &aws.Operation{
		Name:       opDeleteSchema,
		HTTPMethod: "PUT",
		HTTPPath:   "/amazonclouddirectory/2017-01-11/schema",
	}

	if input == nil {
		input = &DeleteSchemaInput{}
	}

	req := c.newRequest(op, input, &DeleteSchemaOutput{})
	return DeleteSchemaRequest{Request: req, Input: input, Copy: c.DeleteSchemaRequest}
}

// DeleteSchemaRequest is the request type for the
// DeleteSchema API operation.
type DeleteSchemaRequest struct {
	*aws.Request
	Input *DeleteSchemaInput
	Copy  func(*DeleteSchemaInput) DeleteSchemaRequest
}

// Send marshals and sends the DeleteSchema API request.
func (r DeleteSchemaRequest) Send(ctx context.Context) (*DeleteSchemaResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteSchemaResponse{
		DeleteSchemaOutput: r.Request.Data.(*DeleteSchemaOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteSchemaResponse is the response type for the
// DeleteSchema API operation.
type DeleteSchemaResponse struct {
	*DeleteSchemaOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteSchema request.
func (r *DeleteSchemaResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
