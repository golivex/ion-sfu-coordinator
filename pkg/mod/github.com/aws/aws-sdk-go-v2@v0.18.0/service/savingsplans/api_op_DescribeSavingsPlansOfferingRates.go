// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package savingsplans

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DescribeSavingsPlansOfferingRatesInput struct {
	_ struct{} `type:"structure"`

	// The filters.
	Filters []SavingsPlanOfferingRateFilterElement `locationName:"filters" type:"list"`

	// The maximum number of results to return with a single call. To retrieve additional
	// results, make another call with the returned token value.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// The token for the next page of results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The specific AWS operation for the line item in the billing report.
	Operations []string `locationName:"operations" type:"list"`

	// The AWS products.
	Products []SavingsPlanProductType `locationName:"products" type:"list"`

	// The IDs of the offerings.
	SavingsPlanOfferingIds []string `locationName:"savingsPlanOfferingIds" type:"list"`

	// The payment options.
	SavingsPlanPaymentOptions []SavingsPlanPaymentOption `locationName:"savingsPlanPaymentOptions" type:"list"`

	// The plan types.
	SavingsPlanTypes []SavingsPlanType `locationName:"savingsPlanTypes" type:"list"`

	// The services.
	ServiceCodes []SavingsPlanRateServiceCode `locationName:"serviceCodes" type:"list"`

	// The usage details of the line item in the billing report.
	UsageTypes []string `locationName:"usageTypes" type:"list"`
}

// String returns the string representation
func (s DescribeSavingsPlansOfferingRatesInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeSavingsPlansOfferingRatesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Filters != nil {
		v := s.Filters

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "filters", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Operations != nil {
		v := s.Operations

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "operations", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.Products != nil {
		v := s.Products

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "products", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.SavingsPlanOfferingIds != nil {
		v := s.SavingsPlanOfferingIds

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "savingsPlanOfferingIds", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.SavingsPlanPaymentOptions != nil {
		v := s.SavingsPlanPaymentOptions

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "savingsPlanPaymentOptions", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.SavingsPlanTypes != nil {
		v := s.SavingsPlanTypes

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "savingsPlanTypes", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.ServiceCodes != nil {
		v := s.ServiceCodes

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "serviceCodes", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.UsageTypes != nil {
		v := s.UsageTypes

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "usageTypes", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

type DescribeSavingsPlansOfferingRatesOutput struct {
	_ struct{} `type:"structure"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// Information about the Savings Plans offering rates.
	SearchResults []SavingsPlanOfferingRate `locationName:"searchResults" type:"list"`
}

// String returns the string representation
func (s DescribeSavingsPlansOfferingRatesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeSavingsPlansOfferingRatesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SearchResults != nil {
		v := s.SearchResults

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "searchResults", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opDescribeSavingsPlansOfferingRates = "DescribeSavingsPlansOfferingRates"

// DescribeSavingsPlansOfferingRatesRequest returns a request value for making API operation for
// AWS Savings Plans.
//
// Describes the specified Savings Plans offering rates.
//
//    // Example sending a request using DescribeSavingsPlansOfferingRatesRequest.
//    req := client.DescribeSavingsPlansOfferingRatesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/savingsplans-2019-06-28/DescribeSavingsPlansOfferingRates
func (c *Client) DescribeSavingsPlansOfferingRatesRequest(input *DescribeSavingsPlansOfferingRatesInput) DescribeSavingsPlansOfferingRatesRequest {
	op := &aws.Operation{
		Name:       opDescribeSavingsPlansOfferingRates,
		HTTPMethod: "POST",
		HTTPPath:   "/DescribeSavingsPlansOfferingRates",
	}

	if input == nil {
		input = &DescribeSavingsPlansOfferingRatesInput{}
	}

	req := c.newRequest(op, input, &DescribeSavingsPlansOfferingRatesOutput{})
	return DescribeSavingsPlansOfferingRatesRequest{Request: req, Input: input, Copy: c.DescribeSavingsPlansOfferingRatesRequest}
}

// DescribeSavingsPlansOfferingRatesRequest is the request type for the
// DescribeSavingsPlansOfferingRates API operation.
type DescribeSavingsPlansOfferingRatesRequest struct {
	*aws.Request
	Input *DescribeSavingsPlansOfferingRatesInput
	Copy  func(*DescribeSavingsPlansOfferingRatesInput) DescribeSavingsPlansOfferingRatesRequest
}

// Send marshals and sends the DescribeSavingsPlansOfferingRates API request.
func (r DescribeSavingsPlansOfferingRatesRequest) Send(ctx context.Context) (*DescribeSavingsPlansOfferingRatesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeSavingsPlansOfferingRatesResponse{
		DescribeSavingsPlansOfferingRatesOutput: r.Request.Data.(*DescribeSavingsPlansOfferingRatesOutput),
		response:                                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeSavingsPlansOfferingRatesResponse is the response type for the
// DescribeSavingsPlansOfferingRates API operation.
type DescribeSavingsPlansOfferingRatesResponse struct {
	*DescribeSavingsPlansOfferingRatesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeSavingsPlansOfferingRates request.
func (r *DescribeSavingsPlansOfferingRatesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
