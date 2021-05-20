// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatch

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeAnomalyDetectorsInput struct {
	_ struct{} `type:"structure"`

	// Limits the results to only the anomaly detection models that are associated
	// with the specified metric dimensions. If there are multiple metrics that
	// have these dimensions and have anomaly detection models associated, they're
	// all returned.
	Dimensions []Dimension `type:"list"`

	// The maximum number of results to return in one operation. The maximum value
	// you can specify is 10.
	//
	// To retrieve the remaining results, make another call with the returned NextToken
	// value.
	MaxResults *int64 `min:"1" type:"integer"`

	// Limits the results to only the anomaly detection models that are associated
	// with the specified metric name. If there are multiple metrics with this name
	// in different namespaces that have anomaly detection models, they're all returned.
	MetricName *string `min:"1" type:"string"`

	// Limits the results to only the anomaly detection models that are associated
	// with the specified namespace.
	Namespace *string `min:"1" type:"string"`

	// Use the token returned by the previous operation to request the next page
	// of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeAnomalyDetectorsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeAnomalyDetectorsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeAnomalyDetectorsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.MetricName != nil && len(*s.MetricName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MetricName", 1))
	}
	if s.Namespace != nil && len(*s.Namespace) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Namespace", 1))
	}
	if s.Dimensions != nil {
		for i, v := range s.Dimensions {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Dimensions", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeAnomalyDetectorsOutput struct {
	_ struct{} `type:"structure"`

	// The list of anomaly detection models returned by the operation.
	AnomalyDetectors []AnomalyDetector `type:"list"`

	// A token that you can use in a subsequent operation to retrieve the next set
	// of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeAnomalyDetectorsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeAnomalyDetectors = "DescribeAnomalyDetectors"

// DescribeAnomalyDetectorsRequest returns a request value for making API operation for
// Amazon CloudWatch.
//
// Lists the anomaly detection models that you have created in your account.
// You can list all models in your account or filter the results to only the
// models that are related to a certain namespace, metric name, or metric dimension.
//
//    // Example sending a request using DescribeAnomalyDetectorsRequest.
//    req := client.DescribeAnomalyDetectorsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/monitoring-2010-08-01/DescribeAnomalyDetectors
func (c *Client) DescribeAnomalyDetectorsRequest(input *DescribeAnomalyDetectorsInput) DescribeAnomalyDetectorsRequest {
	op := &aws.Operation{
		Name:       opDescribeAnomalyDetectors,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeAnomalyDetectorsInput{}
	}

	req := c.newRequest(op, input, &DescribeAnomalyDetectorsOutput{})
	return DescribeAnomalyDetectorsRequest{Request: req, Input: input, Copy: c.DescribeAnomalyDetectorsRequest}
}

// DescribeAnomalyDetectorsRequest is the request type for the
// DescribeAnomalyDetectors API operation.
type DescribeAnomalyDetectorsRequest struct {
	*aws.Request
	Input *DescribeAnomalyDetectorsInput
	Copy  func(*DescribeAnomalyDetectorsInput) DescribeAnomalyDetectorsRequest
}

// Send marshals and sends the DescribeAnomalyDetectors API request.
func (r DescribeAnomalyDetectorsRequest) Send(ctx context.Context) (*DescribeAnomalyDetectorsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeAnomalyDetectorsResponse{
		DescribeAnomalyDetectorsOutput: r.Request.Data.(*DescribeAnomalyDetectorsOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeAnomalyDetectorsResponse is the response type for the
// DescribeAnomalyDetectors API operation.
type DescribeAnomalyDetectorsResponse struct {
	*DescribeAnomalyDetectorsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeAnomalyDetectors request.
func (r *DescribeAnomalyDetectorsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
