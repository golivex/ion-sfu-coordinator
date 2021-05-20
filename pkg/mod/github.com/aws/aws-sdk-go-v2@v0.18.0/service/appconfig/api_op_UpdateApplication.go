// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appconfig

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type UpdateApplicationInput struct {
	_ struct{} `type:"structure"`

	// The application ID.
	//
	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"ApplicationId" type:"string" required:"true"`

	// A description of the application.
	Description *string `type:"string"`

	// The name of the application.
	Name *string `min:"1" type:"string"`
}

// String returns the string representation
func (s UpdateApplicationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateApplicationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateApplicationInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateApplicationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "ApplicationId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdateApplicationOutput struct {
	_ struct{} `type:"structure"`

	// The description of the application.
	Description *string `type:"string"`

	// The application ID.
	Id *string `type:"string"`

	// The application name.
	Name *string `min:"1" type:"string"`
}

// String returns the string representation
func (s UpdateApplicationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateApplicationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opUpdateApplication = "UpdateApplication"

// UpdateApplicationRequest returns a request value for making API operation for
// Amazon AppConfig.
//
// Updates an application.
//
//    // Example sending a request using UpdateApplicationRequest.
//    req := client.UpdateApplicationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appconfig-2019-10-09/UpdateApplication
func (c *Client) UpdateApplicationRequest(input *UpdateApplicationInput) UpdateApplicationRequest {
	op := &aws.Operation{
		Name:       opUpdateApplication,
		HTTPMethod: "PATCH",
		HTTPPath:   "/applications/{ApplicationId}",
	}

	if input == nil {
		input = &UpdateApplicationInput{}
	}

	req := c.newRequest(op, input, &UpdateApplicationOutput{})
	return UpdateApplicationRequest{Request: req, Input: input, Copy: c.UpdateApplicationRequest}
}

// UpdateApplicationRequest is the request type for the
// UpdateApplication API operation.
type UpdateApplicationRequest struct {
	*aws.Request
	Input *UpdateApplicationInput
	Copy  func(*UpdateApplicationInput) UpdateApplicationRequest
}

// Send marshals and sends the UpdateApplication API request.
func (r UpdateApplicationRequest) Send(ctx context.Context) (*UpdateApplicationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateApplicationResponse{
		UpdateApplicationOutput: r.Request.Data.(*UpdateApplicationOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateApplicationResponse is the response type for the
// UpdateApplication API operation.
type UpdateApplicationResponse struct {
	*UpdateApplicationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateApplication request.
func (r *UpdateApplicationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
