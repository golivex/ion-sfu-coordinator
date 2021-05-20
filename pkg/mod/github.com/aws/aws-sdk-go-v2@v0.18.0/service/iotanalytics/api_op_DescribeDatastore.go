// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotanalytics

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DescribeDatastoreInput struct {
	_ struct{} `type:"structure"`

	// The name of the data store
	//
	// DatastoreName is a required field
	DatastoreName *string `location:"uri" locationName:"datastoreName" min:"1" type:"string" required:"true"`

	// If true, additional statistical information about the data store is included
	// in the response. This feature cannot be used with a data store whose S3 storage
	// is customer-managed.
	IncludeStatistics *bool `location:"querystring" locationName:"includeStatistics" type:"boolean"`
}

// String returns the string representation
func (s DescribeDatastoreInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDatastoreInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeDatastoreInput"}

	if s.DatastoreName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DatastoreName"))
	}
	if s.DatastoreName != nil && len(*s.DatastoreName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DatastoreName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeDatastoreInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.DatastoreName != nil {
		v := *s.DatastoreName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "datastoreName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.IncludeStatistics != nil {
		v := *s.IncludeStatistics

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "includeStatistics", protocol.BoolValue(v), metadata)
	}
	return nil
}

type DescribeDatastoreOutput struct {
	_ struct{} `type:"structure"`

	// Information about the data store.
	Datastore *Datastore `locationName:"datastore" type:"structure"`

	// Additional statistical information about the data store. Included if the
	// 'includeStatistics' parameter is set to true in the request.
	Statistics *DatastoreStatistics `locationName:"statistics" type:"structure"`
}

// String returns the string representation
func (s DescribeDatastoreOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeDatastoreOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Datastore != nil {
		v := s.Datastore

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "datastore", v, metadata)
	}
	if s.Statistics != nil {
		v := s.Statistics

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "statistics", v, metadata)
	}
	return nil
}

const opDescribeDatastore = "DescribeDatastore"

// DescribeDatastoreRequest returns a request value for making API operation for
// AWS IoT Analytics.
//
// Retrieves information about a data store.
//
//    // Example sending a request using DescribeDatastoreRequest.
//    req := client.DescribeDatastoreRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotanalytics-2017-11-27/DescribeDatastore
func (c *Client) DescribeDatastoreRequest(input *DescribeDatastoreInput) DescribeDatastoreRequest {
	op := &aws.Operation{
		Name:       opDescribeDatastore,
		HTTPMethod: "GET",
		HTTPPath:   "/datastores/{datastoreName}",
	}

	if input == nil {
		input = &DescribeDatastoreInput{}
	}

	req := c.newRequest(op, input, &DescribeDatastoreOutput{})
	return DescribeDatastoreRequest{Request: req, Input: input, Copy: c.DescribeDatastoreRequest}
}

// DescribeDatastoreRequest is the request type for the
// DescribeDatastore API operation.
type DescribeDatastoreRequest struct {
	*aws.Request
	Input *DescribeDatastoreInput
	Copy  func(*DescribeDatastoreInput) DescribeDatastoreRequest
}

// Send marshals and sends the DescribeDatastore API request.
func (r DescribeDatastoreRequest) Send(ctx context.Context) (*DescribeDatastoreResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeDatastoreResponse{
		DescribeDatastoreOutput: r.Request.Data.(*DescribeDatastoreOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeDatastoreResponse is the response type for the
// DescribeDatastore API operation.
type DescribeDatastoreResponse struct {
	*DescribeDatastoreOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeDatastore request.
func (r *DescribeDatastoreResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
