// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// DescribeTapeRecoveryPointsInput
type DescribeTapeRecoveryPointsInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and AWS Region.
	//
	// GatewayARN is a required field
	GatewayARN *string `min:"50" type:"string" required:"true"`

	// Specifies that the number of virtual tape recovery points that are described
	// be limited to the specified number.
	Limit *int64 `min:"1" type:"integer"`

	// An opaque string that indicates the position at which to begin describing
	// the virtual tape recovery points.
	Marker *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeTapeRecoveryPointsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeTapeRecoveryPointsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeTapeRecoveryPointsInput"}

	if s.GatewayARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("GatewayARN"))
	}
	if s.GatewayARN != nil && len(*s.GatewayARN) < 50 {
		invalidParams.Add(aws.NewErrParamMinLen("GatewayARN", 50))
	}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}
	if s.Marker != nil && len(*s.Marker) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Marker", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// DescribeTapeRecoveryPointsOutput
type DescribeTapeRecoveryPointsOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and AWS Region.
	GatewayARN *string `min:"50" type:"string"`

	// An opaque string that indicates the position at which the virtual tape recovery
	// points that were listed for description ended.
	//
	// Use this marker in your next request to list the next set of virtual tape
	// recovery points in the list. If there are no more recovery points to describe,
	// this field does not appear in the response.
	Marker *string `min:"1" type:"string"`

	// An array of TapeRecoveryPointInfos that are available for the specified gateway.
	TapeRecoveryPointInfos []TapeRecoveryPointInfo `type:"list"`
}

// String returns the string representation
func (s DescribeTapeRecoveryPointsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeTapeRecoveryPoints = "DescribeTapeRecoveryPoints"

// DescribeTapeRecoveryPointsRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Returns a list of virtual tape recovery points that are available for the
// specified tape gateway.
//
// A recovery point is a point-in-time view of a virtual tape at which all the
// data on the virtual tape is consistent. If your gateway crashes, virtual
// tapes that have recovery points can be recovered to a new gateway. This operation
// is only supported in the tape gateway type.
//
//    // Example sending a request using DescribeTapeRecoveryPointsRequest.
//    req := client.DescribeTapeRecoveryPointsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/DescribeTapeRecoveryPoints
func (c *Client) DescribeTapeRecoveryPointsRequest(input *DescribeTapeRecoveryPointsInput) DescribeTapeRecoveryPointsRequest {
	op := &aws.Operation{
		Name:       opDescribeTapeRecoveryPoints,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "Limit",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeTapeRecoveryPointsInput{}
	}

	req := c.newRequest(op, input, &DescribeTapeRecoveryPointsOutput{})
	return DescribeTapeRecoveryPointsRequest{Request: req, Input: input, Copy: c.DescribeTapeRecoveryPointsRequest}
}

// DescribeTapeRecoveryPointsRequest is the request type for the
// DescribeTapeRecoveryPoints API operation.
type DescribeTapeRecoveryPointsRequest struct {
	*aws.Request
	Input *DescribeTapeRecoveryPointsInput
	Copy  func(*DescribeTapeRecoveryPointsInput) DescribeTapeRecoveryPointsRequest
}

// Send marshals and sends the DescribeTapeRecoveryPoints API request.
func (r DescribeTapeRecoveryPointsRequest) Send(ctx context.Context) (*DescribeTapeRecoveryPointsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeTapeRecoveryPointsResponse{
		DescribeTapeRecoveryPointsOutput: r.Request.Data.(*DescribeTapeRecoveryPointsOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeTapeRecoveryPointsRequestPaginator returns a paginator for DescribeTapeRecoveryPoints.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeTapeRecoveryPointsRequest(input)
//   p := storagegateway.NewDescribeTapeRecoveryPointsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeTapeRecoveryPointsPaginator(req DescribeTapeRecoveryPointsRequest) DescribeTapeRecoveryPointsPaginator {
	return DescribeTapeRecoveryPointsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeTapeRecoveryPointsInput
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

// DescribeTapeRecoveryPointsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeTapeRecoveryPointsPaginator struct {
	aws.Pager
}

func (p *DescribeTapeRecoveryPointsPaginator) CurrentPage() *DescribeTapeRecoveryPointsOutput {
	return p.Pager.CurrentPage().(*DescribeTapeRecoveryPointsOutput)
}

// DescribeTapeRecoveryPointsResponse is the response type for the
// DescribeTapeRecoveryPoints API operation.
type DescribeTapeRecoveryPointsResponse struct {
	*DescribeTapeRecoveryPointsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeTapeRecoveryPoints request.
func (r *DescribeTapeRecoveryPointsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
