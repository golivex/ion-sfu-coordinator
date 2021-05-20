// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pricing

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetProductsInput struct {
	_ struct{} `type:"structure"`

	// The list of filters that limit the returned products. only products that
	// match all filters are returned.
	Filters []Filter `type:"list"`

	// The format version that you want the response to be in.
	//
	// Valid values are: aws_v1
	FormatVersion *string `type:"string"`

	// The maximum number of results to return in the response.
	MaxResults *int64 `min:"1" type:"integer"`

	// The pagination token that indicates the next set of results that you want
	// to retrieve.
	NextToken *string `type:"string"`

	// The code for the service whose products you want to retrieve.
	ServiceCode *string `type:"string"`
}

// String returns the string representation
func (s GetProductsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetProductsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetProductsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.Filters != nil {
		for i, v := range s.Filters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Filters", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetProductsOutput struct {
	_ struct{} `type:"structure"`

	// The format version of the response. For example, aws_v1.
	FormatVersion *string `type:"string"`

	// The pagination token that indicates the next set of results to retrieve.
	NextToken *string `type:"string"`

	// The list of products that match your filters. The list contains both the
	// product metadata and the price information.
	PriceList []aws.JSONValue `type:"list"`
}

// String returns the string representation
func (s GetProductsOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetProducts = "GetProducts"

// GetProductsRequest returns a request value for making API operation for
// AWS Price List Service.
//
// Returns a list of all products that match the filter criteria.
//
//    // Example sending a request using GetProductsRequest.
//    req := client.GetProductsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pricing-2017-10-15/GetProducts
func (c *Client) GetProductsRequest(input *GetProductsInput) GetProductsRequest {
	op := &aws.Operation{
		Name:       opGetProducts,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &GetProductsInput{}
	}

	req := c.newRequest(op, input, &GetProductsOutput{})
	return GetProductsRequest{Request: req, Input: input, Copy: c.GetProductsRequest}
}

// GetProductsRequest is the request type for the
// GetProducts API operation.
type GetProductsRequest struct {
	*aws.Request
	Input *GetProductsInput
	Copy  func(*GetProductsInput) GetProductsRequest
}

// Send marshals and sends the GetProducts API request.
func (r GetProductsRequest) Send(ctx context.Context) (*GetProductsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetProductsResponse{
		GetProductsOutput: r.Request.Data.(*GetProductsOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetProductsRequestPaginator returns a paginator for GetProducts.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetProductsRequest(input)
//   p := pricing.NewGetProductsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetProductsPaginator(req GetProductsRequest) GetProductsPaginator {
	return GetProductsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetProductsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// GetProductsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetProductsPaginator struct {
	aws.Pager
}

func (p *GetProductsPaginator) CurrentPage() *GetProductsOutput {
	return p.Pager.CurrentPage().(*GetProductsOutput)
}

// GetProductsResponse is the response type for the
// GetProducts API operation.
type GetProductsResponse struct {
	*GetProductsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetProducts request.
func (r *GetProductsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
