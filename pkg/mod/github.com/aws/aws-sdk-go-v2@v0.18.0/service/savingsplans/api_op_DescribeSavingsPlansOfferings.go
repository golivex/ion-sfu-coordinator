// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package savingsplans

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DescribeSavingsPlansOfferingsInput struct {
	_ struct{} `type:"structure"`

	// The currencies.
	Currencies []CurrencyCode `locationName:"currencies" type:"list"`

	// The descriptions.
	Descriptions []string `locationName:"descriptions" type:"list"`

	// The durations, in seconds.
	Durations []int64 `locationName:"durations" type:"list"`

	// The filters.
	Filters []SavingsPlanOfferingFilterElement `locationName:"filters" type:"list"`

	// The maximum number of results to return with a single call. To retrieve additional
	// results, make another call with the returned token value.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// The token for the next page of results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The IDs of the offerings.
	OfferingIds []string `locationName:"offeringIds" type:"list"`

	// The specific AWS operation for the line item in the billing report.
	Operations []string `locationName:"operations" type:"list"`

	// The payment options.
	PaymentOptions []SavingsPlanPaymentOption `locationName:"paymentOptions" type:"list"`

	// The plan type.
	PlanTypes []SavingsPlanType `locationName:"planTypes" type:"list"`

	// The product type.
	ProductType SavingsPlanProductType `locationName:"productType" type:"string" enum:"true"`

	// The services.
	ServiceCodes []string `locationName:"serviceCodes" type:"list"`

	// The usage details of the line item in the billing report.
	UsageTypes []string `locationName:"usageTypes" type:"list"`
}

// String returns the string representation
func (s DescribeSavingsPlansOfferingsInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeSavingsPlansOfferingsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Currencies != nil {
		v := s.Currencies

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "currencies", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.Descriptions != nil {
		v := s.Descriptions

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "descriptions", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.Durations != nil {
		v := s.Durations

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "durations", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.Int64Value(v1))
		}
		ls0.End()

	}
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
	if s.OfferingIds != nil {
		v := s.OfferingIds

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "offeringIds", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

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
	if s.PaymentOptions != nil {
		v := s.PaymentOptions

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "paymentOptions", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.PlanTypes != nil {
		v := s.PlanTypes

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "planTypes", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if len(s.ProductType) > 0 {
		v := s.ProductType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "productType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
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

type DescribeSavingsPlansOfferingsOutput struct {
	_ struct{} `type:"structure"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// Information about the Savings Plans offerings.
	SearchResults []SavingsPlanOffering `locationName:"searchResults" type:"list"`
}

// String returns the string representation
func (s DescribeSavingsPlansOfferingsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeSavingsPlansOfferingsOutput) MarshalFields(e protocol.FieldEncoder) error {
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

const opDescribeSavingsPlansOfferings = "DescribeSavingsPlansOfferings"

// DescribeSavingsPlansOfferingsRequest returns a request value for making API operation for
// AWS Savings Plans.
//
// Describes the specified Savings Plans offerings.
//
//    // Example sending a request using DescribeSavingsPlansOfferingsRequest.
//    req := client.DescribeSavingsPlansOfferingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/savingsplans-2019-06-28/DescribeSavingsPlansOfferings
func (c *Client) DescribeSavingsPlansOfferingsRequest(input *DescribeSavingsPlansOfferingsInput) DescribeSavingsPlansOfferingsRequest {
	op := &aws.Operation{
		Name:       opDescribeSavingsPlansOfferings,
		HTTPMethod: "POST",
		HTTPPath:   "/DescribeSavingsPlansOfferings",
	}

	if input == nil {
		input = &DescribeSavingsPlansOfferingsInput{}
	}

	req := c.newRequest(op, input, &DescribeSavingsPlansOfferingsOutput{})
	return DescribeSavingsPlansOfferingsRequest{Request: req, Input: input, Copy: c.DescribeSavingsPlansOfferingsRequest}
}

// DescribeSavingsPlansOfferingsRequest is the request type for the
// DescribeSavingsPlansOfferings API operation.
type DescribeSavingsPlansOfferingsRequest struct {
	*aws.Request
	Input *DescribeSavingsPlansOfferingsInput
	Copy  func(*DescribeSavingsPlansOfferingsInput) DescribeSavingsPlansOfferingsRequest
}

// Send marshals and sends the DescribeSavingsPlansOfferings API request.
func (r DescribeSavingsPlansOfferingsRequest) Send(ctx context.Context) (*DescribeSavingsPlansOfferingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeSavingsPlansOfferingsResponse{
		DescribeSavingsPlansOfferingsOutput: r.Request.Data.(*DescribeSavingsPlansOfferingsOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeSavingsPlansOfferingsResponse is the response type for the
// DescribeSavingsPlansOfferings API operation.
type DescribeSavingsPlansOfferingsResponse struct {
	*DescribeSavingsPlansOfferingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeSavingsPlansOfferings request.
func (r *DescribeSavingsPlansOfferingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
