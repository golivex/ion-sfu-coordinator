// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package schemas

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DescribeSchemaInput struct {
	_ struct{} `type:"structure"`

	// RegistryName is a required field
	RegistryName *string `location:"uri" locationName:"registryName" type:"string" required:"true"`

	// SchemaName is a required field
	SchemaName *string `location:"uri" locationName:"schemaName" type:"string" required:"true"`

	SchemaVersion *string `location:"querystring" locationName:"schemaVersion" type:"string"`
}

// String returns the string representation
func (s DescribeSchemaInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeSchemaInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeSchemaInput"}

	if s.RegistryName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RegistryName"))
	}

	if s.SchemaName == nil {
		invalidParams.Add(aws.NewErrParamRequired("SchemaName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeSchemaInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.RegistryName != nil {
		v := *s.RegistryName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "registryName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SchemaName != nil {
		v := *s.SchemaName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "schemaName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SchemaVersion != nil {
		v := *s.SchemaVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "schemaVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DescribeSchemaOutput struct {
	_ struct{} `type:"structure"`

	Content *string `type:"string"`

	Description *string `type:"string"`

	LastModified *time.Time `type:"timestamp" timestampFormat:"iso8601"`

	SchemaArn *string `type:"string"`

	SchemaName *string `type:"string"`

	SchemaVersion *string `type:"string"`

	// Key-value pairs associated with a resource.
	Tags map[string]string `locationName:"tags" type:"map"`

	Type *string `type:"string"`

	VersionCreatedDate *time.Time `type:"timestamp" timestampFormat:"iso8601"`
}

// String returns the string representation
func (s DescribeSchemaOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeSchemaOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Content != nil {
		v := *s.Content

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Content", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LastModified != nil {
		v := *s.LastModified

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LastModified",
			protocol.TimeValue{V: v, Format: "iso8601", QuotedFormatTime: true}, metadata)
	}
	if s.SchemaArn != nil {
		v := *s.SchemaArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SchemaArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SchemaName != nil {
		v := *s.SchemaName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SchemaName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SchemaVersion != nil {
		v := *s.SchemaVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SchemaVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.Type != nil {
		v := *s.Type

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Type", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VersionCreatedDate != nil {
		v := *s.VersionCreatedDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "VersionCreatedDate",
			protocol.TimeValue{V: v, Format: "iso8601", QuotedFormatTime: true}, metadata)
	}
	return nil
}

const opDescribeSchema = "DescribeSchema"

// DescribeSchemaRequest returns a request value for making API operation for
// Schemas.
//
// Retrieve the schema definition.
//
//    // Example sending a request using DescribeSchemaRequest.
//    req := client.DescribeSchemaRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/schemas-2019-12-02/DescribeSchema
func (c *Client) DescribeSchemaRequest(input *DescribeSchemaInput) DescribeSchemaRequest {
	op := &aws.Operation{
		Name:       opDescribeSchema,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/registries/name/{registryName}/schemas/name/{schemaName}",
	}

	if input == nil {
		input = &DescribeSchemaInput{}
	}

	req := c.newRequest(op, input, &DescribeSchemaOutput{})
	return DescribeSchemaRequest{Request: req, Input: input, Copy: c.DescribeSchemaRequest}
}

// DescribeSchemaRequest is the request type for the
// DescribeSchema API operation.
type DescribeSchemaRequest struct {
	*aws.Request
	Input *DescribeSchemaInput
	Copy  func(*DescribeSchemaInput) DescribeSchemaRequest
}

// Send marshals and sends the DescribeSchema API request.
func (r DescribeSchemaRequest) Send(ctx context.Context) (*DescribeSchemaResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeSchemaResponse{
		DescribeSchemaOutput: r.Request.Data.(*DescribeSchemaOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeSchemaResponse is the response type for the
// DescribeSchema API operation.
type DescribeSchemaResponse struct {
	*DescribeSchemaOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeSchema request.
func (r *DescribeSchemaResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
